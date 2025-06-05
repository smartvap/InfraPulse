package cmd

import (
	"github.com/clarechu/infrapulse/pkg/server"
	"github.com/clarechu/infrapulse/pkg/utils/homedir"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
)

func ServerCommand(args []string) *cobra.Command {
	config := &server.CmdbConfig{}
	serverCommand := &cobra.Command{
		Use:               "server",
		Short:             "run cmdb server ",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `The new generation of CMDB`,
		Run: func(cmd *cobra.Command, args []string) {
			klog.Info("cmdb start ...")

			if config.DataRoot == "" {
				config.DataRoot = filepath.Join(homedir.HomeDir(), ".cmdb")
			}
			s, err := server.NewCmdb(config)
			if err != nil {
				klog.Errorf("new cmdb config error:%s", err)
				os.Exit(-1)
			}
			s.ListenAndServe()
		},
	}
	AddServerCommandFlag(serverCommand, config)
	return serverCommand
}

func AddServerCommandFlag(serverCommand *cobra.Command, config *server.CmdbConfig) {
	serverCommand.Flags().Int32VarP(&config.Port, "port", "p", 9090, "http server port")
	serverCommand.Flags().Int32Var(&config.ProxyPort, "proxy-port", 9891, "http server proxy port")
}
