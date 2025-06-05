package main

import (
	goflags "flag"
	"github.com/clarechu/infrapulse/cmd"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"os"
)

func init() {
	//	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(goflags.CommandLine)
}

func main() {
	rootCmd := cmd.GetRootCmd(os.Args[1:])
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
		os.Exit(-1)
	}
}
