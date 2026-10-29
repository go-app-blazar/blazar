package blazar

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarCard struct {
	app.Compo
	MixHTML

	IIcon             string
	IHeadline         string
	IHeadlineFunction func() app.UI
	ISubtitle         string
	ISubtitleFunction func() app.UI
	IBody             []app.UI
	IActions          []FormAction
}

func Card() *blazarCard {
	return &blazarCard{}
}

// BEGIN: mixHTML methods

func (c *blazarCard) Attr(name string, value any) *blazarCard {
	c.mixAttr(name, value)
	return c
}

func (c *blazarCard) Class(class ...string) *blazarCard {
	c.mixClass(class...)
	return c
}

func (c *blazarCard) DataSet(name string, value any) *blazarCard {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarCard) Style(name, value string) *blazarCard {
	c.mixStyle(name, value)
	return c
}

func (c *blazarCard) Title(title string) *blazarCard {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarCard) Icon(icon string) *blazarCard {
	c.IIcon = icon
	return c
}

func (c *blazarCard) HeadlineText(text string) *blazarCard {
	c.IHeadline = text
	return c
}

func (c *blazarCard) HeadlineFunction(function func() app.UI) *blazarCard {
	c.IHeadlineFunction = function
	return c
}

func (c *blazarCard) SubtitleText(text string) *blazarCard {
	c.ISubtitle = text
	return c
}

func (c *blazarCard) SubtitleFunction(function func() app.UI) *blazarCard {
	c.ISubtitleFunction = function
	return c
}

func (c *blazarCard) Body(body ...app.UI) *blazarCard {
	c.IBody = body
	return c
}

func (c *blazarCard) Action(actions ...FormAction) *blazarCard {
	c.IActions = actions
	return c
}

func (c *blazarCard) Render() app.UI {
	if debugCard {
		slog.DebugContext(context.TODO(), "Card: Render")
	}

	var finalElement app.UI
	finalElement = app.Div().
		Class("blazar-card", "blazar-mixin-dropshadow").
		Style("width", "100%").
		Body(
			app.If(c.IHeadline != "" || c.IHeadlineFunction != nil || c.ISubtitle != "" || c.ISubtitleFunction != nil, func() app.UI {
				return app.Div().
					Class("blazar-card__header").
					Body(
						app.If(c.IIcon != "", func() app.UI {
							return app.Div().
								Class("blazar-card__icon").
								Body(
									Icon().
										Icon(c.IIcon),
								)
						}),
						app.Div().
							Body(
								app.If(c.IHeadline != "" || c.IHeadlineFunction != nil, func() app.UI {
									return app.Div().
										Class("blazar-card__headline").
										Body(
											app.If(c.IHeadlineFunction != nil, func() app.UI {
												return c.IHeadlineFunction()
											}).Else(func() app.UI {
												return app.Div().
													Class("blazar-card__headline-text").
													Text(c.IHeadline)
											}),
										)
								}),
								app.If(c.ISubtitle != "" || c.ISubtitleFunction != nil, func() app.UI {
									return app.Div().
										Class("blazar-card__subtitle").
										Body(
											app.If(c.ISubtitleFunction != nil, func() app.UI {
												return c.ISubtitleFunction()
											}).Else(func() app.UI {
												return app.Div().
													Class("blazar-card__subtitle-text").
													Text(c.ISubtitle)
											}),
										)
								}),
							),
					)
			}),
			Form().
				Action(c.IActions...).
				Body(
					app.Div().
						Class("blazar-card__body").
						Body(c.IBody...),
				),
		)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
