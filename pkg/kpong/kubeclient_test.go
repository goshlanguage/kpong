package kpong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewK8SClient(t *testing.T) {
	_, err := newK8SClient("")
	assert.NoErrorf(t, err, "Error when creating k8s client: %s", err)
}

func TestGetRandomPod(t *testing.T) {
	pod, err := GetRandomPod("", "")
	assert.NoErrorf(t, err, "Error when creating k8s client: %s", err)

	assert.NotEqual(t, pod.Name, "", "Pod name should not be blank")
}
