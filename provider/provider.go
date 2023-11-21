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

package provider

import (
	"errors"
	"os"

	"github.com/bank-vaults/secret-init/model"
	"github.com/bank-vaults/secret-init/provider/file"
)

func New(providerName string) (model.Provider, error) {
	switch providerName {
	case file.ProviderName:
		provider, err := file.NewProvider(os.Getenv("SECRETS_FILE_PATH"))
		if err != nil {
			return nil, err
		}
		return provider, nil

	default:
		return nil, errors.New("invalid provider specified")
	}
}
