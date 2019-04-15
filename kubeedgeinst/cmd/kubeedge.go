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
	//"fmt"
	"io"

	kubeedgeutil "github.com/kubeedge/kubeedge/kubeedgeinst/cmd/util"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
)

// NewKubeadmCommand returns cobra.Command to run kubeadm command
func NewKubeedgeCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	var rootfsPath string
	cmds := &cobra.Command{
		Use:   "kubeedge",
		Short: "kubeedge: easily bootstrap a secure Kubernetes cluster",
		Long: dedent.Dedent(`

			    ┌──────────────────────────────────────────────────────────┐
			    │ KUBEEDGE                                                 │
			    │ Easily bootstrap a KubeEdge cluster                      │
			    │                                                          │
			    │ Please give us feedback at:                              │
			    │ https://github.com/kubeedge/kubeedge/issues              │
			    └──────────────────────────────────────────────────────────┘

			Example usage:

			    Create a two-machine cluster with one cloud node
                            (which controls the edge cluster), and one edge node
                            (where native containerized application, in the form of
                            pods and deployments run), connects to devices.

			    ┌──────────────────────────────────────────────────────────┐
			    │ On the first machine:                                    │
			    ├──────────────────────────────────────────────────────────┤
			    │ cloud-node# kubeedge cloud init <arguments>                                 │
			    └──────────────────────────────────────────────────────────┘

			    ┌──────────────────────────────────────────────────────────┐
			    │ On the second machine:                                   │
			    ├──────────────────────────────────────────────────────────┤
			    │ edge-node# kubeedge node join <arguments>                │
			    └──────────────────────────────────────────────────────────┘

			    You can then repeat the second step on as many other machines as you like.

		`),

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if rootfsPath != "" {
				if err := kubeedgeutil.Chroot(rootfsPath); err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmds.AddCommand(NewCmdCloud(out))
	cmds.AddCommand(NewCmdNode(out))
	cmds.AddCommand(NewCmdVersion(out))
	cmds.ResetFlags()

	return cmds
}
