package demo

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type CardPage struct {
	app.Compo
}

var _ app.Initializer = (*CardPage)(nil)

func (c *CardPage) OnInit() {
	slog.DebugContext(context.TODO(), "CardPage: OnInit")
}

func (c *CardPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "CardPage: OnNav")
}

func (c *CardPage) Render() app.UI {
	action1Function := func(ctx app.Context) {
		app.Window().Call("alert", "Action 1 called")
	}
	action2Function := func(ctx app.Context) {
		app.Window().Call("alert", "Action 2 called")
	}
	sleepFunction := func(ctx app.Context) {
		time.Sleep(5 * time.Second)
	}

	return blazar.Page().
		Body(
			blazar.Card(),
			blazar.Card().
				Body(app.Text("Card body text")),
			blazar.Card().
				HeadlineText("Card with title").
				Body(app.Text("Card body text")),
			blazar.Card().
				Icon("circle-info").
				HeadlineText("Card with title and icon").
				Body(app.Text("Card body text")),
			blazar.Card().
				Icon("circle-info").
				HeadlineText("Card with title, icon, and subtitle").
				SubtitleText("Subtitle text").
				Body(app.Text("Card body text")),
			blazar.Card().
				Icon("circle-info").
				HeadlineText("Card with title, icon, subtitle, and actions").
				SubtitleText("Subtitle text").
				Action(
					blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function},
					blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function},
					blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction},
				).
				Body(app.Text("Card body text")),
		)
}
