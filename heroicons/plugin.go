package heroicons

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/go-app-blazar/heroicons-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// Config is the configuration for the HeroIcons plugin.
type Config struct {
	Location string
}

// plugin is the implementation of the HeroIcons plugin.
type plugin struct {
	config Config
}

var _ blazarplugin.Plugin = (*plugin)(nil)

// NewPlugin creates a new HeroIcons plugin.
func NewPlugin(config Config) blazarplugin.Plugin {
	if config.Location == "" {
		config.Location = "/web/heroicons/"
	}
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
	slog.InfoContext(context.TODO(), "Registering HeroIcons plugin", "location", location)
	mux.Handle(location, http.StripPrefix(location, p.httpHandler()))

	handler.Styles = append(handler.Styles, p.cssFilenames(p.config.Location)...)

	var fontFiles []string
	{
		files, err := fs.Glob(heroicons.FS, "*.woff2")
		if err != nil {
			slog.ErrorContext(context.TODO(), "Error getting HeroIcons font files", "error", err)
		} else {
			for _, file := range files {
				fontFiles = append(fontFiles, filepath.Join(p.config.Location, file))
			}
		}
	}
	handler.Fonts = append(handler.Fonts, fontFiles...)
	handler.CacheableResources = append(handler.CacheableResources, fontFiles...)
}

// cssFilenames returns the CSS filenames for the plugin.
func (p *plugin) cssFilenames(prefix string) []string {
	var cssFiles []string
	{
		files, err := fs.Glob(heroicons.FS, "*.css")
		if err != nil {
			slog.ErrorContext(context.TODO(), "Error getting HeroIcons CSS files", "error", err)
		} else {
			cssFiles = append(cssFiles, files...)
		}
	}

	slices.Sort(cssFiles)
	for i, filename := range cssFiles {
		cssFiles[i] = filepath.Join(prefix, filename)
	}
	return cssFiles
}

// httpHandler returns the HTTP handler for the plugin.
func (p *plugin) httpHandler() http.Handler {
	return blazarplugin.MimeTypeHandler(http.FileServerFS(heroicons.FS), blazarplugin.DefaultMimeTypeExtensions())
}
