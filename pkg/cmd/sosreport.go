package cmd

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd/api"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"github.com/openshift/client-go/"
)

var Version = "0.0.1-alpha"

type SosreportOptions struct {
	configFlags *genericclioptions.ConfigFlags

	nodeName string
	caseNumber string
	printVersion bool

	rawConfig api.Config
	args           []string

	genericclioptions.IOStreams
}

func NewSosreportOptions(streams genericclioptions.IOStreams) *SosreportOptions {
	return &SosreportOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams: streams,
	}
}

// NewCmdNamespace provides a cobra command wrapping NamespaceOptions
func NewCmdSosreport(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewSosreportOptions(streams)

	cmd := &cobra.Command{
		Use:	   "sosreport [node-name] [case-num] [flags]",
		Short:	   "Generate a sosreport from a specific OpenShift Node",
		RunE: func(c *cobra.Command, args []string) error {
			if o.printVersion {
				fmt.Println(Version)
				os.Exit(0)
			}

			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&o.printVersion, "version", false, "prints version of plugin")
	cmd.Flags().StringVar(&o.nodeName, "node", o.nodeName, "generate sosreport on this node")
	cmd.Flags().StringVar(&o.caseNumber, "case", o.caseNumber, "Red Hat support case number")

	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Complete sets all information required for updating the current context
func (o *SosreportOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args

	var err error
	o.rawConfig, err = o.configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return err
	}


}

// Validate ensures that all required arguments and flag values are provided
func (o *SosreportOptions) Validate() error {
	if len(o.args) == 0 {
		return fmt.Errorf("atleast one argument is required")
	}

	if len(o.args) > 1 {
		return fmt.Errorf("only one argument is allowed")
	}

	return nil
}

// Run fetches the given secret manifest from the cluster, decodes the payload, opens an editor to make changes, and applies the modified manifest when done
func (o *SosreportOptions) Run() error {
	secret, err := secrets.Get(o.kubeclient, o.secretName, o.namespace)
	if err != nil {
		return err
	}

}