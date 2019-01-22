# kube-vwh
Kube Validating WebHook

You can follow the steps here to create your own CA signed cert:
https://deliciousbrains.com/ssl-certificate-authority-for-local-https-development/

No need to install to your localhost trust unless you plan on using in the future.

You'll need to base64 encode the contents of myca.pem and embed in clientconfig.
See assets/vwh.yaml for example.

```sh
minikube start \
--vm-driver kvm2 \
-v 9
--extra-config=apiserver.enable-admission-plugins="NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,Priority,ResourceQuota"
```
