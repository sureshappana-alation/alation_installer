# alation-base-application-getting-started
This is the base helm chart application teams can use for development of their services. 

# Usage
- It's recommended that application teams maintain the helm charts in their code repo root folder under `charts/` directory.
- Under `charts/` you can have one or more helm charts. For example
```
./root
├── charts/
      ├── service A/
            ├── templates/
            ├── Chart.yaml
            ├── .....
      ├── service B/
            ├── templates/
            ├── Chart.yaml
            ├── .....

```
- The example chart provides a base skelton of a chart. Teams have the freedom to modify and update them based on usage.
- This helm chart was developed on top of base skelton created by `helm create` by considering some of the best practices and also added missing templates.


# Best Practices
- Don't hardcode values in template files unless there is no other way to do. Instead pass values from the `values.yaml` file.
- You can have more than one values.yaml file. ONLY if it is required you can split the values.yaml to multiple.
- For any other required k8s resources create a new file in `template/` and manage it from values.yaml.
- Don't make one big helm chart i.e, don't add dependency deployments, services etc in the same helm chart of service. Instead create a separate helm chart for dependencies and add them as subcharts in main chart. This gives more flexibility to manage subcharts.
- Update [Notes.md](charts/templates/NOTES.txt) file with information about your applicaton/service on how to use and access the application/service.

# What is covered in this basic helm chart
- Alation basic helm chart creation.
- Usage of Skaffold together with helm (Skaffold has many other potentials, i.e, creation of profiles etc, but as per our requirement this example scope is limited for local development).
- Manipulating k8s resource manifests from values.yaml
- Dependent helm chart management.

Please look into the [README](charts/README.md) in the `template/` for more information and usage.

# Prerequisite
Follow steps to setup local development environment.
https://confluence.alationdata.com/display/ENG/Kubernetes+Developer+Environment+Setup+and+Use

# Helm
[Helm](https://helm.sh/docs/) is the package manager for Kubernetes. Instead of managing each k8s resource manifest alone, helm brings all resources together and helps to manage them as a single unit.

Helm chart is available in charts/ directory. More information about helm chart is available in it's [ReadMe](charts/README.md) file.

To **install** helm package:

```bash
helm install <custom-name> <charts-dir>
```

To dry-run without actually installing:


```bash
helm install <custom-name> <charts-dir> --dry-run
```

Example:
```bash
helm install mychart ./charts --dry-run
```

(Please note that if you are using the Skaffold with helm, Skaffold automatically performs the install/upgrade. Please refer Skaffold section for usage.)

To **upgrade** installed helm chart:

```bash
helm upgrade <release> <charts-dir>
```

Example:
```bash
helm upgrade mychart ./charts
```

To **uninstall** installed helm chart:

```bash
helm uninstall <release>
```

Example:
```bash
helm uninstall mychart
```

More info on list of helm commands and usage is available [here](https://helm.sh/docs/helm/)

To **package** helm chart:

To package helm chart

```bash
helm package <chart-path>
```

This will create an output file with the extension .tgz
By default this will package all the files & subdirectories exists in charts/ directory. **If it contains subcharts, this command will package them also.**

To update dependencies before packaging use the following flag
```bash
-u (or) --dependency-update        update dependencies from "Chart.yaml" to dir "charts/" subdirectory before packaging
```

# Helm + Skaffold
[Skaffold](https://skaffold.dev/docs/) handles the workflow for building, pushing and deploying the application.
More info on configuring skaffold with helm is available [here](https://skaffold.dev/docs/pipeline-stages/deployers/helm/)


The advantage of Skaffold is that when a developer saves a change to their code, Skaffold will automatically rebuild the appropriate docker images and run helm upgrade to update the service in their cluster.

To run skaffold pipeline in development mode (skaffold watches for file changes and deploys automatically on save):

```bash
skaffold dev
```

To run skaffold pipeline once (Simply deploys to the cluster and exits. Don't watch for file changes.)

```
skaffold run
```

To run specific steps in Skaffold
```bash
skaffold <step>
```
ex:
```bash
skaffold build

skaffold deploy
```

To delete existing deployment

```bash
skaffold delete
```

More information about values you can override and use in skaffold available [here](https://skaffold.dev/docs/references/yaml/#deploy-helm-releases-setValues)

Skaffold also has inline documentation yaml file available [here](https://skaffold.dev/docs/references/yaml/).
