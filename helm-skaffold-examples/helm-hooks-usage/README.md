# helm-hooks-usage

Instructions from [alation-base-application-getting-started/README.md](../alation-base-application-getting-started/README.md) are still applicable here.

Additionally, this example demonstrates the usage of helm pre-install hooks and running kubectl commands from the application/service.

Here an example application is available in example-app. It has a simple docker image which creates a secret as part of helm pre-install hook.

An example also provided in example-app/setup.sh to update an existing secret

Main files here are [charts/templates/role.yaml](charts/templates/role.yaml) and [charts/templates/rolebinding.yaml](charts/templates/rolebinding.yaml). These gives required permissions for the application to connect with APIServer to perform operations on secrets. Similar to permissions mentioned in charts/templates/role.yaml other permissions can be added to any k8s resource.

More information on RBAC is available [here](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
