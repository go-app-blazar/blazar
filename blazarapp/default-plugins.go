package blazarapp

import (
	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/go-app-blazar/blazar/lucide"
)

// DefaultPlugins returns the default plugins for the Blazar app.
//
// This includes:
// - Lucide
// - Blazar
func DefaultPlugins() []blazarplugin.Plugin {
	var plugins []blazarplugin.Plugin
	plugins = append(plugins, lucide.NewPlugin(lucide.Config{
		Location: "/web/lucide/",
	}))
	plugins = append(plugins, blazar.NewPlugin(blazar.Config{
		Location: "/web/blazar/",
	}))
	return plugins
}
