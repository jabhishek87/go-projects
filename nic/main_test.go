package main

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestFetchVfPodInfo(t *testing.T) {
	clientset := fake.NewSimpleClientset(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "pod1"},
	})
	// get all pods in the default namespace
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		t.Error(err)
	}

	// check that there is only one pod
	if len(pods.Items) != 1 {
		t.Errorf("expected 1 pod, got %d", len(pods.Items))
	}

	// check that the name of the pod is "pod1"
	if pods.Items[0].Name != "pod1" {
		t.Errorf("expected pod name to be \"pod1\", got \"%s\"", pods.Items[0].Name)
	}
}
