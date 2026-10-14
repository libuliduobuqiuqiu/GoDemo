package cmd

import (
	"fmt"
	"godemo/internal/goweb/gogin/app"

	"github.com/spf13/cobra"
)

var ginPort int

var ginCmd = &cobra.Command{
	Use:   "gin",
	Short: "Run the local Gin example from the repository root",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.NewRouter("web/templates/index.html").Run(fmt.Sprintf("127.0.0.1:%d", ginPort))
	},
}

func init() {
	rootCmd.AddCommand(ginCmd)
	ginCmd.Flags().IntVarP(&ginPort, "port", "p", 8999, "Local port for the Gin example.")
}
