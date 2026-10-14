/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"godemo/internal/goweb/gogin/proxy/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	address      string
	port         int
	isDebug      bool
	proxyTargets []string
	enableSSH    bool
)

// proxyCmd represents the proxy command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run local HTTP proxy and optional SSH demos.",
	Long: `Run the local Gin proxy demo. HTTP targets must be explicitly allowed with
--proxy-target. SSH command and terminal routes require --enable-ssh and a
GO_DEMO_SSH_KNOWN_HOSTS file. This demo has no client authentication.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if isDebug {
			gin.SetMode(gin.DebugMode)
		}
		return routers.StartProxyService(address, port, proxyTargets, enableSSH)
	},
}

func init() {
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().IntVarP(&port, "port", "p", 8090, "Port to run the server on.")
	proxyCmd.Flags().StringVarP(&address, "address", "a", "127.0.0.1", "Server address.")
	proxyCmd.Flags().BoolVarP(&isDebug, "debug", "d", false, "Run the server debug mode.")
	proxyCmd.Flags().StringArrayVar(&proxyTargets, "proxy-target", nil, "Allowed HTTP or HTTPS origin; repeat for more targets.")
	proxyCmd.Flags().BoolVar(&enableSSH, "enable-ssh", false, "Enable SSH command and terminal routes; requires GO_DEMO_SSH_KNOWN_HOSTS.")
}
