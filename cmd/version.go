package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version   = "1.4.0"
	githash   = "HEAD"
	gitbranch = "master"
	date      = "1970-01-01T00:00:00Z UTC"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Erlenmeyer version %s %s/%s\n", version, gitbranch, githash)
		fmt.Printf("Erlenmeyer build date %s\n", date)
		fmt.Printf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}
