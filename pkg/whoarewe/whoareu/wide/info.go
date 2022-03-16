package wide

import (
	"github.com/watergist/k8s-manifests/pkg/whoarewe/listener"
	"os"
	"path"
)

type PodProperties struct {
	Server      *listener.Listener
	Name        string
	Namespace   string
	Labels      map[string]string
	Annotations map[string]string
	Error       []error
}

func GetPodProperties(listenerProperties *listener.Listener) PodProperties {
	var err error
	podProperties := PodProperties{
		Server:    listenerProperties,
		Name:      os.Getenv("POD_NAME"),
		Namespace: os.Getenv("POD_NAMESPACE"),
	}
	if proprtiesPath := os.Getenv("POD_PROPERTIES_PATH"); proprtiesPath != "" {
		podProperties.Labels, err = DrawPodProperties(path.Join(proprtiesPath, "labels"))
		if err != nil {
			podProperties.Error = append(podProperties.Error, err)
		}
		delete(podProperties.Labels, "pod-template-hash")
		podProperties.Annotations, err = DrawPodProperties(path.Join(proprtiesPath, "annotations"))
		if err != nil {
			podProperties.Error = append(podProperties.Error, err)
		}
	}
	return podProperties
}
