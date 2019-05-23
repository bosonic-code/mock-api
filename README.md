
```
##################### 
#      WARNING      # 
#   EARLY STAGE OF  #
#     DEVELOPMENT   #
#     DO NOT EAT    #
#####################
```
# mock-api 

The purpose of Mock API is to act as a small and quick way of mocking external REST-apis for use in integration tests. 

Available on docker hub:
https://hub.docker.com/r/bosonic/mock-api

## Sample usage  (docker compose)

NOTE: This sample usage can be found in the examples: https://github.com/bosonic-code/mock-api/tree/master/example/demo. 

In this example, we have an API that checks if a twitch user is currently streaming by using the github.com/nicklaw5/helix twitch API client. 


### Configuring the docker-compose suite
We start by adding our subject container to the docker-compose suite:


https://github.com/bosonic-code/mock-api/blob/master/example/demo/integration-tests/docker-compose.yml
``` 
version: '3'

services:

  # Services
  api:
    container_name: test_api
    image: yourcompany/api
    env_file: ./test.env
```

And  you create an integration test program that calls your api

https://github.com/bosonic-code/mock-api/blob/master/example/demo/integration-tests/docker-compose.yml
```
version: '3'

services:

 # Image running integration tests for API
 integration-test:
    container_name: test_integration-test
    depends_on:
      - "api"

  # Services
  api:
    container_name: test_api
    image: yourcompany/api

```

You  can  add bosonic/mock-api into the docker-compose suite of services, and point yourcompany/api to send traffic to your mockserver. 

What  mock-api allows you to do is run as a service in your docker-compose suite, allowing your integration-tests to configure how it should respond to various requests.

https://github.com/bosonic-code/mock-api/blob/master/example/demo/integration-tests/docker-compose.yml
```


 # Image running integration tests for API
 integration-test:
    container_name: test_integration-test
    depends_on:
      - "api"
    env_file: ./test.env
    depends_on:
      - "twitch-mock"

  # Services
  api:
    container_name: test_api
    image: yourcompany/api
    depends_on:
      - "twitch-mock"
    env_file: ./test.env

  # Twitch API mock
  twitch-mock:
    container_name: test_twitch-mock
    image: bosonic/mock-api
    env_file: ./mock-api.env

```

To finish the setup in docker-compose - the api needs to be configured to route requests to the twitch-mock container instead of the original endpoint. This can be achieved by using e.g. environment variables:

https://github.com/bosonic-code/mock-api/blob/master/example/demo/integration-tests/test.env
 ```

    TWITCH_API_BASE_URL=http://twitch-mock:1337/
    TWITCH_AUTH_BASE_URL=http://twitch-mock:1337/

 ``` 

### Writing a test 
 Once the traffic has been routed to the mock-server in your integration test suite, you can configure the mock server to respond in a predictable, testable mannner. We can run a quick test to verify that our service returns the  correct result when a user has any active streams. 

 We start by innitializing a client that we can use to configure the mock server:

https://github.com/bosonic-code/mock-api/blob/master/example/demo/integration-tests/streamcheck_test.go)
```

var (
	mo mocker.MockerClient
)

func init() {
	var err error

	if mo, err = mocker.NewClient("twitch-mock:9999"); err != nil {
		log.Fatalf("Unable to create mock client %v", err)
	}
}

```

We can then configure a scenario in the server:

https://github.com/bosonic-code/mock-api/blob/master/example/demo/integration-tests/streamcheck_test.go)
```

// TestGetStreamerStatus verifies that when we
// get the streamer status for a user - it returns
// true if the user has an ongoing stream on twitch

func TestGetStreamerStatus(t *testing.T) {
	var (
	  // This describes  what  we expect the request to
		// twitch to look like
		request = mocker.RequestMatcher{
			Method: http.MethodGet,
			Path:   "/helix/streams",
			Query: map[string]string{
				"type":    "all",
				"user_id": "1",
			},
		}

		// This is how we expect twitch to respond
		response = mocker.MatcherResponse{
			Status: http.StatusOK,
			Body: `
		{
			"data": [{
				"id": "26007494656",
				"user_id": "1",
				...
			}],
            ...
		}`,
		}
	)



	// Register the scenario in the server
	addHandlerRequest := &mocker.AddHandlerRequest{
		RequestMatcher: &request,
		Response:       &response}

	if _, err := mo.AddHandler(
		context.TODO(),
		addHandlerRequest,
	); err != nil {
		log.Fatalf("Error setting up handle %v", err)
	}

    ...
```

We can the n run our tests and verify that the API behaves correctly:

https://github.com/bosonic-code/mock-api/blob/master/example/demo/integration-tests/streamcheck_test.go)
```

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

```

## Configuring responses

Currently - you can configure the server to match requests by Method, Path, Headers and Query. See our tests for example usage: https://github.com/bosonic-code/mock-api/blob/master/cmd/server/mock_server_test.go

