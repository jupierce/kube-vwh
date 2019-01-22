package main

import (
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
        klog.Fatalf("node-name is required")
      }
      if startOpts.keypath == "" {
        klog.Fatalf("node-name is required")
      }
        server.Serve(startOpts.certpath, startOpts.keypath)
      },
    }

    startOpts struct {
        certpath             string
        keypath              string
        port                 int
    }
)

func main() {
  var rootCmd = &cobra.Command{Use: "kube-vwh"}
  rootCmd.AddCommand(cmdServer)
  cmdServer.PersistentFlags().StringVar(&startOpts.certpath, "certpath", "", "Path to server cert")
  cmdServer.PersistentFlags().StringVar(&startOpts.keypath, "keypath", "", "Path to private server key")
  cmdServer.PersistentFlags().IntVar(&startOpts.port, "port", 443, "port to listen serve requests")
  rootCmd.Execute()
}
