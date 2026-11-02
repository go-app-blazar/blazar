package blazarapp

import (
	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/go-app-blazar/blazar/plugin/tabler"
)

// DefaultPlugins returns the default plugins for the Blazar app.
//
// This includes:
// - Tabler (icons)
// - Blazar
func DefaultPlugins() []blazarplugin.Plugin {
	var plugins []blazarplugin.Plugin
	plugins = append(plugins, tabler.NewPlugin(tabler.Config{
		Location: "/web/tabler/",
	}))
	plugins = append(plugins, blazar.NewPlugin(blazar.Config{
		Location: "/web/blazar/",
	}))
	return plugins
}
