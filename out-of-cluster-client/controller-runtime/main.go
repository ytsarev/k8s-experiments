package main

import (
	"context"
	"flag"
	"runtime"

	"k8s.io/client-go/tools/clientcmd"

	cnatv1alpha1 "github.com/.../cnat/cnat-kubebuilder/pkg/apis/cnat/v1alpha1"
	ohmyglbv1beta1 "github.com/AbsaOSS/ohmyglb/pkg/apis/ohmyglb/v1beta1"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	kubeconfig = flag.String("kubeconfig", "~/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	crScheme := runtime.NewScheme()
	cnatv1alpha1.AddToScheme(crScheme)
	ohmyglbv1beta1.AddToScheme(crScheme)

	cl, _ := runtimeclient.New(config, client.Options{
		Scheme: crScheme,
	})
	list := &cnatv1alpha1.AtList{}
	err := cl.List(context.TODO(), client.InNamespace("default"), list)

}
