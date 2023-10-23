package main

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	config, _    = clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	clientset, _ = kubernetes.NewForConfig(config)
)

func main() {
	log.Info("operator starting")
	list, err := clientset.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range list.Items {
		log.Info(item.Name)
	}
}
