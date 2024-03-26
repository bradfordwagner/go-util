package flag_helper

import (
	"time"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// CreateFlag creates a flag and binds it to viper (so env can be marshalled into a struct seamlessly)
func CreateFlag[T any](
	flags *flag.FlagSet,
	valRef interface{},
	name string,
	short string,
	defaultValue T,
	description string,
) {
	switch castedDefaultValue := any(defaultValue).(type) {
	case string:
		ref := valRef.(*string)
		flags.StringVarP(ref, name, short, castedDefaultValue, description)
	case time.Duration:
		ref := valRef.(*time.Duration)
		flags.DurationVarP(ref, name, short, castedDefaultValue, description)
	case bool:
		ref := valRef.(*bool)
		flags.BoolVarP(ref, name, short, castedDefaultValue, description)
	default:
		panic("unsupported type")
	}
	viper.BindPFlag(name, flags.Lookup(name))
}

// Load loads the flags into the viper
func Load[T any](args *T) {
	viper.AutomaticEnv()
	viper.Unmarshal(args) // when verbs have divergent args this will need to be moved into cmd specific methods
}
