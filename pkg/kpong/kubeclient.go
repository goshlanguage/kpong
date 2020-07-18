package kpong

import (
	"fmt"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/clientcmd"
)

// K8SClient holds a kubernetes clientset, set as interface for testability
type K8SClient struct {
	Clientset kubernetes.Interface
	Failed    bool   // Failed indicates if there have been errors that keep our clientset from being functional
	Namespace string // Namespace represents the kubernetes namespace we should be operating in, "" for all namespaces
}

func newK8SClient(kubeconfig string, namespace string) *K8SClient {
	var failed bool
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Encountered error when setting up kubernetes clientcmd: %s\n", err.Error())
		failed = true
	}

	var clientset *kubernetes.Clientset
	if !failed {
		// create the clientset
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			fmt.Printf("Encountered error when setting up kubernetes clientset: %s\n", err.Error())
			failed = true
		}
	}

	if failed {
		return &K8SClient{
			fake.NewSimpleClientset(),
			true,
			"",
		}
	}
	return &K8SClient{
		Clientset: clientset,
		Namespace: namespace,
	}
}

// GetRandomPod selects a pod from a particular namespace to be on the hook
// If an empty string is supplied for namespace, all namespaces will be enumerated
// GetRandom pod tries to fail gracefully by providing a fake object if the client fails
func (k8s *K8SClient) GetRandomPod() (*v1.Pod, error) {
	if k8s.Failed {
		return &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("fakey-mcfake-fake-%v", rand.Intn(64)),
					Namespace: "fakespace",
				},
			},
			fmt.Errorf("The kubernetes client failed to initialize")
	}

	pods, err := k8s.Clientset.CoreV1().Pods(k8s.Namespace).List(metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Encountered error when setting up kubernetes clientset: %s\n", err.Error())
		return &v1.Pod{}, err
	}

	if len(pods.Items) == 0 {
		return &v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "no-pods-found",
					Namespace: "",
				},
			},
			nil
	}
	die := rand.Intn(len(pods.Items))
	randoPod := &pods.Items[die]
	return randoPod, nil
}

// CyclePod is responsible for removing the given pod, and returning a new one
func (k8s *K8SClient) CyclePod(pod *v1.Pod) (*v1.Pod, error) {
	if !k8s.Failed {
		deletePolicy := metav1.DeletePropagationForeground
		k8s.Clientset.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		})

		newPod, err := k8s.GetRandomPod()
		if err != nil {
			return &v1.Pod{}, err
		}
		return newPod, err
	}
	return &v1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("fakey-mcfake-fake-%v", rand.Intn(64)),
				Namespace: "fakespace",
			},
		},
		fmt.Errorf("The kubernetes client failed to initialize")
}
