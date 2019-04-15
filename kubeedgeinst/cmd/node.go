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

	node "github.com/kubeedge/kubeedge/kubeedgeinst/cmd/cmdnode"

	"github.com/spf13/cobra"
)

// nodeCmd represents the node command
func NewCmdNode(out io.Writer) *cobra.Command {
	var join joinOptions
	var cmd = &cobra.Command{
		Use:   "node",
		Short: "Edge component command option for KubeEdge",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Work your own magic here
			fmt.Println("node called")
		},
	}
	// joinOptions defines all the options exposed via flags by kubeadm join.
	// Please note that this structure includes the public kubeedge config API, but only a subset of the options
	// supported by this api will be exposed as a flag.

	cmd.AddCommand(NewCmdJoin(out, &join))
	cmd.AddCommand(node.NewCmdReset(out))
	cmd.AddCommand(node.NewCmdInit(out))
	return cmd
}

func init() {
	//add logic
}
