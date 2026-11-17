package blazar

import "github.com/maxence-charriere/go-app/v11/pkg/app"

// MixComposer is a mixin for components that need to implement app.Composer.
type MixComposer struct {
	app.Composer
}
