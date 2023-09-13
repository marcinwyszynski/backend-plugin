package backendplugin

import "github.com/hashicorp/go-plugin"

const BackendPluginName = "backend"

// PluginMap is the map of plugins we can dispense.
var Plugins = map[string]plugin.Plugin{
	BackendPluginName: &GRPCPlugin{},
}

// Server creates a plugin.ServeConfig for the given BackendPlugin.
func Server(impl BackendPlugin) *plugin.ServeConfig {
	return &plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			BackendPluginName: &GRPCPlugin{Impl: impl},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	}
}
