package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"
	informers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	k8sruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Set up Kubernetes client configuration
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = fmt.Sprintf("%s/.kube/config", home)
	}
	configFlag := flag.String("kubeconfig", kubeconfig, "path to Kubernetes config file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *configFlag)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Stop signal for the informer
	stopper := make(chan struct{})
	defer close(stopper)

	// Set resync interval to 10 minutes
	ResyncInterval := 10 * time.Minute

	// Create a shared informer factory with resync period
	factory := informers.NewSharedInformerFactoryWithOptions(clientSet, ResyncInterval)

	// Create a pod informer
	podInformer := factory.Core().V1().Pods().Informer()

	// Handle panics gracefully
	defer k8sruntime.HandleCrash()

	// Add event handlers to the pod informer
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(podObj interface{}) {
			pod, ok := podObj.(*corev1.Pod)
			if ok {
				fmt.Printf("Added Pod: %s. Current status: %s\n", pod.GetName(), pod.Status.Phase)
			}
		},
		UpdateFunc: func(oldPodObj, newPodObj interface{}) {
			// Old version of the Pod Object
			oldPod, okOld := oldPodObj.(*corev1.Pod)
			// New version of the Pod Object
			newPod, okNew := newPodObj.(*corev1.Pod)

			if okOld && okNew {
				// If resource version is different, there was a change with the pod
				if oldPod.ResourceVersion != newPod.ResourceVersion {
					fmt.Printf("Updated Pod: %s. Current status: %s\n", newPod.GetName(), newPod.Status.Phase)
				} else {
					fmt.Printf("Nothing to update for Pod: %s\n", newPod.GetName())
				}
			}
		},
		DeleteFunc: func(podObj interface{}) {
			pod, ok := podObj.(*corev1.Pod)
			if ok {
				fmt.Printf("Deleted Pod: %s\n", pod.GetName())
			}
		},
	})

	// Start the shared informer factory
	go factory.Start(stopper)

	// Wait for termination signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	log.Println("Shutting down...")
}
