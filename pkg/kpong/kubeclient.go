package kpong

import (
	"fmt"
	"math/rand"
	"os"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func newK8SClient(kubeconfig string) (*kubernetes.Clientset, error) {
	if kubeconfig == "" {
		kubeconfig = fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Encountered error when setting up kubernetes clientcmd: %s\n", err.Error())
		return &kubernetes.Clientset{}, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Encountered error when setting up kubernetes clientset: %s\n", err.Error())
		return &kubernetes.Clientset{}, err
	}
	return clientset, nil
}

// GetRandomPod selects a pod from a particular namespace to be on the hook
// If an empty string is supplied for namespace, all namespaces will be enumerated
func GetRandomPod(kubeconfig string, namespace string) (v1.Pod, error) {
	client, err := newK8SClient(kubeconfig)
	if err != nil {
		return v1.Pod{}, err
	}

	pods, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Encountered error when setting up kubernetes clientset: %s\n", err.Error())
		return v1.Pod{}, err
	}

	die := rand.Intn(len(pods.Items))
	randoPod := pods.Items[die]
	return randoPod, nil
}
