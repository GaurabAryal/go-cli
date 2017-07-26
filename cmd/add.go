// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
	"go-cli/todo"
)
var priority int
// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add will create a new todo item to the list`,
	Run: addRun,
}

func addRun(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataFile)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		for _, x := range args {
			item := todo.Item{Text: x}
			item.SetPriority(priority)
			items = append(items, item)
		}
		err = todo.SaveItems(dataFile, items)
		if err != nil {
			fmt.Errorf("%v", err)
		}
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority,
		"priority", "p", 2, "Priority: 1, 2, 3")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
