package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/config"
	"k8s.io/client-go/kubernetes"
)

func main() {
	// load the kubeconfig file
	config, err := config.LoadFromDefaultLocation()
	if err != nil {
		log.Fatal(err)
	}

	// create a kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// get all pods in the default namespace
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// print the name of each pod
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
