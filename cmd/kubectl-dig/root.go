package main

import (
	"os"

	"github.com/sysdiglabs/kubectl-dig/pkg/cmd"
	"github.com/spf13/pflag"
    "github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-dig", pflag.ExitOnError)
	pflag.CommandLine = flags

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	root := cmd.NewDigCommand(streams)
	children := subCommands(root)
	routeFrom(root, children)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func hasAlias(root *cobra.Command) string {
	if len(root.Aliases) == 0 {
        return ""
	}
	return root.Aliases[0]
}

func routeFrom(root *cobra.Command, children []string) {
	rootAlias := hasAlias(root)
	if rootAlias == "" || len(os.Args) == 1 || (len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h")) {
		return
	}
    // route to the correct subcommand if first argument matches one of them
    for _, c := range children {
        if os.Args[1] == c {
            return
        }
	}
	// Otherwise construct the right entire command line
    os.Args = append([]string{os.Args[0], root.Aliases[0]}, os.Args[1:]...)
}

// subCommands returns the list of children commands
// excluding the one aliased by the root command.
func subCommands(root *cobra.Command) (ls []string) {
    rootAlias := hasAlias(root)
	if rootAlias == "" || len(os.Args) == 1 || (len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h")) {
		return
	}
    // Iterate children
    for _, child := range root.Commands() {
        isAlias := false
        for _, a := range append(child.Aliases, child.Name()) {
            if a == rootAlias {
                isAlias = true
                break
            }
        }
        if !isAlias {
            ls = append(ls, child.Name())
            ls = append(ls, child.Aliases...)
        }
    }
    return
}