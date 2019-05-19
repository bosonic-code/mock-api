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

func TestMatch(t *testing.T) {
	type sample struct {
		in             *proto.HandleRequest
		expectedHeader int
		expectedBody   []byte
		r              *http.Request
	}
	samples := []sample{
		// Success case - match path
		sample{
			in: &proto.HandleRequest{
				Request: &proto.Request{
					Path: "/success",
				},
				Response: &proto.Response{
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

		// Fail case - match path
		sample{
			in: &proto.HandleRequest{
				Request: &proto.Request{
					Path: "/success",
				},
				Response: &proto.Response{
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

		// Success case - match path + query
		sample{
			in: &proto.HandleRequest{
				Request: &proto.Request{
					Path: "/success",
					Query: map[string]string{
						"a": "b",
					},
				},
				Response: &proto.Response{
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

		// Fail case - match path + query
		sample{
			in: &proto.HandleRequest{
				Request: &proto.Request{
					Path: "/success",
					Query: map[string]string{
						"a": "b",
					},
				},
				Response: &proto.Response{
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

		// Success case - match path + headers
		sample{
			in: &proto.HandleRequest{
				Request: &proto.Request{
					Path: "/success",
					Headers: map[string]*proto.HeaderValue{
						"a": &proto.HeaderValue{Value: []string{"b", "c"}},
					},
				},
				Response: &proto.Response{
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

		// Fail case - match path + headers
		sample{
			in: &proto.HandleRequest{
				Request: &proto.Request{
					Path: "/success",
					Headers: map[string]*proto.HeaderValue{
						"a": &proto.HeaderValue{Value: []string{"b", "c"}},
					},
				},
				Response: &proto.Response{
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
	}
	for i, sample := range samples {

		t.Logf("Running sample %v", i)

		server := &Server{}
		if _, err := server.Handle(context.TODO(), sample.in); err != nil {
			t.Fatalf("Failed to handle input %v", err)
		}

		rr := httptest.NewRecorder()
		server.ServeHTTP(rr, sample.r)

		if rr.Code != sample.expectedHeader {
			t.Errorf("Invalid response code %v, expected %v",
				rr.Code,
				sample.expectedHeader)
		}

		var (
			err  error
			body []byte
		)

		if body, err = ioutil.ReadAll(rr.Body); err != nil {
			t.Fatalf("Failed to read response body %v", err)
		}

		if bytes.Equal(body, sample.expectedBody) == false {
			t.Errorf("Invalid response body %v, expected %v", string(body), string(sample.expectedBody))
		}
	}
}
