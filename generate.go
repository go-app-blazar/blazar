package blazar

//go:generate go run ./tool/generate-mixin/... --source-directory=./internal/base/blazar

//go:generate go run ./tool/generate-interface/... --source-directory=./internal/base/blazar --target-directory=./blazar --private-prefix=blazar --public-prefix=Blazar
