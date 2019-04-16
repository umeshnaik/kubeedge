// Copyright © 2019 Kubeedge Authors <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io"

	"github.com/kubeedge/kubeedge/kubeedgeinst/cmd/options"
	"github.com/kubeedge/kubeedge/kubeedgeinst/cmd/phases/workflow"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/sets"
)

// initOptions defines all the options exposed via flags by kubeadm join.
// Please note that this structure includes the public kubeedge config API, but only a subset of the options
// supported by this api will be exposed as a flag.
type InitOptions struct {
	cfgPath               string
	token                 string
	controlPlane          bool
	ignorePreflightErrors []string
	//externalcfg           *kubeadmapiv1beta1.JoinConfiguration
	certificateKey string
	server         string
}

// joinData defines all the runtime information used when running the kubeadm join worklow;
// this data is shared across all the phases that are included in the workflow.
type initData struct {
	//cfg                   *kubeadmapi.JoinConfiguration
	skipTokenPrint bool
	//initCfg               *kubeadmapi.InitConfiguration
	//tlsBootstrapCfg       *clientcmdapi.Config
	//clientSet             *clientset.Clientset
	ignorePreflightErrors sets.String
	outputWriter          io.Writer
	certificateKey        string
}

// initCmd represents the init command
func NewCmdInit(out io.Writer, init *InitOptions) *cobra.Command {
	if init == nil {
		init = NewinitOptions()
	}
	joinRunner := workflow.NewRunner()
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Bootstraps cloud component. Checks and install (if required) the pre-requisites.",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Work your own magic here
			fmt.Println("init called")
		},
		// We accept the control-plane location as an optional positional argument
		Args: cobra.MaximumNArgs(1),
	}
	//addJoinConfigFlags(cmd.Flags(), joinOptions.externalcfg)
	addJoinOtherFlags(cmd.Flags(), init)

	//joinRunner.AppendPhase(phases.NewPreflightPhase())

	// sets the data builder function, that will be used by the runner
	// both when running the entire workflow or single phases
	joinRunner.SetDataInitializer(func(cmd *cobra.Command, args []string) (workflow.RunData, error) {
		fmt.Println("args is %v", args)
		return newJoinData(cmd, args, init, out)
	})

	// binds the Runner to kubeadm join command by altering
	// command help, adding --skip-phases flag and by adding phases subcommands
	joinRunner.BindToCommand(cmd)

	return cmd
}
func NewinitOptions() *InitOptions {
	// initialize the public kubeEdge config API by applying defaults

	// Add optional config objects to host flags.
	// un-set objects will be cleaned up afterwards (into newJoinData func)

	// Apply defaults
	//kubeadmscheme.Scheme.Default(externalcfg)

	return &InitOptions{
		//externalcfg: externalcfg,
	}
}

// newJoinData returns a new joinData struct to be used for the execution of the kubeedge node init workflow.
// This func takes care of validating joinOptions passed to the command, and then it converts
// options into the internal JoinConfiguration type that is used as input all the phases in the kubeadm join workflow
func newJoinData(cmd *cobra.Command, args []string, opt *InitOptions, out io.Writer) (*initData, error) {
	// Re-apply defaults to the public kubeadm API (this will set only values not exposed/not set as a flags)
	//kubeadmscheme.Scheme.Default(opt.externalcfg)

	// Validate standalone flags values and/or combination of flags and then assigns
	// validated values to the public kubeadm config API when applicable

	// if a token is provided, use this value for both discovery-token and tls-bootstrap-token when those values are not provided
	if len(opt.token) > 0 {
		//add logic
	}

	// if a file or URL from which to load cluster information was not provided, unset the Server.File object
	//if len(opt.externalcfg.Discovery.File.KubeConfigPath) == 0 {
	//	opt.externalcfg.Discovery.File = nil
	//}

	// if an APIServerEndpoint from which to retrieve cluster information was not provided, unset the Discovery.BootstrapToken object
	if len(args) == 0 {
		//add logic
	}

	// if not joining a control plane, unset the ControlPlane object
	if !opt.controlPlane {

	}

	// if the admin.conf file already exists, use it for skipping the discovery process.
	// NB. this case can happen when we are joining a control-plane node only (and phases are invoked atomically)
	//var adminKubeConfigPath = kubeadmconstants.GetAdminKubeConfigPath()

	return &initData{
		//tlsBootstrapCfg:       tlsBootstrapCfg,
		//ignorePreflightErrors: ignorePreflightErrorsSet,
		outputWriter:   out,
		certificateKey: opt.certificateKey,
	}, nil
}

func addJoinOtherFlags(flagSet *flag.FlagSet, initOptions *InitOptions) {
	flagSet.StringVar(
		&initOptions.cfgPath, options.KubeedgeVersion, initOptions.cfgPath,
		"use this key to download and use the required KubeEdge version (Optional, default will be Latest)",
	)
	flagSet.StringSliceVar(
		&initOptions.ignorePreflightErrors, options.DockerVersion, initOptions.ignorePreflightErrors,
		"use this key to download and use the required Docker version (Optional, default will be Latest)",
	)

	flagSet.StringVar(
		&initOptions.certificateKey, options.Kubernetesversion, "",
		"use this key to download and use the required Kubernetes version (Optional, default will be Latest)",
	)

}

func init() {
	//add logic
}
