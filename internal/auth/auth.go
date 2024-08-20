package auth

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type PcbeAccessTokenProvider struct {
	token string
}

func (p *PcbeAccessTokenProvider) GetAuthorizationToken(
	ctx context.Context, url *url.URL,
	additionalAuthenticationContext map[string]interface{},
) (string, error) {
	return p.token, nil
}

func (p *PcbeAccessTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}

func NewPcbeAccessTokenProvider(token string) *PcbeAccessTokenProvider {
	p := PcbeAccessTokenProvider{
		token: token,
	}

	return &p
}
