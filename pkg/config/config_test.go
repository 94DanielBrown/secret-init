// Copyright © 2023 Bank-Vaults Maintainers
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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name       string
		env        map[string]string
		wantConfig *Config
	}{
		{
			name: "Valid configuration",
			env: map[string]string{
				defaultEnvPrefix + LogLevelEnv:  "debug",
				defaultEnvPrefix + JSONLogEnv:   "true",
				defaultEnvPrefix + LogServerEnv: "",
				defaultEnvPrefix + DaemonEnv:    "true",
				defaultEnvPrefix + ProviderEnv:  "vault",
			},
			wantConfig: &Config{
				LogLevel:  "debug",
				JSONLog:   true,
				LogServer: "",
				Daemon:    true,
				Provider:  "vault",
			},
		},
	}

	for _, tt := range tests {
		ttp := tt
		t.Run(ttp.name, func(t *testing.T) {
			for envKey, envVal := range ttp.env {
				os.Setenv(envKey, envVal)
			}

			config, err := NewConfig()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			assert.Equal(t, ttp.wantConfig, config, "Unexpected config")

			// unset envs for the next test
			for envKey := range ttp.env {
				os.Unsetenv(envKey)
			}
		})
	}
}
