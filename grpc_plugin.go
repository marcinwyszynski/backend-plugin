package backendplugin

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/marcinwyszynski/backendplugin/proto"
)

type GRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin

	logger hclog.Logger

	// Concrete implementation, written in Go.
	Impl BackendPlugin
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterBackendServer(s, &GRPCServer{Impl: p.Impl, logger: p.logger})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewBackendClient(c)}, nil
}
