// Copyright 2020 Burak Sezer
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

package lib

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-discover"
	"github.com/hashicorp/go-discover/provider/k8s"
	"github.com/mitchellh/mapstructure"
)

type CloudDiscovery struct {
	config   *Config
	log      *log.Logger
	discover *discover.Discover
}

type Config struct {
	Provider string
	Args     interface{}
}

func (c *CloudDiscovery) checkErrors() error {
	if c.config == nil {
		return fmt.Errorf("config cannot be nil")
	}
	if c.log == nil {
		return fmt.Errorf("logger cannot be nil")
	}
	if c.config.Provider != "k8s" {
		_, ok := discover.Providers[c.config.Provider]
		if !ok {
			return fmt.Errorf("invalid provider: %s", c.config.Provider)
		}
	}
	return nil
}

func (c *CloudDiscovery) Initialize() error {
	if err := c.checkErrors(); err != nil {
		return err
	}

	m := map[string]discover.Provider{}
	if c.config.Provider == "k8s" {
		m[c.config.Provider] = &k8s.Provider{}
	} else {
		provider, _ := discover.Providers[c.config.Provider]
		m[c.config.Provider] = provider
	}

	opt := discover.WithProviders(m)
	d, err := discover.New(opt)
	if err != nil {
		return fmt.Errorf("discover.New returned an error: %w", err)
	}
	c.discover = d
	c.log.Printf("[INFO] Service discovery plugin is enabled, provider: %s", c.config.Provider)
	return nil
}

func (c *CloudDiscovery) SetLogger(l *log.Logger) {
	c.log = l
}

func (c *CloudDiscovery) SetConfig(cfg map[string]interface{}) error {
	var cg Config
	err := mapstructure.Decode(cfg, &cg)
	if err != nil {
		return err
	}
	c.config = &cg
	return nil
}

func (c *CloudDiscovery) getArgs() string {
	result := fmt.Sprintf("provider=%s", c.config.Provider)

	args, ok := c.config.Args.(string)
	if ok {
		return fmt.Sprintf("%s %s", result, args)
	}

	for key, value := range c.config.Args.(map[string]string) {
		result += fmt.Sprintf("%s=%s", key, value)
	}
	return result
}

func (c *CloudDiscovery) DiscoverPeers() ([]string, error) {
	peers, err := c.discover.Addrs(c.getArgs(), c.log)
	if err != nil {
		return nil, err
	}
	if len(peers) == 0 {
		return nil, fmt.Errorf("no peer found")
	}
	return peers, nil
}

func (c *CloudDiscovery) Register() error { return nil }

func (c *CloudDiscovery) Deregister() error { return nil }

func (c *CloudDiscovery) Close() error { return nil }
