package fontawesome

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// Config is the configuration for the FontAwesome plugin.
type Config struct {
	Location string
	Minify   bool
}

// plugin is the implementation of the FontAwesome plugin.
type plugin struct {
	config Config
}

var _ blazarplugin.Plugin = (*plugin)(nil)

// NewPlugin creates a new FontAwesome plugin.
func NewPlugin(config Config) blazarplugin.Plugin {
	return &plugin{
		config: config,
	}
}

// Register registers the plugin against the go-app handler and the HTTP mux.
func (p *plugin) Register(handler *app.Handler, mux *http.ServeMux) {
	location := p.config.Location
	if handler.Resources != nil {
		location = handler.Resources.Resolve(location)
	}
	slog.InfoContext(context.TODO(), "Registering FontAwesome plugin", "location", location)
	mux.Handle(location, http.StripPrefix(location, p.httpHandler()))

	handler.Styles = append(handler.Styles, p.cssFilenames(p.config.Location)...)
	handler.CacheableResources = append(handler.CacheableResources,
		filepath.Join(p.config.Location, "webfonts/fa-brands-400.woff2"),
		filepath.Join(p.config.Location, "webfonts/fa-regular-400.woff2"),
		filepath.Join(p.config.Location, "webfonts/fa-solid-900.woff2"),
		filepath.Join(p.config.Location, "webfonts/fa-v4compatibility.woff2"),
	)
}

// cssFilenames returns the CSS filenames for the plugin.
func (p *plugin) cssFilenames(prefix string) []string {
	var cssFiles []string
	if p.config.Minify {
		cssFiles = append(cssFiles,
			"css/fontawesome.min.css",
			"css/brands.min.css",
			"css/regular.min.css",
			"css/solid.min.css",
		)
	} else {
		cssFiles = append(cssFiles,
			"css/fontawesome.css",
			"css/brands.css",
			"css/regular.css",
			"css/solid.css",
		)
	}
	slices.Sort(cssFiles)
	for i, filename := range cssFiles {
		cssFiles[i] = filepath.Join(prefix, filename)
	}
	return cssFiles
}

// httpHandler returns the HTTP handler for the plugin.
func (p *plugin) httpHandler() http.Handler {
	newFS, err := fs.Sub(embeddedFS, "embedded")
	if err != nil {
		return http.NotFoundHandler()
	}
	return http.FileServer(http.FS(newFS))
}
