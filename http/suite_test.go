/*
Copyright 2020 Noah Kantrowitz

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

package http

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	cu "github.com/coderanger/controller-utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	migrationsv1beta1 "github.com/coderanger/migrations-operator/api/v1beta1"
)

var suiteHelper *cu.FunctionalSuiteHelper
var url string

func TestHttp(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"API Server Suite",
		[]Reporter{printer.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	// Set up where the API server will listen.
	port := 40000 + rand.Intn(10000)
	os.Setenv("API_LISTEN", fmt.Sprintf("localhost:%d", port))
	url = fmt.Sprintf("http://localhost:%d/", port)

	By("bootstrapping test environment")
	suiteHelper = cu.Functional().
		API(migrationsv1beta1.AddToScheme).
		MustBuild()

	close(done)
}, 60)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	suiteHelper.MustStop()
})
