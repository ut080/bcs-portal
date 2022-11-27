package main

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/ut080/bcs-portal/app/dpdocs"
)

// BaseVersion is the SemVer-formatted string that defines the current version of ag7if.
// Build information will be added at compile-time.
const BaseVersion = "0.1.0-develop"

// BuildTime is a timestamp of when the build is run. This variable is set at compile-time.
var BuildTime string

// GitRevision is the current Git commit ID. If the tree is dirty at compile-time, an "x-" is prepended to the hash.
// This variable is set at compile-time.
var GitRevision string

// GitBranch is the name of the active Git branch at compile-time. This variable is set at compile-time.
var GitBranch string

func main() {
	version := fmt.Sprintf(
		"%s+%s.%s.%s",
		BaseVersion,
		GitBranch,
		GitRevision,
		BuildTime,
	)
	viper.Set("version", version)
	dpdocs.Execute()
}
