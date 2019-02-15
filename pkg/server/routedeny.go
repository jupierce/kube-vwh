/*
Copyright 2019 Red Hat, Inc. and/or its affiliates

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
	"fmt"
	"encoding/json"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	routeapi "github.com/openshift/origin/pkg/route/apis/route"
)


func routeCreateDeny(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	klog.Errorf("admitting route")

	routeresource := metav1.GroupVersionResource{Group: "route.openshift.io", Version: "v1", Resource: "routes"}
	if ar.Request.Resource != routeresource {
		klog.Errorf("expect resource to be %s, found %v", routeresource, ar.Request.Resource)
		return nil
	}
	reviewResponse := v1beta1.AdmissionResponse{}
	isPriv := checkNamespace(ar.Request.Namespace)
	reviewResponse.Allowed = true
	if isPriv {
		klog.Errorf("privileged namespace approved for route")
		return &reviewResponse
	}
	raw := ar.Request.Object.Raw
	route := routeapi.Route{}
	err := json.Unmarshal(raw, &route)
	if err != nil {
		klog.Error(err)
		return toAdmissionResponse(err)
	}
	klog.Errorf(fmt.Sprintf("route.Spec.Host = %v", route.Spec.Host))
	if route.Spec.Host != "" {
		reviewResponse.Allowed = false
		reviewResponse.Result = &metav1.Status{
			Reason: "route with custom host may not be created in this namespace.",
		}
	}
	return &reviewResponse
}
