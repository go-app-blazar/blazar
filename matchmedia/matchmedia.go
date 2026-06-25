package matchmedia

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// MatchMedia is a wrapper around the matchMedia API.
//
// This is used to watch for changes in a given media query.
type MatchMedia struct {
	ctx                  app.Context                       // The context of the component that is using this MatchMedia.
	mediaQuery           string                            // The media query to watch for changes.
	onChangeCallback     func(ctx app.Context, value bool) // The callback to invoke when the media query changes.
	jsMediaQuery         app.Value                         // The JavaScript value of the media query.
	jsMediaQueryListener app.Func                          // The JavaScript function to listen for changes.
	value                bool                              // The current value of the media query.
}

// New creates a new MatchMedia instance.
//
// The ctx given is the context of the component that is using this MatchMedia.
func New(ctx app.Context, mediaQuery string) *MatchMedia {
	if debugMatchMedia {
		slog.DebugContext(context.TODO(), "MatchMedia: New", "mediaQuery", mediaQuery)
	}
	m := &MatchMedia{
		ctx: ctx,
	}
	m.SetQuery(mediaQuery)
	return m
}

// SetQuery sets the media query to watch for changes.
func (m *MatchMedia) SetQuery(mediaQuery string) {
	if debugMatchMedia {
		slog.DebugContext(context.TODO(), "MatchMedia: Query", "mediaQuery", mediaQuery)
	}

	// If the query is the same, then don't do anything.
	if m.mediaQuery == mediaQuery {
		return
	}

	if m.jsMediaQueryListener != nil {
		m.jsMediaQuery.Call("removeEventListener", "change", m.jsMediaQueryListener)
		m.jsMediaQueryListener.Release()
		m.jsMediaQueryListener = nil
	}

	m.mediaQuery = mediaQuery
	m.jsMediaQuery = app.Window().Call("matchMedia", mediaQuery)
	if m.jsMediaQuery.IsNull() {
		slog.ErrorContext(context.TODO(), "MatchMedia: Query: jsMediaQuery is null")
		return
	}
	//slog.DebugContext(context.TODO(), "MatchMedia: jsMediaQuery", "jsMediaQuery", m.jsMediaQuery)

	m.jsMediaQueryListener = app.FuncOf(func(this app.Value, args []app.Value) any {
		//slog.DebugContext(context.TODO(), "MatchMedia: OnChange", "args", args)

		m.value = args[0].Get("matches").Bool()
		if debugMatchMedia {
			slog.DebugContext(context.TODO(), "MatchMedia: OnChange", "value", m.value)
		}

		if m.onChangeCallback != nil {
			if len(args) > 0 {
				m.onChangeCallback(m.ctx, m.value)
			}
		}
		return nil
	})

	m.jsMediaQuery.Call("addEventListener", "change", m.jsMediaQueryListener)

	if debugMatchMedia {
		slog.DebugContext(context.TODO(), "MatchMedia: SetQuery: Invoking event listener manually")
	}
	m.jsMediaQueryListener.Invoke(m.jsMediaQuery)
}

// SetOnChange sets the callback to invoke when the media query changes.
func (m *MatchMedia) SetOnChange(callback func(ctx app.Context, value bool)) {
	m.onChangeCallback = callback

	// Call the callback with the current value.
	m.onChangeCallback(m.ctx, m.value)
}
