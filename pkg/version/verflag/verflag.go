package verflag

import (
	flag "github.com/spf13/pflag"
)

type versionValue int

const (
	versionFalse versionValue = 0
)

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() interface{} {
	return v
}

const versionFlagName = "version"

func AddFlags(fs *flag.FlagSet) {
	//fs.AddFlag(flag.Lookup(versionFlagName))
}
