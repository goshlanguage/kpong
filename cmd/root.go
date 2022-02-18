package cmd

import (
	"fmt"

	"github.com/goshlanguage/kpong/pkg/kpong"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	host       bool
	hostIP     string
	kubeconfig string
	namespace  string
	rootCmd    = &cobra.Command{
		Use:   "kpong",
		Short: "A high stakes game of pong",
		Long:  `kpong is a kubernetes chaos game. Each player represents a pod. Lose and you lose your pod. Goodluck.`,
		Run: func(cmd *cobra.Command, args []string) {
			conf := kpong.GameConfig{
				Host:         host,
				HostIP:       hostIP,
				Kubeconfig:   kubeconfig,
				Namespace:    namespace,
				ScreenHeight: 800,
				ScreenWidth:  1024,
			}
			kpong.MainMenu(conf)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&hostIP, "join", "127.0.0.1:27017", "join the IP specified for a match of pong, used for multiplayer")
	rootCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig you'd like to use")
	rootCmd.Flags().StringVar(&namespace, "namespace", "", "The namespace you want to play out of. Leave blank for all namespaces. ex: kube-system")

	v := viper.New()
	v.BindEnv("kubeconfig")
	v.BindPFlag("kubeconfig", rootCmd.Flags().Lookup("kubeconfig"))

	kubeconfig = v.GetString("kubeconfig")
	fmt.Printf("Using kubeconfig: %s\n", kubeconfig)
}
