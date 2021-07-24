package main

import (
	"context"
	"flag"
	"path/filepath"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

const (
	Service   = "ingress-nginx"
	NameSpace = "kube-system"

	defaultConfig = ".kube/config"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	if len(externalIP) == 0 {
		klog.Fatal("empty external ip")
	}

	// Build the kubeconfig from kubeConfig
	if len(kubeconfig) == 0 {
		kubeconfig = defaultConfig
	}

	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), defaultConfig))
	if err != nil {
		klog.Fatal(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	service, err := clientSet.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		klog.Fatal(err)
	}

	for _, extIp := range service.Spec.ExternalIPs {
		if externalIP == extIp {
			klog.Infof("externalIp %q already binded on service %s/%s", externalIP, namespace, name)
			return
		}
	}

	service.Spec.ExternalIPs = []string{externalIP}
	service.Spec.Type = v1.ServiceTypeLoadBalancer
	if _, err = clientSet.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{}); err != nil {
		klog.Fatal(err)
	}
	klog.Infof("externalIp %q binded on service %s/%s", externalIP, namespace, name)
}

var (
	kubeconfig string
	name       string
	namespace  string
	externalIP string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&name, "name", Service, "service name")
	flag.StringVar(&namespace, "namespace", NameSpace, "service namespace")
	flag.StringVar(&externalIP, "externalip", "", "service external ip")
}
