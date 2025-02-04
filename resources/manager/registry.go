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

package manager

import (
	"context"
	"fmt"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/cloudbeat/resources/fetching"
)

type FetchersRegistry interface {
	Register(key string, f fetching.Fetcher, c ...fetching.Condition) error
	Keys() []string
	ShouldRun(key string) bool
	Run(ctx context.Context, key string) ([]fetching.Resource, error)
	Stop(ctx context.Context)
}

type fetchersRegistry struct {
	reg map[string]registeredFetcher
}

type registeredFetcher struct {
	f fetching.Fetcher
	c []fetching.Condition
}

func NewFetcherRegistry() FetchersRegistry {
	return &fetchersRegistry{
		reg: make(map[string]registeredFetcher),
	}
}

// Register registers a Fetcher implementation.
func (r *fetchersRegistry) Register(key string, f fetching.Fetcher, c ...fetching.Condition) error {
	if _, ok := r.reg[key]; ok {
		return fmt.Errorf("fetcher key collision: %q is already registered", key)
	}

	r.reg[key] = registeredFetcher{
		f: f,
		c: c,
	}

	return nil
}

func (r *fetchersRegistry) Keys() []string {
	keys := []string{}
	for k := range r.reg {
		keys = append(keys, k)
	}

	return keys
}

func (r *fetchersRegistry) ShouldRun(key string) bool {
	registered, ok := r.reg[key]
	if !ok {
		return false
	}

	for _, condition := range registered.c {
		if !condition.Condition() {
			logp.L().Infof("Conditional fetcher %q should not run because %q", key, condition.Name())
			return false
		}
	}

	return true
}

func (r *fetchersRegistry) Run(ctx context.Context, key string) ([]fetching.Resource, error) {
	registered, ok := r.reg[key]
	if !ok {
		return nil, fmt.Errorf("fetcher %v not found", key)
	}

	return registered.f.Fetch(ctx)
}

func (r *fetchersRegistry) Stop(ctx context.Context) {
	for key, rfetcher := range r.reg {
		rfetcher.f.Stop()
		logp.L().Infof("Fetcher for key %q stopped", key)
	}
}
