// Copyright © 2017 Orinoco Payments
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

	"github.com/orinocopay/go-etherutils/cli"
	"github.com/orinocopay/go-etherutils/ens"
	"github.com/spf13/cobra"
)

// resolverCmd represents the resolver command
var resolverCmd = &cobra.Command{
	Use:   "resolver",
	Short: "Obtain the resolver of an ENS name",
	Long: `Obtain the resolver of a name registered with the Ethereum Name Service (ENS).  For example:

    ens resolver enstest.eth

In quiet mode this will return 0 if the name has a resolver, otherwise 1.`,

	Run: func(cmd *cobra.Command, args []string) {
		registrarContract, err := ens.RegistrarContract(client, rpcclient)
		inState, err := ens.NameInState(registrarContract, args[0], "Owned")
		cli.ErrAssert(inState, err, quiet, "Name not in a suitable state to obtain the resolver")

		registryContract, err := ens.RegistryContract(client, rpcclient)
		cli.ErrCheck(err, quiet, "Failed to obtain registry contract")
		resolver, err := ens.Resolver(registryContract, args[0])
		cli.ErrCheck(err, quiet, "No resolver for that name")
		if !quiet {
			fmt.Println(resolver.Hex())
		}
	},
}

func init() {
	RootCmd.AddCommand(resolverCmd)
}
