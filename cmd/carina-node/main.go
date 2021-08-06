/*
   Copyright @ 2021 bocloud <fushaosong@beyondcent.com>.

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
package main

import (
	"github.com/bocloud/carina/cmd/carina-node/run"
	"github.com/bocloud/carina/utils/log"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"os"
)

var gitCommitID = "dev"

func main() {
	printWelcome()
	run.Execute()
}

func printWelcome() {
	if gitCommitID == "" {
		gitCommitID = "dev"
	}
	log.Info("-------- Welcome to use Carina Node Server --------")
	log.Infof("Git Commit ID : %s", gitCommitID)
	log.Infof("node name : %s", os.Getenv("NODE_NAME"))
	log.Info("------------------------------------")
}
