package server

import (
	"github.com/sirupsen/logrus"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func containerStateFromStatus(status string) runtimeapi.ContainerState {
	switch status {
	case ContainerRunning:
		return runtimeapi.ContainerState_CONTAINER_RUNNING
	default:
		logrus.Warningf("got invalid container status: %q", status)
		return runtimeapi.ContainerState_CONTAINER_UNKNOWN
	}
}
