# appset-test-plugin

1. Install Argo CD v2.8 or newer - [RELEASES](https://github.com/argoproj/argo-cd/releases)
2. Add the plugin `kustomize build manifests/ | kubectl apply -f -`
3. Test "barebones" example `kubectl apply -f manifests/1.sample-appset.yaml`
4. To test parameters, use `kubectl apply -f manifests/2.sample-appset.yaml`

> **NOTE** Run `kubectl delete -f manifests/1.sample-appset.yaml` before testing the second one
