package main

import (
	"fmt"
	"os"

	"github.com/goshlanguage/kpong/pkg/kpong"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	host       bool
	hostIP     string
	kubeconfig string
	namespace  string
)

func newRootCmd(args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
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

	cmd.Flags().StringVar(&hostIP, "join", "127.0.0.1:27017", "join the IP specified for a match of pong, used for multiplayer")
	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig you'd like to use")
	cmd.Flags().StringVar(&namespace, "namespace", "", "The namespace you want to play out of. Leave blank for all namespaces. ex: kube-system")
	flags := cmd.Flags()
	flags.Parse(args)

	v := viper.New()
	v.BindEnv("kubeconfig")
	v.BindPFlag("kubeconfig", cmd.Flags().Lookup("kubeconfig"))

	kubeconfig = v.GetString("kubeconfig")
	fmt.Printf("Using kubeconfig: %s\n", kubeconfig)

	return cmd, nil
}

func main() {
	cmd, _ := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
