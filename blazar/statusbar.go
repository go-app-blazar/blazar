package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func StatusBar() *blazarStatusBar {
	return &blazarStatusBar{}
}

type blazarStatusBar struct {
	app.Compo
	MixHTML

	IText    string
	IGood    *bool
	IActions []StatusBarAction
}

var _ app.Composer = (*blazarStatusBar)(nil)

type StatusBarAction struct {
	Name     string
	Icon     string
	To       string
	Function func(ctx app.Context)
	Disabled bool
}

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

func (c *blazarStatusBar) Action(actions ...StatusBarAction) *blazarStatusBar {
	c.IActions = actions
	return c
}

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
			app.If(len(c.IActions) > 0, func() app.UI {
				return app.Div().
					Class("blazar-statusbar__spacer").
					Style("flex", "1")
			}),
			app.Range(c.IActions).Slice(func(i int) app.UI {
				action := c.IActions[i]

				button := Button().
					Class("blazar-statusbar__action").
					Style("--color", "var(--blazar-theme-on-primary)").
					Label(action.Name).
					Icon(action.Icon).
					Flat(true).
					Outline(false).
					Disabled(action.Disabled)
				if action.To != "" {
					button = button.To(action.To)
				}
				if action.Function != nil {
					button = button.OnClick(func(ctx app.Context, e htmlevent.MouseEvent) {
						action.Function(ctx)
					})
				}
				return button
			}),
		)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
