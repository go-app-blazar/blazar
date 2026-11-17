![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/go-app-blazar/blazar?label=version&logo=version&sort=semver)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-app-blazar/blazar)](https://pkg.go.dev/github.com/go-app-blazar/blazar)

# blazar
Quasar-like components for go-app

# Demo
Check out the component demo here: https://blazar.go-app-blazar.com/

# Build your own components
Blazar components are just go-app components, but Blazar includes some tooling to auto-generate the code and interfaces to make them behave like normal go-app components.
This includes functions like `ID`, `Class`, `Attr`, `OnClick`, etc.

## Setup
In your Go project, make sure that you include the Blazar tools.

```
go get -tool github.com/go-app-blazar/blazar/tool/generate-mixin-files
go get -tool github.com/go-app-blazar/blazar/tool/generate-interface-package
```

Determine your package name (such as `example`).

Put your actual components in an `internal` directory (such as `internal/base/example`).
Ensure that each file contains _at most one_ component struct definition.

Create a top-level `generate.go` file that calls the Blazar tools.

```
//go:generate go tool github.com/go-app-blazar/blazar/tool/generate-mixin-files --source-directory=./internal/base/example

//go:generate go tool github.com/go-app-blazar/blazar/tool/generate-interface-package --source-directory=./internal/base/example --target-directory=./example --private-prefix=example --public-prefix=Example
```

Create a Blazar plugin package under `plugin/example`.

That package should contain the CSS file in an `embed.FS`.

It should also define a new `blazarplugin.Plugin`.

## Mixins
The Blazar tools auto-generate component functions based on the mixins that you include on your component.

### `app.Compo`
Your component is expected to embed the stock go-app `app.Compo`.
If not, it will be skipped.

### `mix.MixComposer`
Embed `mix.MixComposer` if your component needs to implement the `app.Composer` interface.

### `mix.MixEvents`
Embed `mix.MixEvents` if your component needs to handle standard HTML mouse and keyboard events.

In your `Render` function, be sure to call `c.MixEvents.Wrap(element)` on whatever `app.UI` element your component ultimately creates.
This will register the events appropriately.

### `mix.MixHTML`
Embed `mix.MixHTML` if your component needs to accept standard HTML attributes (like `ID`, `Class`, `Attr`, etc.).

In your `Render` function, be sure to call `c.MixHTML.Wrap(element)` on whatever `app.UI` element your component ultimately creates.
This will register the HTML properties appropriately.

### `mix.MixRouterViewComponent`
Embed `mix.MixRouterviewComponent` if your component needs to implement `router.RouterViewInterface`.

## Usage

Once done, you can add your plugin to your Blazar app.

```
blazarApp.AddPlugin(example.NewPlugin(example.Config{
    Location: "/web/example/",
}))
```

And you can use your component normally.

```
app.Div().
    Body(
        example.Banner().
            HeadlineText("This is my banner!"),
    )
```

# Tools
This package publishes the following tools.

## `generate-mixin-files`
This generates any additional functions for your components based on any `mix` structs that you used.

## `generate-interface-package`
This generates a whole new copy of your package with only interfaces exposed, similar to how go-app looks and feels.

# Development
All of the files in the `blazar` package are auto-generated based on the files in `internal/base/blazar`.
This allows us to expose true interfaces for Blazar, similar to how go-app exposes its HTML interfaces.
