package blazar

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func StatusBar() *blazarStatusBar {
	return &blazarStatusBar{}
}

type blazarStatusBar struct {
	app.Compo

	TextValue string
	GoodValue *bool
}

var _ app.Composer = (*blazarStatusBar)(nil)

func (c *blazarStatusBar) Text(text string) *blazarStatusBar {
	c.TextValue = text
	return c
}

func (c *blazarStatusBar) Good() *blazarStatusBar {
	c.GoodValue = new(bool)
	*c.GoodValue = true
	return c
}

func (c *blazarStatusBar) Bad() *blazarStatusBar {
	c.GoodValue = new(bool)
	*c.GoodValue = false
	return c
}

func (c *blazarStatusBar) Neutral() *blazarStatusBar {
	c.GoodValue = nil
	return c
}

func (c *blazarStatusBar) Render() app.UI {
	goodClass := ""
	goodIcon := "circle-info"
	if c.GoodValue != nil {
		if *c.GoodValue {
			goodClass = "good"
			goodIcon = "circle-check"
		} else {
			goodClass = "bad"
			goodIcon = "circle-exclamation"
		}
	}

	return app.Div().
		Class("blazar-statusbar").
		Class(goodClass).
		Body(
			app.If(goodIcon != "", func() app.UI {
				return Icon().
					Icon(goodIcon)
			}),
			app.Span().
				Class("blazar-button-label").
				Text(c.TextValue),
		)
}
