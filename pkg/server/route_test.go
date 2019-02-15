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
	"io/ioutil"
	"testing"
	"k8s.io/api/admission/v1beta1"
)

type routeRequestTestCase struct {
	exampleRequestFile string
	shouldPass	bool
}

var routeRequestTestCases = [...]routeRequestTestCase {
	routeRequestTestCase {exampleRequestFile: "../../util/example_request_route_custom_host.json", shouldPass: false},
	routeRequestTestCase {exampleRequestFile: "../../util/example_request_route_custom_host_allowed.json", shouldPass: true},
}

func TestRouteHost(t *testing.T) {
	for i, tcase := range routeRequestTestCases {
		dat, ferr := ioutil.ReadFile(tcase.exampleRequestFile)
		if ferr != nil {
			t.Fatal("Unable to read in json file")
		}
		requestedAdmissionReview := v1beta1.AdmissionReview{}
		responseAdmissionReview := v1beta1.AdmissionReview{}
		deserializer := codecs.UniversalDeserializer()
		if _, _, err := deserializer.Decode(dat, nil, &requestedAdmissionReview); err != nil {
			t.Fatal("Unable to process input json testcase.")
		} else {
			responseAdmissionReview.Response = routeCreateDeny(requestedAdmissionReview)
		}
		if responseAdmissionReview.Response.Allowed != tcase.shouldPass {
			t.Fatal(fmt.Sprintf("routeRequestTestCase %v failed", i))
		}
	}
}


func TestRouteInvalid(t *testing.T) {
	dat, ferr := ioutil.ReadFile("../../util/example_request.json")
	if ferr != nil {
		t.Fatal("Unable to read in json file")
	}
	requestedAdmissionReview := v1beta1.AdmissionReview{}
	responseAdmissionReview := v1beta1.AdmissionReview{}
	deserializer := codecs.UniversalDeserializer()
	if _, _, err := deserializer.Decode(dat, nil, &requestedAdmissionReview); err != nil {
		t.Fatal("Unable to process input json testcase.")
	} else {
		responseAdmissionReview.Response = routeCreateDeny(requestedAdmissionReview)
	}
	if responseAdmissionReview.Response != nil {
		t.Fatal("Expected nil response for wrong request resource type for route.")
	}
}
