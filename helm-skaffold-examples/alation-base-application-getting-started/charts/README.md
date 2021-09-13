# Explanation about Helm Chart
The following is the directory structure of a helm chart:

```
.
├── Chart.yaml ---> Contains the information about the helm chart (kind of helm chart metadata)
├── LICENSE ---> Contains the license information
├── README.md ---> Information about helm chart. When you update the helm chart update some information about your service here.
├── templates  ---> Contains the templates for Kubernetes resources
│   ├── NOTES.txt  ---> Output to be displayed to the console after successful installation
│   ├── _helpers.tpl  ---> Generic functions helpful for using in the template files.
│   ├── configmap.yaml  ---> Create one or more configmap files (file mounts or regular key-value pairs)  [Official docs](https://kubernetes.io/docs/concepts/configuration/configmap/)
│   ├── deployment.yaml  ---> Create one or more secret files (file mounts or regular key-value pairs) [Official docs](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
│   ├── hpa.yaml ---> K8s manifest template for horizontal pod auto scaler [Official Docs](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
│   ├── ingress.yaml ---> K8s manifest template for ingress definition [Official Docs](https://kubernetes.io/docs/concepts/services-networking/ingress/)
│   ├── secrets.yaml ---> K8s manifest template for creating one or more secrets based on information from values.yaml  [Official docs](https://kubernetes.io/docs/concepts/configuration/secret/)
│   ├── service.yaml ---> K8s manifest template for service [Official Docs](https://kubernetes.io/docs/concepts/services-networking/service/)
│   ├── serviceaccount.yaml ---> K8s manifest template for service account [Official Docs](https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/)
└── values.yaml ---> Values file that drive the K8s manifests
```

# Recommendation
- In general, you don't need to modify the templates, everything should be driven from the values file. The reason is for the customization is to reuse the same helm chart for different environments.

- In case if you require to add any k8s resources create a template file and drive them from values.yaml.


# Dependency Management

## Step 1: Add remote repository 
If subcharts exists locally, skip to step2.

Repository is the combination of repo name and url where the charts exists. To add repo to helm repository:
```sh
helm repo add <repo-name> <repo-url>
```
For example in the above case:

```sh
helm repo add bitnami https://charts.bitnami.com/bitnami
```
So that in all places you can refer the chart with `@bintanmi` instead of specifying full url in each and every place.


## Step2: Add subchart in charts.yaml file
One helm chart can have zero or more dependencies. Dependencies are defined in `Chart.yaml` file.

Example dependency is:

```yaml
dependencies: # A list of the chart requirements (optional)
  - name: nginx # The name of the chart
    version: 9.1.0 # The version of the chart
    repository: "@bitnami" # (optional) The repository URL ("https://example.com/charts") or alias ("@repo-name")
                           # Local charts can be specified with file://../
    condition: nginx.enabled #(optional) A yaml path that resolves to a boolean, used for enabling/disabling charts (e.g. subchart1.enabled )
    tags: # (optional)
      - nginx # 
    # import-values: # (optional)
    #   - ImportValues holds the mapping of source values to parent key to be imported. Each item can be a string or pair of child/parent sublist items.
    alias: nginx
```

## Step 3: Communication between parent and subchart (optional)

### Data passing from parent chart to subchart

Values from parent chart to subchart can be passed similar to other values in values.yaml file.
For example, to pass values from parent chart to above nginx subchart:

```yaml
nginx:
  key1: value1
```

### Data passing from subchart to parent chart
Import-values are used for this purpose.

For example:

Parent's Chart.yaml file:
```yaml
dependencies:
  - name: subchart
    repository: http://localhost:10191
    version: 0.1.0
    import-values:
      - data
```
Subchart's values.yaml file:

```yaml
exports:
  data:
    myint: 99
```

From the above `data` holds the values from subchart.

### Global values
Global values are values that can be accessed from any chart or subchart by exactly the same name. Global values are defined values.yaml file like

```yaml
global:
  key1: value1
```
In templates those values can be used in the similar way

```yaml
{{ .Values.global.key1 }}
```

More information about helm subcharts and global values are available [here](https://helm.sh/docs/chart_template_guide/subcharts_and_globals/).


## Step 4: Fetch Subcharts (in case of subcharts referenced from remote repository)
After defining subcharts in charts.yaml, to download subcharts from remote repository to local

```sh
helm dependency update
```

This will create a chart.lock file and charts/ directory in the location where chart.yaml file exists.

`charts` subdirectory contains the subchart .tgz files.

`chart.lock` file contains the versions of the charts downloaded. To have consistent chart dependencies version check-in this file to code repository.

An example chart.lock file looks like

```yaml
dependencies:
- name: nginx
  repository: https://charts.bitnami.com/bitnami
  version: 9.1.0
digest: sha256:387925c987914f8e502b5cf92d17358bd3a044a8d3934d11b1ece2ef575101ed
generated: "2021-06-08T11:18:23.338911-04:00"
```
