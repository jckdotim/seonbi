// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

// checkSpellCmd represents the checkSpell command
var checkSpellCmd = &cobra.Command{
	Use:   "checkSpell",
	Short: "Check given Korean's spell",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.PostForm("https://speller.cs.pusan.ac.kr/results",
			url.Values{"text1": {args[0]}})
		if err != nil {
			fmt.Println("error raised")
		} else {
			fmt.Println(resp.Body)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkSpellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkSpellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkSpellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
