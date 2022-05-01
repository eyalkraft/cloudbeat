// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package evaluator

import "context"

type Evaluator interface {
	Decision(context.Context, interface{}) (interface{}, error)
	Stop(context.Context)
	Decode(result interface{}) ([]Finding, error)
}

type Metadata struct {
	Version string `json:"opa_version"`
}

type RuleResult struct {
	Findings []Finding `json:"findings"`
	Metadata Metadata  `json:"metadata"`
	// Golang 1.18 will introduce generics which will be useful for typing the resource field
	Resource interface{} `json:"resource"`
}

type Finding struct {
	Result Result `json:"result"`
	Rule   Rule   `json:"rule"`
}

type Result struct {
	Evaluation string      `json:"evaluation"`
	Expected   interface{} `json:"expected"`
	Evidence   interface{} `json:"evidence"`
}

type Rule struct {
	Id                    string    `json:"id"`
	Name                  string    `json:"name"`
	Profile_Applicability string    `json:"profile_applicability"`
	Description           string    `json:"description"`
	Rationale             string    `json:"rationale"`
	Audit                 string    `json:"audit"`
	Remediation           string    `json:"remediation"`
	Impact                string    `json:"impact"`
	Default_Value         string    `json:"default_value"`
	References            string    `json:"references"`
	Section               string    `json:"section"`
	Version               string    `json:"version"`
	Tags                  []string  `json:"tags"`
	Benchmark             Benchmark `json:"benchmark"`
}

type Benchmark struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
