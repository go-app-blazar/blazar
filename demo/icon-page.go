package demo

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/go-app-blazar/fontawesome-font"
	"github.com/go-app-blazar/heroicons-font"
	"github.com/go-app-blazar/lucide-font"
	"github.com/go-app-blazar/tabler-font"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IconPage struct {
	app.Compo

	pack              string
	size              float64
	icons             []string
	filter            string
	filterPlaceholder string
}

var _ app.Initializer = (*IconPage)(nil)

func (c *IconPage) OnInit() {
	slog.DebugContext(context.TODO(), "IconPage: OnInit")

	c.pack = "tabler"
	c.size = 2.0
	c.icons = tabler.Icon()
	c.filter = ""
}

func (c *IconPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "IconPage: OnNav")
}

func (c *IconPage) Render() app.UI {
	var filterParts []string
	for _, part := range strings.Split(c.filter, " ") {
		part = strings.TrimSpace(part)
		if part != "" {
			filterParts = append(filterParts, part)
		}
	}

	var filteredIcons []string
	for _, icon := range c.icons {
		allMatch := true
		if len(filterParts) > 0 {
			for _, part := range filterParts {
				if !strings.Contains(icon, part) {
					allMatch = false
					break
				}
			}
		}

		if !allMatch {
			continue
		}

		filteredIcons = append(filteredIcons, icon)
	}

	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Configuration"),
					app.Div().Body(
						blazar.Select().
							Label("Icon Pack").
							AllowedValue(
								blazar.SelectOption{
									Value: "fontawesome",
									Label: "Font Awesome",
								},
								blazar.SelectOption{
									Value: "heroicons",
									Label: "Heroicons",
								},
								blazar.SelectOption{
									Value: "lucide",
									Label: "Lucide",
								},
								blazar.SelectOption{
									Value: "tabler",
									Label: "Tabler",
								},
							).
							Bind(&c.pack).
							On("change", func(ctx app.Context, e app.Event) {
								switch c.pack {
								case "lucide":
									c.icons = lucide.Icon()
								case "fontawesome":
									c.icons = fontawesome.Icon()
								case "heroicons":
									c.icons = heroicons.Icon()
								case "tabler":
									c.icons = tabler.Icon()
								}
								ctx.Update()
							}),
						blazar.Input[float64]().
							Label("Size").
							Min(0.1).
							Max(10.0).
							Step(0.1).
							Bind(&c.size),
						blazar.Input[string]().
							Label("Filter").
							Bind(&c.filterPlaceholder).
							On("change", func(ctx app.Context, e app.Event) {
								c.filter = c.filterPlaceholder
							}).
							OnKeyUp(func(ctx app.Context, e htmlevent.KeyboardEvent) {
								c.filter = e.Target.Get("value").String()
							}),
					),
				),
			app.Div().
				Style("display", "flex").
				Style("flex-wrap", "wrap").
				Style("font-size", fmt.Sprintf("%0.1fem", c.size)).
				Style("gap", "0.1em").
				Body(
					app.Range(filteredIcons).Slice(func(i int) app.UI {
						icon := filteredIcons[i]

						return blazar.Icon().
							Icon(icon).
							Title(icon)
					}),
				),
		)
}
