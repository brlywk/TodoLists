package templates

import "embed"

//go:embed *.html partials/*.html api/*.html
var Files embed.FS
