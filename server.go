package backendplugin

import (
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

const BackendPluginName = "backend"

// PluginMap is the map of plugins we can dispense.
var Plugins = map[string]plugin.Plugin{
	BackendPluginName: &GRPCPlugin{},
}

// Server creates a plugin.ServeConfig for the given BackendPlugin.
func Server(impl BackendPlugin) *plugin.ServeConfig {
	logger := hclog.New(&hclog.LoggerOptions{
		// We send all output to terraform. Go-plugin will take the output and
		// pass it through another hclog.Logger on the client side where it can
		// be filtered.
		Level:      hclog.Trace,
		JSONFormat: true,
	})
	log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))

	return &plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			BackendPluginName: &GRPCPlugin{Impl: impl},
		},
		GRPCServer: plugin.DefaultGRPCServer,
		Logger:     logger,
	}
}
