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
func GetRandomPod(clientset *kubernetes.Clientset, namespace string) (*v1.Pod, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Encountered error when setting up kubernetes clientset: %s\n", err.Error())
		return &v1.Pod{}, err
	}

	die := rand.Intn(len(pods.Items))
	randoPod := &pods.Items[die]
	return randoPod, nil
}

// CyclePod is responsible for removing the given pod, and returning a new one
func CyclePod(clientset *kubernetes.Clientset, pod *v1.Pod) (*v1.Pod, error) {
	deletePolicy := metav1.DeletePropagationForeground
	clientset.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	newPod, err := GetRandomPod(clientset, "")
	if err != nil {
		return &v1.Pod{}, err
	}
	return newPod, err
}
