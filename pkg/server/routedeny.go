/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	//batchv1beta1 "k8s.io/api/batch/v1beta1"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
)

// deny configmaps with specific key-value pair.
func routeCreateDeny(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	klog.V(2).Info("admitting route")
	/* unsure how to do this for route
	cronjobresource := metav1.GroupVersionResource{Group: "batch", Version: "v1beta1", Resource: "cronjobs"}
	if ar.Request.Resource != cronjobresource {
		klog.Errorf("expect resource to be %s", cronjobresource)
		return nil
	}
	*/
	reviewResponse := v1beta1.AdmissionResponse{}
	isPriv := checkNamespace(ar.Request.Namespace)
	reviewResponse.Allowed = true
	if isPriv {
		klog.V(2).Info("privileged namespace approved for route")
		return &reviewResponse
	}
	raw := ar.Request.Object.Raw
	route := routeapi.Route{}
	deserializer := codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(raw, nil, &route); err != nil {
		klog.Error(err)
		return toAdmissionResponse(err)
	}
	if route.Spec.Host != "" {
		reviewResponse.Allowed = false
		reviewResponse.Result = &metav1.Status{
			Reason: "route with custom host may not be created in this namespace.",
		}
	}
	return &reviewResponse
}
