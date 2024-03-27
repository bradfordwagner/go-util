package flag_helper

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

// TestArgs is a struct that will be used to test the flag_helper
type TestArgs struct {
	Hello    string        `mapstructure:"HELLO_WORLD"`
	Truthy   bool          `mapstructure:"TRUTHY"`
	Duration time.Duration `mapstructure:"DURATION"`
	Int      int           `mapstructure:"INT"`
}

type args[T any] struct {
	name         string
	shortFlag    string
	defaultValue T
	description  string

	getField func(a TestArgs) *T

	// when overriding with env var
	envKey   string
	envValue string
}
type res[T any] struct {
	expected T
}

func test[T any](a args[T], r res[T]) {
	var args TestArgs
	var cmd = &cobra.Command{}
	flags := cmd.PersistentFlags()
	CreateFlag(flags, a.getField(args), a.name, a.shortFlag, a.defaultValue, a.description)

	// override env var
	if a.envKey != "" {
		GinkgoT().Setenv(a.envKey, a.envValue)
	}

	// load env/defaults
	Load(&args)

	var result T = *a.getField(args)
	Expect(result).To(Equal(r.expected))
}

var _ = Describe("FlagHelper", func() {
	It("string,default=default_arg,env=''", func() {
		test[string](
			args[string]{
				name:         "hello_world",
				shortFlag:    "h",
				defaultValue: "default_arg",
				description:  "hello_world, if not provided (env=HELLO_WORLD)",
				getField:     func(a TestArgs) *string { return &a.Hello },
				envKey:       "",
				envValue:     "",
			},
			res[string]{
				expected: "default_arg",
			},
		)
	})
	It("string,default=env_override,env=env_override", func() {
		test[string](
			args[string]{
				name:         "hello_world",
				shortFlag:    "h",
				defaultValue: "default_arg",
				description:  "hello_world, if not provided (env=HELLO_WORLD)",
				getField:     func(a TestArgs) *string { return &a.Hello },
				envKey:       "HELLO_WORLD",
				envValue:     "env_override",
			},
			res[string]{
				expected: "env_override",
			},
		)
	})
	It("bool,default=false,env=''", func() {
		test[bool](
			args[bool]{
				name:         "truthy",
				shortFlag:    "t",
				defaultValue: false,
				description:  "truthy, if not provided (env=TRUTHY)",
				getField:     func(a TestArgs) *bool { return &a.Truthy },
				envKey:       "",
				envValue:     "",
			},
			res[bool]{
				expected: false,
			},
		)
	})
	It("bool,default=true,env=''", func() {
		test[bool](
			args[bool]{
				name:         "truthy",
				shortFlag:    "t",
				defaultValue: true,
				description:  "truthy, if not provided (env=TRUTHY)",
				getField:     func(a TestArgs) *bool { return &a.Truthy },
				envKey:       "",
				envValue:     "",
			},
			res[bool]{
				expected: true,
			},
		)
	})
	It("bool,default=false,env=true", func() {
		test[bool](
			args[bool]{
				name:         "truthy",
				shortFlag:    "t",
				defaultValue: false,
				description:  "truthy, if not provided (env=TRUTHY)",
				getField:     func(a TestArgs) *bool { return &a.Truthy },
				envKey:       "TRUTHY",
				envValue:     "true",
			},
			res[bool]{
				expected: true,
			},
		)
	})
	It("duration,default=1s,env=''", func() {
		test[time.Duration](
			args[time.Duration]{
				name:         "duration",
				shortFlag:    "d",
				defaultValue: time.Second,
				description:  "duration, if not provided (env=DURATION)",
				getField:     func(a TestArgs) *time.Duration { return &a.Duration },
				envKey:       "",
				envValue:     "",
			},
			res[time.Duration]{
				expected: time.Second,
			},
		)
	})
	It("duration,default=1s,env=2s", func() {
		test[time.Duration](
			args[time.Duration]{
				name:         "duration",
				shortFlag:    "d",
				defaultValue: time.Second,
				description:  "duration, if not provided (env=DURATION)",
				getField:     func(a TestArgs) *time.Duration { return &a.Duration },
				envKey:       "DURATION",
				envValue:     "2s",
			},
			res[time.Duration]{
				expected: 2 * time.Second,
			},
		)
	})
	It("int,default=1,env=''", func() {
		test[int](
			args[int]{
				name:         "int",
				shortFlag:    "i",
				defaultValue: 1,
				description:  "int, if not provided (env=INT)",
				getField:     func(a TestArgs) *int { return &a.Int },
				envKey:       "",
				envValue:     "",
			},
			res[int]{
				expected: 1,
			},
		)
	})
	It("int,default=1,env=2", func() {
		test[int](
			args[int]{
				name:         "int",
				shortFlag:    "i",
				defaultValue: 1,
				description:  "int, if not provided (env=INT)",
				getField:     func(a TestArgs) *int { return &a.Int },
				envKey:       "INT",
				envValue:     "2",
			},
			res[int]{
				expected: 2,
			},
		)
	})
})
