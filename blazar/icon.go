package blazar

import (
	"context"
	"log/slog"
	"strings"

	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// defaultIconClass will be directly added to the actual icon element.
var defaultIconClass = ""

// iconPrefix will be prefixed to the icon class if it is not present.
var iconPrefix = ""

// SetDefaultIconClass sets the default icon class.
func SetDefaultIconClass(class string) {
	slog.InfoContext(context.TODO(), "Setting default icon class", "class", class)
	defaultIconClass = class
}

// SetIconPrefix sets the icon prefix.
func SetIconPrefix(prefix string) {
	slog.InfoContext(context.TODO(), "Setting icon prefix", "prefix", prefix)
	iconPrefix = prefix
}

func Icon() *blazarIcon {
	return &blazarIcon{}
}

type blazarIcon struct {
	app.Compo
	UseEvents
	MixHTML

	IconValue string
}

var _ app.Composer = (*blazarIcon)(nil)

// BEGIN: mixHTML methods

func (c *blazarIcon) Attr(name string, value any) *blazarIcon {
	c.mixAttr(name, value)
	return c
}

func (c *blazarIcon) Class(class ...string) *blazarIcon {
	c.mixClass(class...)
	return c
}

func (c *blazarIcon) DataSet(name string, value any) *blazarIcon {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarIcon) Style(name, value string) *blazarIcon {
	c.mixStyle(name, value)
	return c
}

func (c *blazarIcon) Title(title string) *blazarIcon {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

// Icon sets the name of the icon.
//
// If this does not start with "fa-", it will be prefixed with "fa-".
//
// If this contains a space, the first part will be the icon class and the
// second part will be the font class (default is "fa-solid").
func (c *blazarIcon) Icon(icon string) *blazarIcon {
	c.IconValue = icon
	return c
}

func (c *blazarIcon) On(event string, function func(ctx app.Context, e app.Event)) *blazarIcon {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarIcon) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarIcon {
	c.UseEvents.OnClick(function)
	return c
}

func (c *blazarIcon) Render() app.UI {
	fontClass := defaultIconClass

	iconClass := ""
	parts := strings.Split(c.IconValue, " ")
	if len(parts) > 0 {
		iconClass = parts[0]
		parts = parts[1:]

		if iconPrefix == "" {
			if len(parts) > 0 {
				fontClass = parts[0]
			}
		} else {
			if !strings.HasPrefix(iconClass, iconPrefix) {
				iconClass = iconPrefix + iconClass
			}

			if len(parts) > 0 {
				if strings.HasPrefix(parts[0], iconPrefix) {
					fontClass = parts[0]
				}
			}
		}
	}

	var finalElement app.UI
	finalElement = c.UseEvents.Wrap(
		app.Span().
			Class("blazar-icon").
			DataSet("icon", c.IconValue).
			Body(
				app.I().
					Class("blazar-icon__icon", iconClass, fontClass),
			),
	)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
