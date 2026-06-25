module github.com/go-app-blazar/blazar

go 1.26.2

require (
	github.com/go-app-blazar/router v0.1.0
	github.com/joho/godotenv v1.5.1
	github.com/lmittmann/tint v1.1.3
	github.com/mattn/go-isatty v0.0.22
	github.com/maxence-charriere/go-app/v11 v11.0.4
	golang.org/x/text v0.38.0
)

require (
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)

replace github.com/maxence-charriere/go-app/v11 => github.com/tekkamanendless/fork-of-maxence-charriere-go-app/v11 v11.0.0-20260624062618-c2f3e531f36d

//replace github.com/maxence-charriere/go-app/v11 => ../../tekkamanendless/fork-of-maxence-charriere-go-app
