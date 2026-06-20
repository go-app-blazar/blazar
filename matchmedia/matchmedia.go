package matchmedia

import (
	"context"
	"log/slog"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type MatchMedia struct {
	mediaQuery           string
	onChangeCallback     func(ctx app.Context, value bool)
	jsMediaQuery         app.Value
	jsMediaQueryListener app.Func
}

func New(mediaQuery string) *MatchMedia {
	m := &MatchMedia{
		mediaQuery: mediaQuery,
	}

	slog.InfoContext(context.TODO(), "MatchMedia: New", "mediaQuery", mediaQuery)
	m.jsMediaQuery = app.Window().Call("matchMedia", mediaQuery)
	if m.jsMediaQuery.IsNull() {
		slog.ErrorContext(context.TODO(), "MatchMedia: jsMediaQuery is null")
		return nil
	}
	slog.InfoContext(context.TODO(), "MatchMedia: jsMediaQuery", "jsMediaQuery", m.jsMediaQuery)

	m.jsMediaQueryListener = app.FuncOf(func(this app.Value, args []app.Value) any {
		slog.InfoContext(context.TODO(), "MatchMedia: OnChange", "args", args)
		if m.onChangeCallback != nil {
			if len(args) > 0 {
				slog.InfoContext(context.TODO(), "MatchMedia: OnChange", "matches", args[0].Get("matches").Bool())
				//m.onChangeCallback(nil, args[0].Get("matches").Bool())
			}
		}
		return nil
	})

	m.jsMediaQuery.Call("addEventListener", "change", m.jsMediaQueryListener)

	return m
}

func (m *MatchMedia) SetQuery(mediaQuery string) *MatchMedia {
	m.mediaQuery = mediaQuery
	// TODO: Update the media query listener and such.
	return m
}

func (m *MatchMedia) OnChange(callback func(ctx app.Context, value bool)) *MatchMedia {
	m.onChangeCallback = callback
	return m
}
