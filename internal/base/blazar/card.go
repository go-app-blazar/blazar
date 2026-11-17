package blazar

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarCard struct {
	app.Compo
	MixHTML

	IIcon         string
	IHeadlineText string
	IHeadlineBody []app.UI
	ISubtitleText string
	ISubtitleBody []app.UI
	IBody         []app.UI
	IActions      []FormAction
}

func Card() *blazarCard {
	return &blazarCard{}
}

func (c *blazarCard) Icon(icon string) *blazarCard {
	c.IIcon = icon
	return c
}

func (c *blazarCard) HeadlineText(text string) *blazarCard {
	c.IHeadlineText = text
	return c
}

func (c *blazarCard) HeadlineBody(body ...app.UI) *blazarCard {
	c.IHeadlineBody = body
	return c
}

func (c *blazarCard) SubtitleText(text string) *blazarCard {
	c.ISubtitleText = text
	return c
}

func (c *blazarCard) SubtitleBody(body ...app.UI) *blazarCard {
	c.ISubtitleBody = body
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
			app.If(c.IHeadlineText != "" || len(c.IHeadlineBody) > 0 || c.ISubtitleText != "" || len(c.ISubtitleBody) > 0, func() app.UI {
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
								app.If(c.IHeadlineText != "" || len(c.IHeadlineBody) > 0, func() app.UI {
									return app.Div().
										Class("blazar-card__headline").
										Body(
											app.IfSlice(len(c.IHeadlineBody) > 0, func() []app.UI {
												return c.IHeadlineBody
											}).Else(func() app.UI {
												return app.Div().
													Class("blazar-card__headline-text").
													Text(c.IHeadlineText)
											}),
										)
								}),
								app.If(c.ISubtitleText != "" || len(c.ISubtitleBody) > 0, func() app.UI {
									return app.Div().
										Class("blazar-card__subtitle").
										Body(
											app.IfSlice(len(c.ISubtitleBody) > 0, func() []app.UI {
												return c.ISubtitleBody
											}).Else(func() app.UI {
												return app.Div().
													Class("blazar-card__subtitle-text").
													Text(c.ISubtitleText)
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
