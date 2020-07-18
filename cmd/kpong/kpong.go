package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/ryanhartje/kpong/pkg/kpong"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	kubeconfig string
	namespace  string
)

func newRootCmd(args []string) (*cobra.Command, error) {
	homedir, err := homedir.Dir()
	if err != nil {
		return &cobra.Command{}, err
	}
	defaultConfig := fmt.Sprintf("%s/.kube/config", homedir)

	v := viper.New()
	cmd := &cobra.Command{
		Use:   "kpong",
		Short: "A high stakes game of pong",
		Long:  `kpong is a kubernetes chaos game. Each player represents a pod. Lose and you lose your pod. Goodluck.`,
		Run: func(cmd *cobra.Command, args []string) {
			kpong.Start(kubeconfig, namespace)
		},
	}

	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", defaultConfig, "Path to the kubeconfig you'd like to use")
	cmd.Flags().StringVar(&namespace, "namespace", "", "The namespace you want to play out of. Leave blank for all namespaces. ex: kube-system")
	flags := cmd.Flags()
	flags.Parse(args)

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
