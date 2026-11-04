package blazar

import (
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func StatusBar() *blazarStatusBar {
	return &blazarStatusBar{}
}

type blazarStatusBar struct {
	app.Compo
	MixHTML

	IText string
	IGood *bool
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
	c.IText = text
	return c
}

func (c *blazarStatusBar) Good() *blazarStatusBar {
	c.IGood = new(bool)
	*c.IGood = true
	return c
}

func (c *blazarStatusBar) Bad() *blazarStatusBar {
	c.IGood = new(bool)
	*c.IGood = false
	return c
}

func (c *blazarStatusBar) Neutral() *blazarStatusBar {
	c.IGood = nil
	return c
}

func (c *blazarStatusBar) Render() app.UI {
	goodClass := ""
	goodIcon := tabler.IconOutlineInfoCircle
	if c.IGood != nil {
		if *c.IGood {
			goodClass = "blazar-statusbar--good"
			goodIcon = tabler.IconOutlineCircleCheck
		} else {
			goodClass = "blazar-statusbar--bad"
			goodIcon = tabler.IconOutlineAlertCircle
		}
	}

	var finalElement app.UI
	finalElement = app.Div().
		Class("blazar-statusbar", goodClass).
		Body(
			app.If(goodIcon != "", func() app.UI {
				return Icon().
					Class("blazar-statusbar__icon").
					Icon(goodIcon)
			}),
			app.Span().
				Class("blazar-statusbar__text").
				Text(c.IText),
		)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
