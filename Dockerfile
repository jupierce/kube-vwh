FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 AS builder
WORKDIR /go/src/github.com/michaelgugino/kube-vwh
COPY . .
RUN make build

FROM registry.svc.ci.openshift.org/openshift/origin-v4.0:base
COPY --from=builder /go/src/github.com/michaelgugino/kube-vwh/bin/kube-vwh /usr/bin/
ENTRYPOINT ["/usr/bin/kube-vwh"]
