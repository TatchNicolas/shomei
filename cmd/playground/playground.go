package playground

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "playground",
	Run: playground,
}

func playground(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
