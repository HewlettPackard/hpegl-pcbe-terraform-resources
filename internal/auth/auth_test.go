package auth

import (
	"context"
	"net/http"
	"testing"

	"github.com/h2non/gock"
)

func TestAuthOk(t *testing.T) {
	mockURL := "http://example.com/api"
	mockBody := `grant_type=client_credentials&client_id=id123&client_secret=secret123`
	expectedToken := "token123"

	defer gock.Off()

	gock.New(mockURL).
		Post("/api").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		JSON(mockBody).
		Reply(200).
		JSON(map[string]string{"access_token": expectedToken})

	creds := Credentials{
		URL:    mockURL,
		ID:     "id123",
		Secret: "secret123",
	}

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	token, err := GetToken(context.Background(), creds, WithHTTPClient(client))
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}

	if token != expectedToken {
		t.Fatalf("GetToken returned unexpected value: %v", token)
	}
}

func TestAuthBad(t *testing.T) {
	mockURL := "http://example.com/api"
	mockBody := `grant_type=client_credentials&client_id=id123&client_secret=secret123`

	defer gock.Off()

	gock.New(mockURL).
		Post("/api").
		MatchHeader("Content-Type", "application/x-www-form-urlencoded").
		JSON(mockBody).
		Reply(401).
		BodyString("invalid")

	creds := Credentials{
		URL:    mockURL,
		ID:     "id123",
		Secret: "secret123",
	}

	client := &http.Client{Transport: &http.Transport{}}
	gock.InterceptClient(client)

	_, err := GetToken(context.Background(), creds, WithHTTPClient(client))
	if err == nil {
		t.Fatalf("GetToken should have failed")
	}

	if err.Error() != "failed to get token (http code: 401 Unauthorized)" {
		t.Fatalf("GetToken returned unexpected error: %v", err)
	}
}
