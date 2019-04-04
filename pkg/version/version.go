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

package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// BuildVersion injected by the golang compiler
var BuildVersion string

// ErrVersionParse definition
var ErrVersionParse = errors.New("Failed to parse BuildVersion string")

// Version defines a version object that holds build information
type Version struct {
	Major         int
	Minor         int
	Patch         int
	GitHash       string
	GitAuthorDate string
	GoVersion     string
	GoArch        string
}

// FromString parses a version string to a version object
func FromString(v string) (*Version, error) {
	fields := strings.Split(v, " ")
	if len(fields) != 7 {
		return nil, ErrVersionParse
	}

	vFields := strings.Split(fields[0], ".")
	if len(vFields) != 3 {
		return nil, ErrVersionParse
	}

	major, err := strconv.Atoi(vFields[0])
	if err != nil {
		return nil, ErrVersionParse
	}

	minor, err := strconv.Atoi(vFields[1])
	if err != nil {
		return nil, ErrVersionParse
	}

	patch, err := strconv.Atoi(vFields[2])
	if err != nil {
		return nil, ErrVersionParse
	}

	gitHash := fields[1]
	if len(gitHash) != 7 {
		return nil, ErrVersionParse
	}

	gitAuthorDate := fields[2]
	if len(gitAuthorDate) != 25 {
		return nil, ErrVersionParse
	}

	return &Version{
		Major:         major,
		Minor:         minor,
		Patch:         patch,
		GitHash:       gitHash,
		GitAuthorDate: gitAuthorDate,
		GoVersion:     fields[5],
		GoArch:        fields[6],
	}, nil
}

// ToString converts a version object to a string
func ToString(v *Version) string {
	return fmt.Sprintf("%d.%d.%d %s %s go version %s %s",
		v.Major, v.Minor, v.Patch,
		v.GitHash, v.GitAuthorDate,
		v.GoVersion, v.GoArch)
}

// Current returns the BuildVersion as Version object
func Current() (*Version, error) {
	return FromString(BuildVersion)
}
