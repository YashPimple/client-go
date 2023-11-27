package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/yashnileshpimple/.kube/config", "Location to config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//handle err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handle err
	}
	//pods are from corev1 resources()
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("keptn-lifecycle-toolkit-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handle err
	}
	fmt.Println("Pods from keptn-lifecycle-toolkit-system namespace are")
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("keptn-lifecycle-toolkit-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handle funct
	}
	fmt.Println("Deployments from keptn-lifecycle-toolkit-system namespace are")
	for _, deployment := range deployments.Items {
		fmt.Printf("%s", deployment.Name)
	}

}
