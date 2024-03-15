module github.com/institute-atri/wastrap

go 1.22.1

// Package dependencies - ATRI
require (
	github.com/institute-atri/glogger v0.0.0-20240308200612-9c47543fa3f6
   	github.com/institute-atri/gnet v0.0.0-20240312004547-aac3a62e404a
)

// Package dependencies
require (
	github.com/spf13/cobra v1.8.0
	gopkg.in/yaml.v3 v3.0.1
)

// Package dependencies - Indirect
require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
