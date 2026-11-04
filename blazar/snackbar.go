package blazar

import (
	"context"
	"log/slog"
	"time"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Snackbar() *blazarSnackbar {
	return &blazarSnackbar{}
}

type blazarSnackbar struct {
	app.Compo
	MixHTML

	IObserve string

	observer *app.Observer
	message  blazarSnackbarMessage
	timer    *time.Timer
}

var _ app.Composer = (*blazarSnackbar)(nil)
var _ app.Mounter = (*blazarSnackbar)(nil)

func (c *blazarSnackbar) OnMount(ctx app.Context) {
	observer := ctx.ObserveState(c.IObserve, &c.message)
	observer.OnChange(func() {
		if debugSnackbar {
			slog.DebugContext(context.TODO(), "Snackbar: Message changed.", "text", c.message.text, "good", c.message.good)
		}

		c.JSValue().Get("classList").Call("add", "blazar-snackbar--visible")
		if c.timer != nil {
			if debugSnackbar {
				slog.DebugContext(context.TODO(), "Snackbar: Stopping existing timer.")
			}
			c.timer.Stop()
		}
		c.timer = time.AfterFunc(10*time.Second, func() {
			if debugSnackbar {
				slog.DebugContext(context.TODO(), "Snackbar: Timer expired normally.")
			}
			if c == nil || c.JSValue() == nil || c.JSValue().Get("classList").IsUndefined() {
				return
			}
			c.JSValue().Get("classList").Call("remove", "blazar-snackbar--visible")
		})
	})
	c.observer = &observer
	if debugSnackbar {
		slog.DebugContext(context.TODO(), "Snackbar: Observer created.", "IObserve", c.IObserve)
	}
}

// BEGIN: mixHTML methods

func (c *blazarSnackbar) Attr(name string, value any) *blazarSnackbar {
	c.mixAttr(name, value)
	return c
}

func (c *blazarSnackbar) Class(class ...string) *blazarSnackbar {
	c.mixClass(class...)
	return c
}

func (c *blazarSnackbar) DataSet(name string, value any) *blazarSnackbar {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarSnackbar) Style(name, value string) *blazarSnackbar {
	c.mixStyle(name, value)
	return c
}

func (c *blazarSnackbar) Title(title string) *blazarSnackbar {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarSnackbar) Observe(id string) *blazarSnackbar {
	c.IObserve = id
	return c
}

func (c *blazarSnackbar) Render() app.UI {
	var body []app.UI

	statusBar := StatusBar().
		Text(c.message.text)
	if c.message.good != nil {
		if *c.message.good {
			statusBar = statusBar.Good()
		} else {
			statusBar = statusBar.Bad()
		}
	}

	body = append(body, statusBar)

	var finalElement app.UI
	finalElement = app.Div().
		Class("blazar-snackbar").
		Role("status").
		Aria("live", "polite").
		Body(
			body...,
		)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}

type blazarSnackbarMessage struct {
	text string
	good *bool
}

func SnackbarMessage() *blazarSnackbarMessage {
	return &blazarSnackbarMessage{}
}

func (c *blazarSnackbarMessage) Text(text string) *blazarSnackbarMessage {
	c.text = text
	return c
}

func (c *blazarSnackbarMessage) Good() *blazarSnackbarMessage {
	c.good = new(bool)
	*c.good = true
	return c
}

func (c *blazarSnackbarMessage) Bad() *blazarSnackbarMessage {
	c.good = new(bool)
	*c.good = false
	return c
}

func SendSnackbarMessage(ctx app.Context, id string, message *blazarSnackbarMessage) {
	ctx.SetState(id, message)
}
