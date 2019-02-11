package edb

import (
	"context"
	"github.com/UnlawfulMonad/edb-operator/pkg/apis/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CanUseDB(c client.Client, namespace string, database *v1alpha1.ExternalDatabase) (bool, error) {
	selector := labels.Set(database.Spec.Selector.MatchLabels).AsSelector()

	nsls := &corev1.NamespaceList{}
	listOptions := &client.ListOptions{LabelSelector: selector}
	err := c.List(context.TODO(), listOptions, nsls)
	if err != nil {
		return false, err
	}

	for _, ns := range nsls.Items {
		// Check if the database is one of the namespaces that the ExternalDB covers
		if ns.Name == namespace {
			return true, nil
		}
	}

	return false, nil
}
