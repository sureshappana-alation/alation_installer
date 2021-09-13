# alation_installer
AL-69821

# Unified installer Build Pipeline
Unified installer build pipeline is available in `.github/unified-installer-build-pipeline.yml`. It consists of two jobs:
    
    1. Build: Builds the go installer using docker image
    2. App-setup: Consolidates all module and dependency artifacts and generate compressed archive
        a. Reads version files from `alation_installer/versions/` directory.
        b. Apply overrides (in case of manually triggered the workflow).
        c. Download dependency ( Kurl, nginx etc) and module tar files from S3.
        d. Compress the final build and uploads to S3.

## Versions

1. Versions/ directory consists of version of modules (OCF, AA etc) and other required dependencies.
2. Module name in version file should match with the artifact available in S3.

    Ex: `versions/kurl.env` consists of entry as `KURL=d2b213e-86607afd5b` where a file exists in S3 with the name `kurl-d2b213e-86607afd5b.tar.gz`

3. Entry name in versions file should follow the following naming convention

    a. Name should contain only numbers, characters

    b. Name should not start with number

    c. Names should be in uppercase
    
    d. No special characters allowed


Module and dependency versions are very important to consider. Versions are handled in the following way:

- `alation_installer/versions/` directory has versions of all modules and dependencies( Kurl, nginx etc ). These are considered as **base versions**.

- When triggering manually or REST, versions passed will be considered as **override versions**. These values will be overridden on base versions.

- Overridden version can contain version or can be left as empty. If the overridden value is left empty then base version is considered.

### Who can update the versions?
- All required dependencies ( Kurl, nginx etc ) version will be maintained by the Cloud Infra Team.
- Module teams will have the provision to update their versions by raising PR.

### Module Packages:
- Modules should be available in S3 bucket **unified-installer-build-pipeline-dev** in `<module-name in small letters>-<version>.tar.gz` format for build pipeline to consider.

- Module tar file content format is available in (confluence](https://alationcorp.atlassian.net/wiki/spaces/PLAT/pages/5595555720/Alation+Installer+Platform+Contract+Design+Document)

## Steps to add new module/dependency?:
- Application teams open Jira ticket to onboard a module/application to unified installer build pipeline.
- Cloud infra team works on the ticket to onboard the new module.

## Trigger the Build
Build could be triggered in 3 ways:
1. UI: Manually by going to github actions workflow and entering parameters.
2. By using REST API.
3. Automatically when PR merged successfully.

Currently Manual trigger is only enabled. REST API examples will be added later.

### Optional parameters when triggering the build:

1. Override versions of modules
2. exclude modules list: This is comma separated modules names (Names should be in the same format mentioned in versions/*.env)

### Special treatment for Kurl package:
- Kurl version consists of two strings delimited by hyphen `d2b213e-86607afd5b1cf95baa25cc28d4668568` where first section indicates hash received from Kurl.sh and second section indicates md5 hash of air gapped Kurl package (In this case hash of kurl-d2b213e.tar.gz). This is to make sure Kurl package used in the unified installer is the one developed and tested against. (Workflow can't get the package from Kurl.sh because of incompatibility issue, we might needs to see how to automate that in future.)
- Github actions caches the kurl package to avoid repeated downloads from S3.

## Unified installer setup:
### Env variables:
The following secrets are managed in github actions settings. These values will be pulled dynamically at the time of workflow run.

`AWS_ACCESS_KEY_ID:` AWS access key id for github actions

`AWS_SECRET_ACCESS_KEY:` AWS access key value for github actions

`AWS_S3_REGION`: AWS S3 bucket region

`S3_DEV_BUCKET_NAME`: AWS S3 bucket used for input (for dependency and module tar files)

`S3_RELEASE_BUCKET_NAME`: AWS S3 bucket where unified installer bundles will be stored.

### Versioning of Unified installer bundle:
Currently unified installer bundle uses the following versioning schema `alation-k8s-<branch>-<date in YYYYMMDD>.<build-run-number>.tar.gz`
