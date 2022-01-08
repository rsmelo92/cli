module jus-cli

go 1.17

require (
	github.com/spf13/cobra v1.3.0
	internal/cli_access v1.0.0
	internal/global_envs v1.0.0
)

require (
	github.com/briandowns/spinner v1.18.0 // indirect
	github.com/enescakir/emoji v1.0.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	internal/utils v1.0.0 // indirect
)

replace (
	internal/cli_access => ./internal/cli_access
	internal/global_envs => ./internal/global_envs
	internal/utils => ./internal/utils
)
