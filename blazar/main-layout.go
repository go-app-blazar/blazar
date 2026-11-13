package blazar

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/matchmedia"
	"github.com/go-app-blazar/router"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarMainLayout struct {
	app.Compo
	router.RouterViewComponent
	MixHTML

	IHeadlineText    string
	IHeadlineBody    []app.UI
	ISubtitleText    string
	ISubtitleBody    []app.UI
	ITrailerBody     []app.UI
	IDrawerBody      []app.UI
	IResponsiveWidth string

	matchMedia    *matchmedia.MatchMedia
	narrow        bool
	drawerVisible bool
}

var _ router.RouterViewInterface = (*blazarMainLayout)(nil)
var _ app.Mounter = (*blazarMainLayout)(nil)

func MainLayout() *blazarMainLayout {
	return &blazarMainLayout{
		IResponsiveWidth: "900px",
	}
}

func (c *blazarMainLayout) ResponsiveWidth(width string) *blazarMainLayout {
	c.IResponsiveWidth = width
	if c.matchMedia != nil {
		c.matchMedia.SetQuery("screen and (max-width: " + c.IResponsiveWidth + ")")
	}
	return c
}

func (c *blazarMainLayout) HeadlineText(text string) *blazarMainLayout {
	c.IHeadlineText = text
	return c
}

func (c *blazarMainLayout) HeadlineBody(body ...app.UI) *blazarMainLayout {
	c.IHeadlineBody = body
	return c
}

func (c *blazarMainLayout) SubtitleText(text string) *blazarMainLayout {
	c.ISubtitleText = text
	return c
}

func (c *blazarMainLayout) SubtitleBody(body ...app.UI) *blazarMainLayout {
	c.ISubtitleBody = body
	return c
}

func (c *blazarMainLayout) TrailerBody(body ...app.UI) *blazarMainLayout {
	c.ITrailerBody = body
	return c
}

func (c *blazarMainLayout) DrawerBody(body ...app.UI) *blazarMainLayout {
	c.IDrawerBody = body
	return c
}

func (c *blazarMainLayout) OnMount(ctx app.Context) {
	if debugMainLayout {
		slog.DebugContext(ctx.Context, "MainLayout: OnMount", "responsiveWidth", c.IResponsiveWidth)
	}

	c.matchMedia = matchmedia.New(ctx, "screen and (max-width: "+c.IResponsiveWidth+")")
	c.matchMedia.SetOnChange(func(ctx app.Context, value bool) {
		if debugMainLayout {
			slog.DebugContext(ctx.Context, "MainLayout: MatchMedia: OnChange", "value", value)
		}
		c.narrow = value

		ctx.Update()
	})
}

func (c *blazarMainLayout) toggleDrawer() {
	if debugMainLayout {
		slog.DebugContext(context.TODO(), "MainLayout: toggleDrawer")
	}

	c.drawerVisible = !c.drawerVisible
}

func (c *blazarMainLayout) Render() app.UI {
	if debugMainLayout {
		slog.DebugContext(context.TODO(), "MainLayout: Render")
	}

	iconVisible := false
	drawerVisibleClass := ""
	if len(c.IDrawerBody) > 0 {
		drawerVisibleClass = "visible"

		if c.narrow {
			iconVisible = true
			drawerVisibleClass = ""
			if c.drawerVisible {
				drawerVisibleClass = "visible"
			}
		}
	}

	var finalElement app.UI
	finalElement = app.Div().
		Class("blazar-main-layout").
		Body(
			AppBar().
				Class("blazar-main-layout__header").
				NoIcon(!iconVisible).
				Icon(tabler.IconFilledMenu2).
				IconFunction(func(ctx app.Context) {
					c.toggleDrawer()
				}).
				HeadlineText(c.IHeadlineText).
				HeadlineBody(c.IHeadlineBody...).
				SubtitleText(c.ISubtitleText).
				SubtitleBody(c.ISubtitleBody...).
				TrailerBody(c.ITrailerBody...),
			app.Div().
				Class("blazar-main-layout__body").
				Body(
					app.If(len(c.IDrawerBody) > 0, func() app.UI {
						return app.Div().
							Class("blazar-main-layout__drawer", drawerVisibleClass).
							Body(c.IDrawerBody...)
					}),
					app.Div().
						Class("blazar-main-layout__content").
						Body(
							app.If(c.RouterViewComponent.RouterView() != nil, func() app.UI {
								return c.RouterViewComponent.RouterView()
							}),
						),
				),
		)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
