package blazar

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func StatusBar() *blazarStatusBar {
	return &blazarStatusBar{}
}

type blazarStatusBar struct {
	app.Compo
	MixHTML

	TextValue string
	GoodValue *bool
}

var _ app.Composer = (*blazarStatusBar)(nil)

// BEGIN: mixHTML methods

func (c *blazarStatusBar) Attr(name string, value any) *blazarStatusBar {
	c.mixAttr(name, value)
	return c
}

func (c *blazarStatusBar) Class(class ...string) *blazarStatusBar {
	c.mixClass(class...)
	return c
}

func (c *blazarStatusBar) DataSet(name string, value any) *blazarStatusBar {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarStatusBar) Style(name, value string) *blazarStatusBar {
	c.mixStyle(name, value)
	return c
}

func (c *blazarStatusBar) Title(title string) *blazarStatusBar {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

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

	var finalElement app.UI
	finalElement = app.Div().
		Class("blazar-statusbar", goodClass).
		Body(
			app.If(goodIcon != "", func() app.UI {
				return Icon().
					Icon(goodIcon)
			}),
			app.Span().
				Class("blazar-button-label").
				Text(c.TextValue),
		)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
