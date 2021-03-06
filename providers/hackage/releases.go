//
// Copyright 2016-2017 Bryan T. Meyers <bmeyers@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package hackage

import (
	"github.com/DataDrake/cuppa/results"
)

// Releases is a representation of one or more Hackage releases
type Releases struct {
	Releases []Release
}

// Convert turns a Hackage result set into a Cuppa result set
func (hrs *Releases) Convert(name string) *results.ResultSet {
	rs := results.NewResultSet(name)
	for _, rel := range hrs.Releases {
		r := rel.Convert()
		if r != nil {
			rs.AddResult(r)
		}
	}
	return rs
}
