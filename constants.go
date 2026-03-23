package katarive

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "KATARIVE_VERSION",
	MagicCookieValue: "v1",
}

var PluginMap = map[string]plugin.Plugin{
	"source":  &SourcePlugin{},
	"speaker": &SpeakerPlugin{},
}
