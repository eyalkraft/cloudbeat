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

package fetchers

import (
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/cloudbeat/resources/fetching"
	"github.com/elastic/cloudbeat/resources/manager"
	k8s "k8s.io/client-go/kubernetes"
)

type KubeFactory struct {
}

func init() {
	manager.Factories.ListFetcherFactory(fetching.KubeAPIType, &KubeFactory{})
}

func (f *KubeFactory) Create(c *common.Config) (fetching.Fetcher, error) {
	cfg := KubeApiFetcherConfig{}
	err := c.Unpack(&cfg)
	if err != nil {
		return nil, err
	}
	return f.CreateFrom(cfg, kubernetes.GetKubernetesClient)
}

func (f *KubeFactory) CreateFrom(cfg KubeApiFetcherConfig, provider func(string, kubernetes.KubeClientOptions) (k8s.Interface, error)) (fetching.Fetcher, error) {
	fe := &KubeFetcher{
		cfg:            cfg,
		clientProvider: provider,
		watchers:       make([]kubernetes.Watcher, 0),
	}

	logp.L().Infof("Kube Fetcher created with the following config: Name: %s, Interval: %s, "+
		"Kubeconfig: %s", cfg.Name, cfg.Interval, cfg.Kubeconfig)
	return fe, nil
}
