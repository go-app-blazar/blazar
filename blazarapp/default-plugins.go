package blazarapp

import (
	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/go-app-blazar/blazar/fontawesome"
)

// DefaultPlugins returns the default plugins for the Blazar app.
//
// This includes:
// - FontAwesome
// - Blazar
func DefaultPlugins() []blazarplugin.Plugin {
	var plugins []blazarplugin.Plugin
	plugins = append(plugins, fontawesome.NewPlugin(fontawesome.Config{
		Location: "/web/fontawesome/",
		Minify:   false,
	}))
	plugins = append(plugins, blazar.NewPlugin(blazar.Config{
		Location: "/web/blazar/",
	}))
	return plugins
}
