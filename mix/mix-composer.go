package mix

import "github.com/maxence-charriere/go-app/v11/pkg/app"

// MixComposer is a mixin for components that need to implement app.MixComposer.
//
// This is read by the parser and will append the app.MixComposer interface to the interface that gets generated.
type MixComposer struct {
	app.Composer
}
