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
	kubeconfig := flag.String("kubeconfig", "/Users/yashnileshpimple/.kube/config", "Location to our config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Unable to find kubeconfig file")
	}

	var namespace string
	fmt.Println("Enter the Namespace to get resources")
	fmt.Scan(&namespace)
	// use to interact with diff resources present in k8 it like an set of clients
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("clienset not running")
	}

	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Unable to find resources pod")
	}
	fmt.Println("Pods in following namespace are:")
	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Unable to find resource deployment")
	}
	fmt.Println("Deployments in following namespace are:")
	for _, deployment := range deployments.Items {
		fmt.Printf("%s\n", deployment.Name)
	}

	services, err := clientset.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("Unable to find resource service")
	}
	fmt.Println("Service in following namespace are:")
	for _, service := range services.Items {
		fmt.Printf("%s\n", service.Name)
	}
	// fmt.Println(clientset)
	// fmt.Println(config)
}
