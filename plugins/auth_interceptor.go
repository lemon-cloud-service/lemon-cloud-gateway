package plugins

import (
	_ "github.com/micro/cli/v2"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/plugin"
)

func RegisterAuthPlugin() {
	plugin.Register(cors.NewPlugin())
}
