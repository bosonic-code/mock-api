package main

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/bosonic-code/mock-api/internal/proto"
)

type MockServer struct {
	handlers []*proto.AddHandlerRequest
}

func (s *MockServer) AddHandler(ctx context.Context, in *proto.AddHandlerRequest) (*proto.AddHandlerResponse, error) {
	if s.handlers == nil {
		s.handlers = make([]*proto.AddHandlerRequest, 0)
	}

	s.handlers = append(s.handlers, in)
	return &proto.AddHandlerResponse{}, nil
}

func (s *MockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handle := range s.handlers {
		if match(r, handle.RequestMatcher) == true {
			w.WriteHeader(int(handle.Response.Status))
			w.Write([]byte(handle.Response.Body))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func match(r *http.Request, req *proto.RequestMatcher) bool {
	if req.Path != r.URL.Path {
		return false
	}

	if len(req.Body) > 0 {
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			return false
		}

		if string(b) != req.Body {
			return false
		}
	}

	if len(req.Method) > 0 && req.Method != r.Method {
		return false
	}

	for k, v := range req.Query {
		if r.URL.Query().Get(k) != v {
			return false
		}
	}

	for k, v := range req.Headers {
		var (
			headers []string
			ok      bool
		)

		if headers, ok = r.Header[k]; !ok {
			return false
		}

		for _, h := range v.Value {
			if contains(headers, h) == false {
				return false
			}
		}
	}

	return true
}

func contains(a []string, b string) bool {
	for _, v := range a {
		if b == v {
			return true
		}
	}

	return false
}
