package healthz

import (
	"context"

	"github.com/bufbuild/connect-go"

	"github.com/sivchari/chat-example/pkg/log"
	"github.com/sivchari/chat-example/proto/proto"
	"github.com/sivchari/chat-example/proto/proto/protoconnect"
)

type server struct {
	logger log.Handler
}

func New(logger log.Handler) protoconnect.HealthzHandler {
	return &server{
		logger: logger,
	}
}

func (s *server) Check(ctx context.Context, req *connect.Request[proto.CheckRequest]) (*connect.Response[proto.CheckResponse], error) {
	return nil, nil
}
