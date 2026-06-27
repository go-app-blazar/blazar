package blazarplugin

import (
	"net/http"
	"path/filepath"
)

// DefaultMimeTypeExtensions returns the default MIME type extensions.
func DefaultMimeTypeExtensions() map[string]string {
	return map[string]string{
		".avif":  "image/avif",
		".css":   "text/css",
		".eot":   "application/vnd.ms-fontobject",
		".gif":   "image/gif",
		".heic":  "image/heic",
		".heif":  "image/heif",
		".ico":   "image/x-icon",
		".jpeg":  "image/jpeg",
		".jpg":   "image/jpeg",
		".js":    "text/javascript",
		".otf":   "font/otf",
		".png":   "image/png",
		".svg":   "image/svg+xml",
		".ttf":   "font/ttf",
		".webp":  "image/webp",
		".woff":  "font/woff",
		".woff2": "font/woff2",
	}
}

// MimeTypeHandler is a handler that sets the MIME type of the response based on the file extension.
func MimeTypeHandler(h http.Handler, fileExtensionMap map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		extension := filepath.Ext(r.URL.Path)
		mimeType, ok := fileExtensionMap[extension]
		if ok {
			w.Header().Set("Content-Type", mimeType)
		}
		h.ServeHTTP(w, r)
	})
}
