/*
Copyright © 2021 Michael Bruskov <mixanemca@yandex.ru>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/mixanemca/dnscli/app"
	"github.com/mixanemca/dnscli/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// fzDelCmd represents the del command
var fzDelCmd = &cobra.Command{
	Aliases: []string{"rm"},
	Use:     "del",
	Short:   "Delete zone from forwarding",
	Example: "  dnscli fz del --name example.com",
	Run:     fzDelCmdRun,
}

func init() {
	fzCmd.AddCommand(fzDelCmd)

	fzDelCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of forwarding zone")
	fzDelCmd.MarkPersistentFlagRequired("name")
}

func fzDelCmdRun(cmd *cobra.Command, args []string) {
	a, err := app.New(
		app.WithBaseURL(viper.GetString("baseURL")),
		app.WithTLS(viper.GetBool("tls"), viper.GetString("cacert"), viper.GetString("cert"), viper.GetString("key")),
		app.WithTimeout(viper.GetInt64("timeout")),
		app.WithDebuggingOutput(viper.GetBool("debug")),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = a.ForwardZones().DeleteByName(models.Canonicalize(name))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if viper.GetString("output-type") == "json" {
		fmt.Println("{}")
		return
	}
	fmt.Printf("domain %s was removed from forwarding zones\n", name)
}
