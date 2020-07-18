package kpong

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetRandomPod(t *testing.T) {
	clientset := fake.NewSimpleClientset()
	k8s := K8SClient{
		Clientset: clientset,
		Namespace: "",
	}
	pod, err := k8s.GetRandomPod()
	assert.NoErrorf(t, err, "Error when creating k8s client: %s", err)

	assert.NotEqual(t, pod.Name, "", "Pod name should not be blank")
}
