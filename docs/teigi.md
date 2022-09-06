# argocd-vault-and-teigi-plugin

This plugin is an extension of the [official one](https://github.com/argoproj-labs/argocd-vault-plugin) developped by IBM.

This add the support for Teigi service.

# Requirements

The extension for teigi service require `kinit` and `tbag` commands installed. This is can be done via Config Management Plugin in ArgoCD.

# How to use it (command line)

Define a configuration file like:

```yaml
AVP_TYPE: teigisecretsmanager
AVP_TEIGI_USER: myuser
AVP_TEIGI_PORT: 4444
AVP_TEIGI_HOSTNAME: whatever.cern.ch
AVP_TEIGI_PASSWORD: mypassword
```

Pick a secret definition like one below:

```yaml
apiVersion: v1
data:
  config: <path:TEIGI_SERVICE#TEIGI_KEY>
kind: Secret
metadata:
  annotations:
    managed-by: argocd.argoproj.io
  labels:
    argocd.argoproj.io/secret-type: cluster
  name: k8s-ims-dev-b
  namespace: argocd
type: Opaque
```

Generate the binary running `make` and then to test it locally:

`./argocd-vault-plugin generate small_secret.yaml -c teigi-conf.yaml`.

If you want to test changes without rebuild all the time you can actually run

`go run main.go generate small_secret.yaml -c teigi-conf.yaml`