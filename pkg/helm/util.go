/*******************************************************************************
*
* Copyright 2019 SAP SE
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You should have received a copy of the License along with this
* program. If not, you may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

package helm

import (
	"errors"
	"strings"

	"github.com/Masterminds/semver"
	"k8s.io/helm/pkg/proto/hapi/chart"
)

func stringSliceContains(stringSlice []string, searchString string) bool {
	searchString = normalizeString(searchString)
	for _, s := range stringSlice {
		s = normalizeString(s)
		if strings.Contains(s, searchString) || strings.Contains(searchString, s) {
			return true
		}
	}
	return false
}

func normalizeRepoName(repoURL string) string {
	name := strings.TrimPrefix(repoURL, "https://")
	name = strings.TrimSuffix(name, "/")
	name = strings.ReplaceAll(name, "/", "-")
	return strings.ReplaceAll(name, ".", "-")
}

func normalizeString(theString string) string {
	theString = strings.TrimSpace(theString)
	return strings.ToLower(theString)
}

func getChartVersion(c *chart.Chart) (*semver.Version, error) {
	m := c.GetMetadata()
	if m == nil {
		return nil, errors.New("chart has no metdata")
	}

	v := m.GetVersion()
	if v == "" {
		return nil, errors.New("chart has no version")
	}

	return semver.NewVersion(v)
}
