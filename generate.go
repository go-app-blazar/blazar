package blazar

//go:generate go run ./tool/generate-mixin-files/... --source-directory=./internal/base/blazar

//go:generate go run ./tool/generate-interface-package/... --source-directory=./internal/base/blazar --target-directory=./blazar --private-prefix=blazar --public-prefix=Blazar
