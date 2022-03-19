package verflag

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/usmhk/goserver/pkg/version"
	"os"
	"strconv"
)

type versionValue int

const (
	VersionFalse versionValue = 0
	VersionTrue  versionValue = 1
	VersionRaw   versionValue = 2
)

const strRawVersion = "raw"

func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVar, err := strconv.ParseBool(s)
	if boolVar {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", *v == VersionTrue)
}

func (v *versionValue) Type() string {
	return "version"
}

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func Version() *versionValue {
	p := new(versionValue)
	*p = VersionFalse
	flag.Var(p, "version", "Print version information and quit.")
	flag.Lookup("version").NoOptDefVal = "true"
	return p
}

var versionFlag = Version()

func (v *versionValue) Get() interface{} {
	return v
}

const versionFlagName = "version"

func AddFlags(fs *flag.FlagSet) {
	fs.AddFlag(flag.Lookup(versionFlagName))
}

func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
