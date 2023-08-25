//
// Copyright (c) 2021 Red Hat, Inc.
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

package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/redhat-appstudio/service-provider-integration-operator/pkg/serviceprovider"

	"github.com/google/go-github/v45/github"
	"github.com/redhat-appstudio/service-provider-integration-operator/pkg/spi-shared/tokenstorage"
	"golang.org/x/oauth2"
)

type githubClientBuilder struct {
	httpClient   *http.Client
	tokenStorage tokenstorage.TokenStorage
}

var _ serviceprovider.AuthorizedClientBuilder[github.Client] = (*githubClientBuilder)(nil)

var tokenNilError = errors.New("token used to construct authorized client is nil")

func (g *githubClientBuilder) CreateAuthorizedClient(ctx context.Context, token *oauth2.Token) (*github.Client, error) {
	if token == nil {
		return nil, tokenNilError
	}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, g.httpClient)
	ts := oauth2.StaticTokenSource(token)
	return github.NewClient(oauth2.NewClient(ctx, ts)), nil
}
