# kube-vwh
Kube Validating WebHook

# Deployment on OpenShift cluster

On 4.x, everything is configured on the cluster by default.

First thing you'll need is the service-signer CA:

Get your CA bundle here:
```sh
kubectl get configmaps --namespace openshift-service-cert-signer -o yaml signing-cabundle
```

Save the CA contents as a base64 encoded string for later.

Next, you can cd to 'assets' dir of this project and start applying manifests.

```sh
# create namespace named 'development' for testing purposes.  Optional.
kubectl apply -f namespace-dev.json

# Create service and tls keypair our kube-vwh
kubectl apply -f vwh-service.yaml

# Create a deployment of our kube-vwh
kubectl apply -f vwh-deployment.yaml

# Create our first ValidatingWebhookConfiguration
# Be sure to paste your base64 CA into this file before executing:
kubectl apply -f cron-vwh-svc.yaml

# Try to make a cron in an unprivileged namespace; should fail with appropriate
# error message
kubectl apply -f cron-test.yaml --namespace development

# Try again in privileged namespace; should succeed.
kubectl apply -f cron-test.yaml --namespace default

# Create you VWHC for routes:
# Be sure to add your CA in this file.
kubectl apply -f route-vwh-svc.yaml

# Try to make a route with custom host; should fail with appropriate message:
kubectl apply -f route-test.yaml --validate=false --namespace development
```


# Development/Build

## Run tests

```sh
go test ./pkg/server/ -test.v
```

## Local binary:
```sh
make build
```

## Image

```sh
# as root/sudo
buildah bud -t docker.io/mgugino/kube-vwh:latest .
buildah push docker.io/mgugino/kube-vwh:latest
```

# Local development on minikube

If you're not testing against openshift-specific components (such as Routes),
you can use minikube to quickly prototype and iterate validating webhooks.


You can follow the steps here to create your own CA signed cert:
https://deliciousbrains.com/ssl-certificate-authority-for-local-https-development/

No need to install to your localhost trust unless you plan on using in the future.
Note, I couldn't get the above to work with IPs, using an actual DNS entry for the
cert helped.

You'll need to base64 encode the contents of myca.pem and embed in clientconfig.
See assets/vwh-local.yaml for example.

```sh
minikube start \
--vm-driver kvm2 \
-v 9
--extra-config=apiserver.enable-admission-plugins="NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,Priority,ResourceQuota"
```

# Import references

https://docs.okd.io/latest/architecture/additional_concepts/dynamic_admission_controllers.html

https://docs.okd.io/latest/dev_guide/secrets.html#service-serving-certificate-secrets

Good overview: https://banzaicloud.com/blog/k8s-admission-webhooks/
