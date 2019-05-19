package integrationtests

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/bosonic-code/mock-api/mocker"
)

var (
	mo *mocker.Client
)

func init() {
	var err error

	if mo, err = mocker.Create("twitch-mock:9999"); err != nil {
		log.Fatalf("Unable to create mock client %v", err)
	}
}

// TestGetStreamerStatus verifies that when we
// get the streamer status for a user - it returns
// true if the user has an ongoing stream on twitch

func TestGetStreamerStatus(t *testing.T) {
	var (
		// This describes  what  we expect the request to
		// twitch to look like
		request = mocker.Request{
			Method: http.MethodGet,
			Path:   "/helix/streams",
			Query: map[string]string{
				"type":    "all",
				"user_id": "1",
			},
		}

		// This is how we expect twitch to respond
		response = mocker.Response{
			Status: http.StatusOK,
			Body: `
		{
			"data": [{
				"id": "26007494656",
				"user_id": "1",
				"user_name": "LIRIK",
				"game_id": "417752",
				"community_ids": [
					"5181e78f-2280-42a6-873d-758e25a7c313",
					"848d95be-90b3-44a5-b143-6e373754c382",
					"fd0eab99-832a-4d7e-8cc0-04d73deb2e54"
				],
				"type": "live",
				"title": "Hey Guys, It's Monday - Twitter: @Lirik",
				"viewer_count": 32575,
				"started_at": "2017-08-14T16:08:32Z",
				"language": "en",
				"thumbnail_url": "https://static-cdn.jtvnw.net/previews-ttv/live_user_lirik-{width}x{height}.jpg",
				"tag_ids": [
					"6ea6bca4-4712-4ab9-a906-e3336a9d8039"
				]
			}],
			"pagination": {
				"cursor": "eyJiIjpudWxsLCJhIjp7Ik9mZnNldCI6MjB9fQ=="
			}
		}`,
		}
	)

	if err := mo.Handle(request, response); err != nil {
		log.Fatalf("Error setting up handle %v", err)
	}

	resp, err := http.Get("http://api:9999/users/1/is-streaming")
	if err != nil {
		t.Fatalf("Error when getting status %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Invalid status code %v", resp.StatusCode)
	}

	type isStreamingReponse struct {
		IsStreaming bool `json:"isStreaming"`
	}

	var apiResponse isStreamingReponse
	decoder := json.NewDecoder(resp.Body)

	if err = decoder.Decode(&apiResponse); err != nil {
		t.Fatalf("Unable to decode response %v", err)
	}

	if apiResponse.IsStreaming == false {
		t.Errorf("Invalid response %v, expected %v", false, true)
	}
}
