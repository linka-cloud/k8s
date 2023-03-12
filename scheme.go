package k8s

import (
	"k8s.io/apimachinery/pkg/runtime"
	admissionv1 "go.linka.cloud/k8s/admission/v1"
	admissionv1beta1 "go.linka.cloud/k8s/admission/v1beta1"
	admissionregistrationv1 "go.linka.cloud/k8s/admissionregistration/v1"
	admissionregistrationv1alpha1 "go.linka.cloud/k8s/admissionregistration/v1alpha1"
	admissionregistrationv1beta1 "go.linka.cloud/k8s/admissionregistration/v1beta1"
	apidiscoveryv2 "go.linka.cloud/k8s/apidiscovery/v2"
	apidiscoveryv2beta1 "go.linka.cloud/k8s/apidiscovery/v2beta1"
	apiserverinternalv1alpha1 "go.linka.cloud/k8s/apiserverinternal/v1alpha1"
	appsv1 "go.linka.cloud/k8s/apps/v1"
	appsv1beta1 "go.linka.cloud/k8s/apps/v1beta1"
	appsv1beta2 "go.linka.cloud/k8s/apps/v1beta2"
	authenticationv1 "go.linka.cloud/k8s/authentication/v1"
	authenticationv1alpha1 "go.linka.cloud/k8s/authentication/v1alpha1"
	authenticationv1beta1 "go.linka.cloud/k8s/authentication/v1beta1"
	authorizationv1 "go.linka.cloud/k8s/authorization/v1"
	authorizationv1beta1 "go.linka.cloud/k8s/authorization/v1beta1"
	autoscalingv1 "go.linka.cloud/k8s/autoscaling/v1"
	autoscalingv2 "go.linka.cloud/k8s/autoscaling/v2"
	autoscalingv2beta1 "go.linka.cloud/k8s/autoscaling/v2beta1"
	autoscalingv2beta2 "go.linka.cloud/k8s/autoscaling/v2beta2"
	batchv1 "go.linka.cloud/k8s/batch/v1"
	batchv1beta1 "go.linka.cloud/k8s/batch/v1beta1"
	certificatesv1 "go.linka.cloud/k8s/certificates/v1"
	certificatesv1alpha1 "go.linka.cloud/k8s/certificates/v1alpha1"
	certificatesv1beta1 "go.linka.cloud/k8s/certificates/v1beta1"
	coordinationv1 "go.linka.cloud/k8s/coordination/v1"
	coordinationv1alpha1 "go.linka.cloud/k8s/coordination/v1alpha1"
	coordinationv1beta1 "go.linka.cloud/k8s/coordination/v1beta1"
	corev1 "go.linka.cloud/k8s/core/v1"
	discoveryv1 "go.linka.cloud/k8s/discovery/v1"
	discoveryv1beta1 "go.linka.cloud/k8s/discovery/v1beta1"
	eventsv1 "go.linka.cloud/k8s/events/v1"
	eventsv1beta1 "go.linka.cloud/k8s/events/v1beta1"
	extensionsv1beta1 "go.linka.cloud/k8s/extensions/v1beta1"
	flowcontrolv1 "go.linka.cloud/k8s/flowcontrol/v1"
	flowcontrolv1beta1 "go.linka.cloud/k8s/flowcontrol/v1beta1"
	flowcontrolv1beta2 "go.linka.cloud/k8s/flowcontrol/v1beta2"
	flowcontrolv1beta3 "go.linka.cloud/k8s/flowcontrol/v1beta3"
	imagepolicyv1alpha1 "go.linka.cloud/k8s/imagepolicy/v1alpha1"
	networkingv1 "go.linka.cloud/k8s/networking/v1"
	networkingv1alpha1 "go.linka.cloud/k8s/networking/v1alpha1"
	networkingv1beta1 "go.linka.cloud/k8s/networking/v1beta1"
	nodev1 "go.linka.cloud/k8s/node/v1"
	nodev1alpha1 "go.linka.cloud/k8s/node/v1alpha1"
	nodev1beta1 "go.linka.cloud/k8s/node/v1beta1"
	policyv1 "go.linka.cloud/k8s/policy/v1"
	policyv1beta1 "go.linka.cloud/k8s/policy/v1beta1"
	rbacv1 "go.linka.cloud/k8s/rbac/v1"
	rbacv1alpha1 "go.linka.cloud/k8s/rbac/v1alpha1"
	rbacv1beta1 "go.linka.cloud/k8s/rbac/v1beta1"
	resourcev1alpha3 "go.linka.cloud/k8s/resource/v1alpha3"
	schedulingv1 "go.linka.cloud/k8s/scheduling/v1"
	schedulingv1alpha1 "go.linka.cloud/k8s/scheduling/v1alpha1"
	schedulingv1beta1 "go.linka.cloud/k8s/scheduling/v1beta1"
	storagev1 "go.linka.cloud/k8s/storage/v1"
	storagev1alpha1 "go.linka.cloud/k8s/storage/v1alpha1"
	storagev1beta1 "go.linka.cloud/k8s/storage/v1beta1"
	storagemigrationv1alpha1 "go.linka.cloud/k8s/storagemigration/v1alpha1"
)

var groups = []runtime.SchemeBuilder{
	admissionv1.SchemeBuilder,
	admissionv1beta1.SchemeBuilder,
	admissionregistrationv1.SchemeBuilder,
	admissionregistrationv1alpha1.SchemeBuilder,
	admissionregistrationv1beta1.SchemeBuilder,
	apidiscoveryv2.SchemeBuilder,
	apidiscoveryv2beta1.SchemeBuilder,
	apiserverinternalv1alpha1.SchemeBuilder,
	appsv1.SchemeBuilder,
	appsv1beta1.SchemeBuilder,
	appsv1beta2.SchemeBuilder,
	authenticationv1.SchemeBuilder,
	authenticationv1alpha1.SchemeBuilder,
	authenticationv1beta1.SchemeBuilder,
	authorizationv1.SchemeBuilder,
	authorizationv1beta1.SchemeBuilder,
	autoscalingv1.SchemeBuilder,
	autoscalingv2.SchemeBuilder,
	autoscalingv2beta1.SchemeBuilder,
	autoscalingv2beta2.SchemeBuilder,
	batchv1.SchemeBuilder,
	batchv1beta1.SchemeBuilder,
	certificatesv1.SchemeBuilder,
	certificatesv1alpha1.SchemeBuilder,
	certificatesv1beta1.SchemeBuilder,
	coordinationv1.SchemeBuilder,
	coordinationv1alpha1.SchemeBuilder,
	coordinationv1beta1.SchemeBuilder,
	corev1.SchemeBuilder,
	discoveryv1.SchemeBuilder,
	discoveryv1beta1.SchemeBuilder,
	eventsv1.SchemeBuilder,
	eventsv1beta1.SchemeBuilder,
	extensionsv1beta1.SchemeBuilder,
	flowcontrolv1.SchemeBuilder,
	flowcontrolv1beta1.SchemeBuilder,
	flowcontrolv1beta2.SchemeBuilder,
	flowcontrolv1beta3.SchemeBuilder,
	imagepolicyv1alpha1.SchemeBuilder,
	networkingv1.SchemeBuilder,
	networkingv1alpha1.SchemeBuilder,
	networkingv1beta1.SchemeBuilder,
	nodev1.SchemeBuilder,
	nodev1alpha1.SchemeBuilder,
	nodev1beta1.SchemeBuilder,
	policyv1.SchemeBuilder,
	policyv1beta1.SchemeBuilder,
	rbacv1.SchemeBuilder,
	rbacv1alpha1.SchemeBuilder,
	rbacv1beta1.SchemeBuilder,
	resourcev1alpha3.SchemeBuilder,
	schedulingv1.SchemeBuilder,
	schedulingv1alpha1.SchemeBuilder,
	schedulingv1beta1.SchemeBuilder,
	storagev1.SchemeBuilder,
	storagev1alpha1.SchemeBuilder,
	storagev1beta1.SchemeBuilder,
	storagemigrationv1alpha1.SchemeBuilder,
}

func AddToScheme(scheme *runtime.Scheme) error {
	for _, group := range groups {
		if err := group.AddToScheme(scheme); err != nil {
			return err
		}
	}
	return nil
}
