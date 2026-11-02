package blazarapp

import (
	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/go-app-blazar/blazar/heroicons"
)

// DefaultPlugins returns the default plugins for the Blazar app.
//
// This includes:
// - HeroIcons
// - Blazar
func DefaultPlugins() []blazarplugin.Plugin {
	var plugins []blazarplugin.Plugin
	plugins = append(plugins, heroicons.NewPlugin(heroicons.Config{
		Location: "/web/heroicons/",
	}))
	plugins = append(plugins, blazar.NewPlugin(blazar.Config{
		Location: "/web/blazar/",
	}))
	return plugins
}
