package htmlevent

import (
	"fmt"

	"github.com/go-app-blazar/appvalue"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// MustParse parses the app.Event into the target struct and panics if there is an error.
func MustParse(appEvent app.Event, target any) {
	err := Parse(appEvent, target)
	if err != nil {
		panic(fmt.Errorf("htmlevent: MustParse: %w", err))
	}
}

// Parse parses the app.Event into the target struct and returns an error if there is an error.
func Parse(appEvent app.Event, target any) error {
	return appvalue.Unmarshal(appEvent, target)
}
