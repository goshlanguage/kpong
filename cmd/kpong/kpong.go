package main

import (
	"fmt"
	"os"

	"github.com/ryanhartje/kpong/pkg/kpong"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	kubeconfig string
	namespace  string
)

func init() {
	viper.AutomaticEnv()
}

func newRootCmd(args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "kpong",
		Short: "A high stakes game of pong",
		Long:  `kpong is a kubernetes chaos game. Each player represents a pod. Lose and you lose your pod. Goodluck.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Check kubeconfig. If we don't pass it in from the command line, grab from env via Viper
			if kubeconfig == "" && viper.GetString("kubeconfig") != "" {
				kubeconfig = viper.GetString("kubeconfig")
			}

			kpong.Start(kubeconfig, namespace)
		},
	}

	cmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig you'd like to use")
	cmd.PersistentFlags().StringVar(&namespace, "namespace", "", "The namespace you want to play out of. Leave blank for all namespaces. ex: kube-system")
	flags := cmd.PersistentFlags()
	flags.Parse(args)

	viper.BindPFlag("kubeconfig", cmd.PersistentFlags().Lookup("kubeconfig"))
	viper.BindEnv("kubeconfig", "KUBECONFIG")

	return cmd, nil
}

func main() {
	cmd, _ := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
