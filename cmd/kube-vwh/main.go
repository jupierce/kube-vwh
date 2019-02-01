package main

import (
  "flag"
  "github.com/spf13/cobra"
  "github.com/michaelgugino/kube-vwh/pkg/server"
  "k8s.io/klog"
)

var (
    cmdServer = &cobra.Command{
      Use:   "server",
      Short: "run the server",
      Long: `run the server.`,
      Args: cobra.MinimumNArgs(0),
      Run: func(cmd *cobra.Command, args []string) {
      if startOpts.certpath == "" {
        klog.Fatalf("certpath is required")
      }
      if startOpts.keypath == "" {
        klog.Fatalf("keypath is required")
      }
        server.Serve(startOpts.certpath, startOpts.keypath, startOpts.port)
      },
    }

    startOpts struct {
        certpath             string
        keypath              string
        port                 int
    }
)

var rootCmd = &cobra.Command{Use: "kube-vwh"}

func init() {
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}

func main() {

  rootCmd.AddCommand(cmdServer)
  cmdServer.PersistentFlags().StringVar(&startOpts.certpath, "certpath", "", "Path to server cert")
  cmdServer.PersistentFlags().StringVar(&startOpts.keypath, "keypath", "", "Path to private server key")
  cmdServer.PersistentFlags().IntVar(&startOpts.port, "port", 8443, "port to listen serve requests")
  rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
  flag.Set("logtostderr", "true")
  flag.Parse()
  rootCmd.Execute()
}
