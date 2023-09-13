package backendplugin

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "BACKEND_PLUGIN",
	MagicCookieValue: "i-solemnly-swear-i-am-up-to-no-good",
}
