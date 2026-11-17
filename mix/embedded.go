package mix

import (
	"embed"
	"io/fs"
)

//go:embed embedded/*
var embeddedFS embed.FS

// FS is the embedded filesystem.
var FS, _ = fs.Sub(embeddedFS, "embedded")
