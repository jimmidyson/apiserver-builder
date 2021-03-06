/*
Copyright 2017 The Kubernetes Authors.

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

package v1beta1_test

import (
	"os"
	"testing"

	"github.com/kubernetes-incubator/apiserver-builder/example/pkg/apis"
	"github.com/kubernetes-incubator/apiserver-builder/example/pkg/client/clientset_generated/clientset"
	"github.com/kubernetes-incubator/apiserver-builder/example/pkg/openapi"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/test"
	"k8s.io/client-go/rest"
)

var testenv *test.TestEnvironment
var config *rest.Config
var cs *clientset.Clientset

// Do Test Suite setup / teardown
func TestMain(m *testing.M) {
	testenv = test.NewTestEnvironment()
	config = testenv.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	cs = clientset.NewForConfigOrDie(config)
	retCode := m.Run()
	testenv.Stop()
	os.Exit(retCode)
}
