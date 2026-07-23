package blazar

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// ProgressBar creates a new progress bar.
func ProgressBar() *blazarProgressBar {
	return &blazarProgressBar{
		IHeight: "4px",
	}
}

// blazarProgressBar is the component that renders a progress bar.
type blazarProgressBar struct {
	app.Compo

	IAnimated      bool     // If true, the progress bar will be animated.
	IClasses       []string // The classes to add to the progress bar.
	IHeight        string   // The height of the progress bar.
	IIndeterminate bool     // If true, the progress bar will be indeterminate.
	IPercentage    uint     // The percentage of the progress bar.  This is ignored if IIndeterminate is true.
}

var _ app.Composer = (*blazarProgressBar)(nil)

// Animated enables or disables the animated mode of the progress bar.
func (c *blazarProgressBar) Animated(animated bool) *blazarProgressBar {
	c.IAnimated = animated
	return c
}

// Class adds a class to the progress bar.
func (c *blazarProgressBar) Class(class ...string) *blazarProgressBar {
	c.IClasses = class
	return c
}

// Height sets the height of the progress bar.
func (c *blazarProgressBar) Height(height string) *blazarProgressBar {
	c.IHeight = height
	return c
}

// Percentage sets the percentage of the progress bar.
func (c *blazarProgressBar) Percentage(percentage uint) *blazarProgressBar {
	c.IPercentage = percentage
	if c.IPercentage > 100 {
		c.IPercentage = 100
	}
	return c
}

// Indeterminate sets the indeterminate mode of the progress bar.
func (c *blazarProgressBar) Indeterminate(indeterminate bool) *blazarProgressBar {
	c.IIndeterminate = indeterminate
	return c
}

func (c *blazarProgressBar) Render() app.UI {
	var extraClass string
	if c.IIndeterminate {
		extraClass = " blazar-progress-bar__bar--indeterminate"
	} else if c.IAnimated {
		extraClass = " blazar-progress-bar__bar--animated"
	}

	return app.Div().
		Class(append([]string{"blazar-progress-bar"}, c.IClasses...)...).
		Style("--height", c.IHeight).
		Style("--percentage", fmt.Sprintf("%d", c.IPercentage)).
		Body(
			app.Div().
				Class("blazar-progress-bar__bar", extraClass),
		)
}
