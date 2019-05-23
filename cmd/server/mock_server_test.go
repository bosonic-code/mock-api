package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/bosonic-code/mock-api/internal/proto"
)

// TestMatchers runs sets up a single handler in a server
// and executes a http.Request. After each run, we exacming
// header and body in the response.
// The goal is to verify that each supported mathcing method
// works as expected.
func TestMatchers(t *testing.T) {

	type sample struct {
		// name of sample (so we can easily identify failed samples)
		name string
		// The handler to initialize the server w/
		handler *proto.AddHandlerRequest
		// The request we want to test
		r *http.Request
		// Asserts for response
		expectedHeader int
		expectedBody   []byte
	}

	samples := []sample{
		sample{
			name: "Success case - match path",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/success",
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
					Body:   "Hello World",
				},
			},
			expectedHeader: http.StatusOK,
			expectedBody:   []byte("Hello World"),
			r: &http.Request{
				URL: &url.URL{
					Path: "/success",
				},
			},
		},

		sample{
			name: "Fail case - match path",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/success",
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
					Body:   "Hello World",
				},
			},
			expectedHeader: http.StatusNotFound,
			expectedBody:   nil,
			r: &http.Request{
				URL: &url.URL{
					Path: "/fail",
				},
			},
		},

		sample{
			name: "Success case - match path + query",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/success",
					Query: map[string]string{
						"a": "b",
					},
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
					Body:   "Hello World",
				},
			},
			expectedHeader: http.StatusOK,
			expectedBody:   []byte("Hello World"),
			r: &http.Request{
				URL: &url.URL{
					Path:     "/success",
					RawQuery: "a=b",
				},
			},
		},

		sample{
			name: "Fail case - match path + query",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/success",
					Query: map[string]string{
						"a": "b",
					},
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
					Body:   "Hello World",
				},
			},
			expectedHeader: http.StatusNotFound,
			expectedBody:   nil,
			r: &http.Request{
				URL: &url.URL{
					Path:     "/success",
					RawQuery: "a=c",
				},
			},
		},

		sample{
			name: "Success case - match path + headers",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/success",
					Headers: map[string]*proto.HeaderValue{
						"a": &proto.HeaderValue{Value: []string{"b", "c"}},
					},
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
					Body:   "Hello World",
				},
			},
			expectedHeader: http.StatusOK,
			expectedBody:   []byte("Hello World"),
			r: &http.Request{
				URL: &url.URL{
					Path: "/success",
				},
				Header: map[string][]string{
					"a": []string{"b", "c"},
				},
			},
		},

		sample{
			name: "Fail case - match path + headers",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/success",
					Headers: map[string]*proto.HeaderValue{
						"a": &proto.HeaderValue{Value: []string{"b", "c"}},
					},
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
					Body:   "Hello World",
				},
			},
			expectedHeader: http.StatusNotFound,
			expectedBody:   nil,
			r: &http.Request{
				URL: &url.URL{
					Path: "/success",
				},
				Header: map[string][]string{
					"a": []string{"b", "d"},
				},
			},
		},

		sample{
			name: "Success case - match path + http Method",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path:   "/update",
					Method: http.MethodPut,
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
				},
			},
			expectedHeader: http.StatusOK,
			r: &http.Request{
				URL: &url.URL{
					Path: "/update",
				},
				Method: http.MethodPut,
			},
		},

		sample{
			name: "Failure case - match path + http method",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path:   "/update",
					Method: http.MethodPut,
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusOK,
				},
			},
			expectedHeader: http.StatusNotFound,
			r: &http.Request{
				URL: &url.URL{
					Path: "/update",
				},
				Method: http.MethodPost,
			},
		},

		sample{
			name: "Success case - match path + payload ",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/create",
					Body: `{"number":1}`,
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusCreated,
					Body:   `{"result":"good"}`,
				},
			},
			expectedHeader: http.StatusCreated,
			expectedBody:   []byte(`{"result":"good"}`),
			r: &http.Request{
				URL: &url.URL{
					Path: "/create",
				},
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"number":1}`))),
			},
		},

		sample{
			name: "Fail case - match path + payload",
			handler: &proto.AddHandlerRequest{
				RequestMatcher: &proto.RequestMatcher{
					Path: "/create",
					Body: `{"number":1}`,
				},
				Response: &proto.MatcherResponse{
					Status: http.StatusCreated,
					Body:   `{"result":"good"}`,
				},
			},
			expectedHeader: http.StatusNotFound,
			r: &http.Request{
				URL: &url.URL{
					Path: "/create",
				},
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"number":2}`))),
			},
		},
	}

	for _, sample := range samples {

		// Initialize server and add the handler
		server := &MockServer{}

		if _, err := server.AddHandler(context.TODO(), sample.handler); err != nil {
			t.Fatalf("\"%v\": Failed  to handle input %v", sample.name, err)
		}

		// Send the sample http request
		// whith a recorder so we can inspect the response
		rr := httptest.NewRecorder()
		server.ServeHTTP(rr, sample.r)

		// Verify expected rheader
		if rr.Code != sample.expectedHeader {
			t.Errorf("\"%v\": Invalid response code %v, expected %v",
				sample.name,
				rr.Code,
				sample.expectedHeader)
		}

		var (
			err  error
			body []byte
		)

		if body, err = ioutil.ReadAll(rr.Body); err != nil {
			t.Fatalf("\"%v\": Failed to read response body %v", sample.name, err)
		}

		// Verify expected body
		if bytes.Equal(body, sample.expectedBody) == false {
			t.Errorf("\"%v\": Invalid response body %v, expected %v",
				sample.name,
				string(body),
				string(sample.expectedBody))
		}
	}
}
