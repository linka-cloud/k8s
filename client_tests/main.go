package main

import (
	"context"
	"os"
	"reflect"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"go.linka.cloud/k8s"
	appsv1 "go.linka.cloud/k8s/apps/v1"
	corev1 "go.linka.cloud/k8s/core/v1"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = k8s.AddToScheme(scheme)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	config.ContentType = "application/json"
	c, err := client.New(config, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		panic(err)
	}

	pl := &corev1.PodList{}
	if err := c.List(ctx, pl); err != nil {
		panic(err)
	}

	deploy := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       reflect.TypeOf(appsv1.Deployment{}).Name(),
			APIVersion: appsv1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "test",
		},
		Spec: &appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "test",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "test",
					},
				},
				Spec: &corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  k8s.Ref("test"),
							Image: k8s.Ref("traefik/whoami"),
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: k8s.Ref[int32](80),
								},
							},
						},
					},
				},
			},
		},
	}
	// Create the deployment with Patch/Apply
	if err := c.Patch(ctx, deploy.DeepCopy(), client.Apply, client.ForceOwnership, client.FieldOwner("test")); err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	// Re Apply with no changes
	if err := c.Patch(ctx, deploy.DeepCopy(), client.Apply, client.ForceOwnership, client.FieldOwner("test")); err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	n := deploy.DeepCopy()
	n.Spec.Template.Spec.Containers[0].Command = []string{"sleep", "100000"}
	n.Spec.Template.Spec.Containers[0].Image = k8s.Ref("alpine")
	if err := c.Patch(ctx, n, client.Apply, client.ForceOwnership, client.FieldOwner("test")); err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	// Delete the deployment
	if err := c.Delete(ctx, deploy); err != nil {
		panic(err)
	}
}
