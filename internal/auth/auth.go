package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type Credentials struct {
	URL    string
	ID     string
	Secret string
}

type token struct {
	AccessToken string `json:"access_token"`
}

type Config struct {
	HTTPClient *http.Client
}

// WithHTTPClient allows optionally passing in a custom http client
// to the GetToken function, eg:
// c := &http.Client{}
// GetToken(ctx, "https://www.example.com", WithHTTPClient(c))
func WithHTTPClient(client *http.Client) Config {
	return Config{HTTPClient: client}
}

func GetToken(
	ctx context.Context,
	creds Credentials,
	configs ...Config,
) (string, error) {
	var token token

	client := &http.Client{Timeout: 10 * time.Second}

	for _, config := range configs {
		if config.HTTPClient != nil {
			client = config.HTTPClient
		}
	}

	grant := fmt.Sprintf(
		"grant_type=client_credentials&client_id=%s&client_secret=%s",
		creds.ID,
		creds.Secret,
	)

	req, err := http.NewRequest("POST", creds.URL, bytes.NewBuffer([]byte(grant)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get token (http code: " + resp.Status + ")")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}

	if token.AccessToken == "" {
		return "", errors.New("failed to get token: empty access token")
	}

	tflog.Debug(ctx, "token ok")

	return token.AccessToken, nil
}

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
