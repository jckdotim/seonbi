// Copyright © 2019 JC Kim
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
	"io/ioutil"
	"strings"
	"regexp"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/TylerBrock/colorjson"
)

// checkSpellCmd represents the checkSpell command
var checkSpellCmd = &cobra.Command{
	Use:   "checkSpell",
	Short: "Check given Korean's spell",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.PostForm("http://speller.cs.pusan.ac.kr/results",
			url.Values{"text1": {args[0]}})
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		if !strings.Contains(string(body), "맞춤법과 문법 오류를 찾지") {
			re := regexp.MustCompile(`data = (\[[^;].*\]);`)
			match := re.FindStringSubmatch(string(body))
			var obj map[string]interface{}
			json.Unmarshal([]byte(string(match[1])), &obj)
			result, _ := colorjson.Marshal(obj)
			fmt.Println(string(match[1]))
			fmt.Println(string(result))
		}
	},
}

func init() {
	rootCmd.AddCommand(checkSpellCmd)
}
