package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/yashnileshpimple/.kube/config", "Location to config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		//handle err
		fmt.Println("Error in building config fron flags\n", err.Error())
		config, err := rest.InClusterConfig()
		if err != nil{
			fmt.Println("Erro in getting inclusterconfig()\n", err.Error())
		}

	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//handle err
		fmt.Println("error in creating clientset\n", err.Error())
	}
	//pods are from corev1 resources()
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("keptn-lifecycle-toolkit-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handle err
		fmt.Println("error %s, while listing all the pods from Keptn-lifecycle-namespace\n", err.Error())
	}
	fmt.Println("Pods from keptn-lifecycle-toolkit-system namespace are")
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
	}

	deployments, err := clientset.AppsV1().Deployments("keptn-lifecycle-toolkit-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		//handle funct
		fmt.Println("listing deployments %s\n", err.Error())
	}
	fmt.Println("Deployments from keptn-lifecycle-toolkit-system namespace are")
	for _, deployment := range deployments.Items {
		fmt.Printf("%s", deployment.Name)
	}

}
