package scw

import (
	"net/http"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/auth"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

const (
	testEndpoint    = "https://api.example.com/"
	defaultEndpoint = "https://api.scaleway.com"
	testAccessKey   = "some access key"
	testSecretKey   = "some secret key"
)

func TestNewClientWithDefaults(t *testing.T) {

	options := []ClientOption{
		WithoutAuth(),
		WithInsecure(),
	}

	client, err := NewClient(options...)
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, defaultEndpoint, client.baseUrl)
	testhelpers.Equals(t, auth.NewNoAuth(), client.auth)

}

func TestNewClientWithOptions(t *testing.T) {

	someHTTPClient := &http.Client{}

	options := []ClientOption{
		WithEndpoint(testEndpoint),
		WithAuth(testAccessKey, testSecretKey),
		WithHttpClient(someHTTPClient),
	}

	client, err := NewClient(options...)
	testhelpers.Ok(t, err)

	testhelpers.Equals(t, testEndpoint, client.baseUrl)
	testhelpers.Equals(t, auth.NewToken(testAccessKey, testSecretKey), client.auth)

	testhelpers.Equals(t, someHTTPClient, client.httpClient)

}