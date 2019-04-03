// Copyright 2019 OpenCICD Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/opencicd/opencicd/pkg/version"
)

var _ = Describe("Version", func() {

	Describe("Version", func() {

		It("Should not be nil", func() {
			Expect(&Version{}).NotTo(BeNil())
		})
	})

	Describe("FromString", func() {

		It("Should return empty Version object if string is empty", func() {
			v := FromString("")

			Expect(v).To(BeZero())
		})
	})
})
