package main

import (
	"context"
	"net/http"

	"github.com/bosonic-code/mock-api/internal/proto"
)

type Server struct {
	Handles []*proto.HandleRequest
}

func (s *Server) Handle(ctx context.Context, in *proto.HandleRequest) (*proto.HandleResponse, error) {
	if s.Handles == nil {
		s.Handles = make([]*proto.HandleRequest, 0)
	}

	s.Handles = append(s.Handles, in)
	return &proto.HandleResponse{}, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handle := range s.Handles {
		if match(r, handle.Request) == true {
			w.WriteHeader(int(handle.Response.Status))
			w.Write([]byte(handle.Response.Body))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func match(r *http.Request, req *proto.Request) bool {
	if req.Path != r.URL.Path {
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
