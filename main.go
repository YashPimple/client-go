package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/yashnileshpimple/.kube/config", "Location to our config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s building kubeconfig file", err.Error())
		//config, err := rest.InClusterConfig()
		if err != nil {
			fmt.Print("Error %s getting incluster config", err.Error())
		}
	}

	var namespace string
	fmt.Println("Enter the Namespace to get resources")
	fmt.Scan(&namespace)
	// use to interact with diff resources present in k8 it like an set of clients
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error %s creating clienset", err.Error())
	}

	// these code provide us info from all shared namespace
	informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	//these code will provide us only info of selected namespace
	informers.NewFilteredSharedInformerFactory(clientset, 10*time.Second, namespace, func(lo *metav1.ListOptions) {
		lo.LabelSelector = ""
		lo.APIVersion = ""
	})

	podinformer := informerfactory.Core().V1().Pods()
	podinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) {
			fmt.Println("add was called")
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Println("updated was called")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("delete was called")
		},
	})
	informerfactory.Start(wait.NeverStop)            //start
	informerfactory.WaitForCacheSync(wait.NeverStop) //waits for cache
	pod, err := podinformer.Lister().Pods(namespace).Get(namespace)
	if err != nil {
		fmt.Println("Unable to fetch resources")
	}
	fmt.Println(pod)

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
