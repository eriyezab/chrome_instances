/*
Copyright © 2020 Eriyeza Buwembo <buwembo.eriyeza@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
    "os/exec"
    "strings"

	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves the running chrome instance.",
	Long: `Saves the running chrome window(s) with the given name. If there are
                multiple windows then it assumes the user wants to save all the 
                windows.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Saving open chrome instance...")
        // Run the chrome-cli tool and save the results into a list
        command := exec.Command("chrome-cli", "list", "links")
        stdout, err := command.Output()

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        tabs := strings.Fields(string(stdout))
        fmt.Println(tabs)
        // fmt.Println(string(stdout))

        // If the user does not provide a name for the instance prompt them for one.
        if len(args) < 1 {
            fmt.Println("Name this instance: ")
        }

        // fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
