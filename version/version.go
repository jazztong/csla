package version

import (
	ver "github.com/hashicorp/go-version"
)

// Version is the main version number that is being run at the moment.
var Version = "0.0.0"

// SemVer is an instance of version.Version. This has the secondary
// benefit of verifying during tests and init time that our version is a
// proper semantic version, which should always be the case.
var SemVer *ver.Version

func init() {
	SemVer = ver.Must(ver.NewVersion(Version))
}

// String returns the complete version string
func String() string {
	return Version
}
