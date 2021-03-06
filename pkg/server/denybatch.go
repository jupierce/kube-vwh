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
	// batchv1beta1 "k8s.io/api/batch/v1beta1"
	// batchv1 "k8s.io/api/batch/v1"
)

// deny cronjobs in non-privileged namespaces.
func cronjobCreateDeny(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	klog.V(2).Info("admitting cronjob")
	cronjobresource := metav1.GroupVersionResource{Group: "batch", Version: "v1beta1", Resource: "cronjobs"}
	if ar.Request.Resource != cronjobresource {
		klog.Errorf("expect resource to be %s", cronjobresource)
		return nil
	}

	reviewResponse := v1beta1.AdmissionResponse{}
	isPriv := checkNamespace(ar.Request.Namespace)
	reviewResponse.Allowed = isPriv
	if !isPriv {
		reviewResponse.Result = &metav1.Status{
			Reason: "cronjob may not be created in this namespace.",
		}
	}
	return &reviewResponse
}

// deny Jobs in non-privileged namespaces.
func jobCreateDeny(ar v1beta1.AdmissionReview) *v1beta1.AdmissionResponse {
	klog.V(2).Info("admitting job")
	jobresource := metav1.GroupVersionResource{Group: "batch", Version: "v1", Resource: "jobs"}
	if ar.Request.Resource != jobresource {
		klog.Errorf("expect resource to be %s", jobresource)
		return nil
	}

	reviewResponse := v1beta1.AdmissionResponse{}
	isPriv := checkNamespace(ar.Request.Namespace)
	reviewResponse.Allowed = isPriv
	if !isPriv {
		reviewResponse.Result = &metav1.Status{
			Reason: "job may not be created in this namespace.",
		}
	}
	return &reviewResponse
}
