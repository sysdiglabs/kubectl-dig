package cmd

import (
	"context"
	"fmt"

	"github.com/leodido/kubectl-dig/pkg/attacher"
	"github.com/leodido/kubectl-dig/pkg/digjob"
	"github.com/leodido/kubectl-dig/pkg/factory"
	"github.com/leodido/kubectl-dig/pkg/meta"
	"github.com/leodido/kubectl-dig/pkg/signals"
	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes/scheme"
	batchv1client "k8s.io/client-go/kubernetes/typed/batch/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

var (
	// ImageNameTag represents the default digrunner image
	ImageNameTag = "sysdig/sysdig:latest"
)

var (
	runShort = `Run sysdig on your nodes` // Wrap with i18n.T()

	runLong = runShort

	runExamples = `%[1]s ...`

	runCommand           = "run"
	usageString          = "(NODE)"
	requiredArgErrString = fmt.Sprintf("%s is a required argument for the %s command", usageString, runCommand)
)

// todo > flags/ideas
// --scap
// --filter to directly pass a filter to sysdig
// check for lib module

// RunOptions ...
type RunOptions struct {
	genericclioptions.IOStreams
	clientConfig *rest.Config

	nodeName          string
	namespace         string
	explicitNamespace bool

	// Flags local to this command
	serviceAccount string
	imageName      string

	// Explicit arguments
	resourceArg string
}

// NewRunOptions provides an instance of RunOptions with default values.
func NewRunOptions(streams genericclioptions.IOStreams) *RunOptions {
	return &RunOptions{
		IOStreams: streams,

		serviceAccount: "",
		imageName:      ImageNameTag,
	}
}

// NewRunCommand provides the run command wrapping RunOptions.
func NewRunCommand(factory factory.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewRunOptions(streams)

	cmd := &cobra.Command{
		Use:          fmt.Sprintf("%s %s", runCommand, usageString),
		Short:        runShort,
		Long:         runLong,                             // Wrap with templates.LongDesc()
		Example:      fmt.Sprintf(runExamples, "kubectl"), // Wrap with templates.Examples()
		SilenceUsage: true,
		PreRunE: func(c *cobra.Command, args []string) error {
			return o.Validate(c, args)
		},
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(factory, c, args); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				fmt.Fprintln(o.ErrOut, err.Error())
				return nil
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&o.serviceAccount, "serviceaccount", o.serviceAccount, "Service account for the digrunner to talk to Kubernetes")
	cmd.Flags().StringVar(&o.imageName, "imagename", o.imageName, "Custom image for the digrunner")

	return cmd
}

// Validate validates the arguments and flags populating RunOptions accordingly.
func (o *RunOptions) Validate(cmd *cobra.Command, args []string) error {
	switch len(args) {
	case 1:
		o.resourceArg = args[0]
		break
	default:
		return fmt.Errorf(requiredArgErrString)
	}

	return nil
}

// Complete completes the setup of the command.
func (o *RunOptions) Complete(factory factory.Factory, cmd *cobra.Command, args []string) error {
	// Prepare namespace
	var err error
	o.namespace, o.explicitNamespace, err = factory.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	// Look for the target object
	x := factory.
		NewBuilder().
		WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
		NamespaceParam(o.namespace).
		SingleResourceType().
		ResourceNames("nodes", o.resourceArg). // Search nodes by default
		Do()

	obj, err := x.Object()
	if err != nil {
		return err
	}

	var node *v1.Node

	switch v := obj.(type) {
	case *v1.Node:
		node = v
		break
	default:
		return fmt.Errorf("first argument must be %s", usageString)
	}

	if node == nil {
		return fmt.Errorf("could not determine on which node to run the dig runner")
	}

	labels := node.GetLabels()
	val, ok := labels["kubernetes.io/hostname"]
	if !ok {
		return fmt.Errorf("label kubernetes.io/hostname not found in node")
	}
	o.nodeName = val

	// Prepare client
	o.clientConfig, err = factory.ToRESTConfig()
	if err != nil {
		return err
	}

	return nil
}

// Run executes the run command.
func (o *RunOptions) Run() error {
	jid := uuid.NewUUID()
	jobsClient, err := batchv1client.NewForConfig(o.clientConfig)
	if err != nil {
		return err
	}

	coreClient, err := corev1client.NewForConfig(o.clientConfig)
	if err != nil {
		return err
	}

	djc := &digjob.Client{
		JobClient: jobsClient.Jobs(o.namespace),
	}

	dj := digjob.Job{
		Name:           fmt.Sprintf("%s%s", meta.ObjectNamePrefix, string(jid)),
		Namespace:      o.namespace,
		ServiceAccount: o.serviceAccount,
		ID:             jid,
		Hostname:       o.nodeName,
		ImageNameTag:   o.imageName,
	}

	job, err := djc.CreateJob(dj)
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx = signals.WithStandardSignals(ctx)
	a := attacher.NewAttacher(coreClient, o.clientConfig, o.IOStreams)
	a.WithContext(ctx)
	a.AttachJob(dj.ID, job.Namespace)

	return nil
}
