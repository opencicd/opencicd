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
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/opencicd/opencicd/pkg/version"
)

var _ = Describe("Version", func() {

	Describe("FromString", func() {

	})

	DescribeTable("FromString",
		func(vString string, vObject *Version, errStatus error) {
			v, err := FromString(vString)

			if errStatus != nil {
				Expect(err).To(Equal(errStatus))
			}
			Expect(v).To(Equal(vObject))
		},
		Entry(
			"valid version string",
			"1.2.3 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64",
			&Version{
				Major:         1,
				Minor:         2,
				Patch:         3,
				GitHash:       "95a4b3d",
				GitAuthorDate: "2019-04-04T10:08:26+02:00",
				GoVersion:     "go1.12.1",
				GoArch:        "darwin/amd64",
			},
			nil,
		),
		Entry(
			"too many items",
			"1.2.3 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64 foo bar",
			nil,
			ErrVersionParse,
		),
		Entry(
			"too few items",
			"foo bar",
			nil,
			ErrVersionParse,
		),
		Entry(
			"invalid major version",
			"a.2.3 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64",
			nil,
			ErrVersionParse,
		),
		Entry(
			"invalid minor version",
			"1.a.3 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64",
			nil,
			ErrVersionParse,
		),
		Entry(
			"invalid patch version",
			"1.2.a 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64",
			nil,
			ErrVersionParse,
		),
		Entry(
			"too many version items",
			"1.2.3.4 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64",
			nil,
			ErrVersionParse,
		),
		Entry(
			"too few version items",
			"1.2 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64",
			nil,
			ErrVersionParse,
		),
		Entry(
			"invalid git hash",
			"1.2.3 abc 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64",
			nil,
			ErrVersionParse,
		),
		Entry(
			"invalid git author date",
			"1.2.3 95a4b3d 123abc go version go1.12.1 darwin/amd64",
			nil,
			ErrVersionParse,
		),
	)

	Describe("ToString", func() {

		It("Should create a valid string", func() {
			vObject := &Version{
				Major:         1,
				Minor:         2,
				Patch:         3,
				GitHash:       "95a4b3d",
				GitAuthorDate: "2019-04-04T10:08:26+02:00",
				GoVersion:     "go1.12.1",
				GoArch:        "darwin/amd64",
			}

			vString := ToString(vObject)

			Expect(vString).To(Equal("1.2.3 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64"))
		})
	})

	Describe("Current", func() {

		It("Should create a valid version out of the BuildVersion", func() {
			BuildVersion = "1.2.3 95a4b3d 2019-04-04T10:08:26+02:00 go version go1.12.1 darwin/amd64"

			vObject := &Version{
				Major:         1,
				Minor:         2,
				Patch:         3,
				GitHash:       "95a4b3d",
				GitAuthorDate: "2019-04-04T10:08:26+02:00",
				GoVersion:     "go1.12.1",
				GoArch:        "darwin/amd64",
			}

			v, err := Current()

			Expect(err).NotTo(HaveOccurred())
			Expect(v).To(Equal(vObject))
		})
	})
})
