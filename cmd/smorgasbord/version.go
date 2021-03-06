/*
Copyright 2020 Smorgasbord Authors

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
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	version string = "DEV"
	commit  string = "DEBUG"
)

func newVersionCmd(out io.Writer) *cobra.Command {
	var json bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints version information.",
		Long:  `Prints version information and the commit.`,
		Run: func(cmd *cobra.Command, args []string) {
			if json {
				fmt.Fprintf(out, `{ "app": "smorgasbord", "version": "%s", "commit": "%s" }`, version, commit)
			} else {
				fmt.Fprintf(out, "smorgasbord\nversion: %s commit: %s\n", version, commit)
			}
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&json, "json", false, "Whether to log version and commit as json.")

	return cmd
}
