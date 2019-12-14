module github.com/crdsdev/crds/cli

go 1.13

replace github.com/crdsdev/crds/internal => ../internal

require (
	github.com/crdsdev/crds/internal v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v0.0.5
)
