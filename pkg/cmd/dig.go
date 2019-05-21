package cmd

import (
	"fmt"

	"github.com/sysdiglabs/kubectl-dig/pkg/factory"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	digLong     = `...`
	digExamples = `%[1]s ...`
)

// DigOptions ...
type DigOptions struct {
	configFlags *genericclioptions.ConfigFlags

	genericclioptions.IOStreams
}

// NewDigOptions provides an instance of DigOptions with default values.
func NewDigOptions(streams genericclioptions.IOStreams) *DigOptions {
	return &DigOptions{
		configFlags: genericclioptions.NewConfigFlags(false),

		IOStreams: streams,
	}
}

// NewDigCommand creates the dig command and its nested children.
func NewDigCommand(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewDigOptions(streams)

	cmd := &cobra.Command{
		Use:                   "dig",
		Aliases:               []string{"run"},
		DisableFlagsInUseLine: true,
		Short:                 `Dig your cluster`,                  // Wrap with i18n.T()
		Long:                  digLong,                             // Wrap with templates.LongDesc()
		Example:               fmt.Sprintf(digExamples, "kubectl"), // Wrap with templates.Examples()
		Run: func(c *cobra.Command, args []string) {
			c.SetOutput(streams.ErrOut)
			cobra.NoArgs(c, args)
			c.Help()
		},
	}

	flags := cmd.PersistentFlags()
	o.configFlags.AddFlags(flags)

	matchVersionFlags := factory.NewMatchVersionFlags(o.configFlags)
	matchVersionFlags.AddFlags(flags)

	// flags.AddGoFlagSet(flag.CommandLine) // todo(leodido) > evaluate whether we need this or not

	f := factory.NewFactory(matchVersionFlags)

	cmd.AddCommand(NewRunCommand(f, streams))

	// Override help on all the commands tree
	walk(cmd, func(c *cobra.Command) {
		c.Flags().BoolP("help", "h", false, fmt.Sprintf("Help for the %s command", c.Name()))
	})

	return cmd
}

// walk calls f for c and all of its children.
func walk(c *cobra.Command, f func(*cobra.Command)) {
	f(c)
	for _, c := range c.Commands() {
		walk(c, f)
	}
}
