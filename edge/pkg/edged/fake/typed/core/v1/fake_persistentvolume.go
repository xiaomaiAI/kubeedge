package v1

import (
	"github.com/kubeedge/kubeedge/edge/pkg/metamanager/client"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
)

// FakePersistentVolumes implements PersistentVolumeInterface
type FakePersistentVolumes struct {
	fakecorev1.FakePersistentVolumes
	MetaClient client.CoreInterface
}

// Get takes name of the persistentVolume, and returns the corresponding persistentVolume object, and an error if there is any.
func (c *FakePersistentVolumes) Get(name string, options metav1.GetOptions) (result *corev1.PersistentVolume, err error) {
	return c.MetaClient.PersistentVolumes(metav1.NamespaceDefault).Get(name, options)
}
