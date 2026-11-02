package tabler

import (
	"context"
	"io/fs"
	"log/slog"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/go-app-blazar/blazar/blazarplugin"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// Config is the configuration for the Tabler plugin.
type Config struct {
	Location string
}

// plugin is the implementation of the Tabler plugin.
type plugin struct {
	config Config
}

var _ blazarplugin.Plugin = (*plugin)(nil)

// NewPlugin creates a new Tabler plugin.
func NewPlugin(config Config) blazarplugin.Plugin {
	if config.Location == "" {
		config.Location = "/web/tabler/"
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
	slog.InfoContext(context.TODO(), "Registering Tabler plugin", "location", location)
	mux.Handle(location, http.StripPrefix(location, p.httpHandler()))

	handler.Styles = append(handler.Styles, p.cssFilenames(p.config.Location)...)

	var fontFiles []string
	{
		files, err := fs.Glob(tabler.FS, "*.woff2")
		if err != nil {
			slog.ErrorContext(context.TODO(), "Error getting Tabler font files", "error", err)
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
		files, err := fs.Glob(tabler.FS, "*.css")
		if err != nil {
			slog.ErrorContext(context.TODO(), "Error getting Tabler CSS files", "error", err)
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
	return blazarplugin.MimeTypeHandler(http.FileServerFS(tabler.FS), blazarplugin.DefaultMimeTypeExtensions())
}
