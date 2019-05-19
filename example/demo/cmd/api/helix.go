package main

import (
	"net/http"
	"strings"

	"github.com/nicklaw5/helix"
)

// getHelixClient returns a helix client that can be mocked
func getHelixClient() (*helix.Client, error) {
	httpClient := &twitchHTTPClient{}
	return helix.NewClient(&helix.Options{
		ClientID:   twitchClientID,
		HTTPClient: httpClient,
	})
}

type twitchHTTPClient struct {
}

// Do intercepts HTTP requests used by helix client.
// To allow integration tests of the helix client (which has the API url hardcoded...)
// we replace scheme+host with API url injected via env. variables
func (t *twitchHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if strings.Index(helix.APIBaseURL, req.URL.Host) != -1 {
		req.URL.Scheme = twitchAuthBaseURL.Scheme
		req.URL.Host = twitchAPIBaseURL.Host
	} else if strings.Index(helix.AuthBaseURL, req.URL.Host) != -1 {
		req.URL.Scheme = twitchAuthBaseURL.Scheme
		req.URL.Host = twitchAuthBaseURL.Host
	}

	return http.DefaultClient.Do(req)
}
