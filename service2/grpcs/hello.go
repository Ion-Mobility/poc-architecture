package grpcs

import (
	context "context"
	"errors"
	"fmt"
)

type Hello struct {
	UnimplementedHelloServer
}

func (h *Hello) SayHello(ctx context.Context, paylod *Unsecure) (*Response, error) {
	return &Response{
		Message: "Hello, i'm insecure from Service 2",
	}, nil
}

func (h *Hello) SayHelloSecure(ctx context.Context, payload *Secure) (*Response, error) {
	if payload.UserId == "" || payload.Email == "" {
		return nil, errors.New("invlaid user")
	}

	return &Response{
		Message: fmt.Sprintf("Hello %s from Service 2", payload.Email),
	}, nil
}
