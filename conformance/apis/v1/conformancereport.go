/*
Copyright 2023 The Kubernetes Authors.

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

package v1

import (
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConformanceReport is a report of conformance testing results including the
// specific conformance profiles that were tested and the results of the tests
// with summaries and statistics.
type ConformanceReport struct {
	metav1.TypeMeta `json:",inline"`
	Implementation  `json:"implementation"`

	// Date indicates the date that this report was generated.
	Date string `json:"date"`

	// GatewayAPIVersion indicates which release version of Gateway API this
	// test report was made for.
	GatewayAPIVersion string `json:"gatewayAPIVersion"`

	// Mode is the operating mode the implementation used to run conformance tests.
	Mode string `json:"mode"`

	// GatewayAPIChannel indicates which release channel of Gateway API this
	// test report was made for.
	GatewayAPIChannel string `json:"gatewayAPIChannel"`

	// ProfileReports is a list of the individual reports for each conformance
	// profile that was enabled for a test run.
	ProfileReports []ProfileReport `json:"profiles"`

	// SucceededProvisionalTests is a list of the names of the provisional tests that
	// have been successfully run.
	SucceededProvisionalTests []string `json:"succeededProvisionalTests,omitempty"`

	// InferredSupportedFeatures indicates whether the supported features were
	// automatically detected by the conformance suite.
	InferredSupportedFeatures bool `json:"inferredSupportedFeatures"`
}

// Implementation provides metadata information on the downstream
// implementation of Gateway API which ran conformance tests.
type Implementation struct {
	// Organization refers to the company, group or individual which maintains
	// the named implementation. Organizations can provide reports for any
	// number of distinct Gateway API implementations they maintain, but need
	// to identify themselves using this organization field for grouping.
	Organization string `json:"organization"`

	// Project indicates the name of the project or repository for a Gateway API
	// implementation.
	Project string `json:"project"`

	// URL indicates a human-usable URL where more information about the
	// implementation can be found. For open source projects this should
	// generally link to the code repository.
	URL string `json:"url"`

	// Version indicates the version of the implementation that was used for
	// testing. This should generally be a semver version when applicable.
	Version string `json:"version"`

	// Contact is contact information for the maintainers so that Gateway API
	// maintainers can get ahold of them as needed. Ideally this should be
	// GitHub usernames (in the form of `@<username>`) or team names (in the
	// form of `@<team>/<name>`), but when that's not possible it can be email
	// addresses.
	// Rather than GitHub usernames or email addresses you can provide a URL to the relevant
	// support pages for the project. Ideally this would be something like the issue creation page
	// on a repository, but for projects without a publicly exposed repository a general support
	// page URL can be provided.
	Contact []string `json:"contact"`
}

// Validate ensures that the Implementation struct has valid fields set
func (i *Implementation) Validate() error {
	// TODO: add data validation https://github.com/kubernetes-sigs/gateway-api/issues/2178
	if i.Organization == "" {
		return errors.New("implementation's organization cannot be empty")
	}
	if i.Project == "" {
		return errors.New("implementation's project cannot be empty")
	}
	if i.URL == "" {
		return errors.New("implementation's url cannot be empty")
	}
	if i.Version == "" {
		return errors.New("implementation's version cannot be empty")
	}
	if len(i.Contact) == 0 {
		return errors.New("implementation's contact cannot be empty")
	}
	return nil
}
