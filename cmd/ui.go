/*
Copyright © 2020 Juan Ezquerro LLanes <arrase@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/arrase/multi-repo-workspace/cli/openurl"
	"github.com/arrase/multi-repo-workspace/cli/server"
	"github.com/spf13/cobra"
	"time"
)

func openBrowser() {
	time.Sleep(1000 * time.Millisecond)
	openurl.OpenBrowser("http://localhost:8080")
}

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Open Web UI",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.WSServer{}
		go openBrowser()
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
