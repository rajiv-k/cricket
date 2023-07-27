package server

import (
	"context"
	"strings"

	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
)

const (
	ServerVersion     = "v0.4.20"
	RuntimeApiVersion = "v1alpha1"
	RuntimeName       = "cricket"
	ContainerRunning  = "RUNNING"
	ContainerStopped  = "STOPPED"
	ContainerExited   = "EXITED"
)

type RuntimeServer struct {
	Ver        string
	Containers map[string]ContainerInfo
}

type ContainerInfo struct {
	Name       string
	Image      string
	Entrypoint string
	Status     string
}

type ImageServer struct {
	Ver string
}

func NewRuntimeServer() *RuntimeServer {
	return &RuntimeServer{
		Ver:        ServerVersion,
		Containers: make(map[string]ContainerInfo),
	}
}

func NewImageServer() *ImageServer {
	return &ImageServer{Ver: ServerVersion}
}

func (r *RuntimeServer) Version(_ context.Context, request *runtimeapi.VersionRequest) (*runtimeapi.VersionResponse, error) {
	resp := &runtimeapi.VersionResponse{
		Version:           r.Ver,
		RuntimeApiVersion: RuntimeApiVersion,
		RuntimeName:       RuntimeName,
		RuntimeVersion:    r.Ver,
	}
	return resp, nil
}

func (r *RuntimeServer) Status(_ context.Context, _ *runtimeapi.StatusRequest) (*runtimeapi.StatusResponse, error) {
	resp := &runtimeapi.StatusResponse{
		Status: &runtimeapi.RuntimeStatus{
			Conditions: []*runtimeapi.RuntimeCondition{
				{
					Status: true,
					Type:   runtimeapi.RuntimeReady,
				},
				{
					Status: true,
					Type:   runtimeapi.NetworkReady,
				},
			},
		},
	}
	return resp, nil
}

func (r *RuntimeServer) Attach(_ context.Context, req *runtimeapi.AttachRequest) (*runtimeapi.AttachResponse, error) {
	return &runtimeapi.AttachResponse{}, nil
}

func (r *RuntimeServer) CheckpointContainer(context.Context, *runtimeapi.CheckpointContainerRequest) (*runtimeapi.CheckpointContainerResponse, error) {
	return &runtimeapi.CheckpointContainerResponse{}, nil
}

func (r *RuntimeServer) RunPodSandbox(context.Context, *runtimeapi.RunPodSandboxRequest) (*runtimeapi.RunPodSandboxResponse, error) {
	return &runtimeapi.RunPodSandboxResponse{}, nil
}

func (r *RuntimeServer) StopPodSandbox(context.Context, *runtimeapi.StopPodSandboxRequest) (*runtimeapi.StopPodSandboxResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) RemovePodSandbox(context.Context, *runtimeapi.RemovePodSandboxRequest) (*runtimeapi.RemovePodSandboxResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) PodSandboxStatus(_ context.Context, _ *runtimeapi.PodSandboxStatusRequest) (*runtimeapi.PodSandboxStatusResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ListPodSandbox(_ context.Context, filter *runtimeapi.ListPodSandboxRequest) (*runtimeapi.ListPodSandboxResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) PortForward(context.Context, *runtimeapi.PortForwardRequest) (*runtimeapi.PortForwardResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) CreateContainer(_ context.Context, req *runtimeapi.CreateContainerRequest) (*runtimeapi.CreateContainerResponse, error) {
	containerName := req.GetConfig().GetMetadata().GetName()
	entry := ContainerInfo{
		Name:       containerName,
		Image:      req.GetConfig().GetImage().String(),
		Entrypoint: strings.Join(req.GetConfig().Command, ""),
		Status:     ContainerRunning,
	}
	r.Containers[containerName] = entry
	resp := &runtimeapi.CreateContainerResponse{
		ContainerId: containerName,
	}
	return resp, nil
}

func (r *RuntimeServer) StartContainer(_ context.Context, _ *runtimeapi.StartContainerRequest) (*runtimeapi.StartContainerResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) StopContainer(_ context.Context, _ *runtimeapi.StopContainerRequest) (*runtimeapi.StopContainerResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) RemoveContainer(_ context.Context, _ *runtimeapi.RemoveContainerRequest) (*runtimeapi.RemoveContainerResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ListContainers(_ context.Context, _ *runtimeapi.ListContainersRequest) (*runtimeapi.ListContainersResponse, error) {
	containers := make([]*runtimeapi.Container, 0, len(r.Containers))
	for k, v := range r.Containers {
		containers = append(containers, &runtimeapi.Container{
			Metadata: &runtimeapi.ContainerMetadata{
				Name: k,
			},
			State: containerStateFromStatus(v.Status),
		})
	}
	resp := &runtimeapi.ListContainersResponse{
		Containers: containers,
	}
	return resp, nil
}

func (r *RuntimeServer) ContainerStatus(_ context.Context, _ *runtimeapi.ContainerStatusRequest) (*runtimeapi.ContainerStatusResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) UpdateContainerResources(context.Context, *runtimeapi.UpdateContainerResourcesRequest) (*runtimeapi.UpdateContainerResourcesResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ExecSync(_ context.Context, request *runtimeapi.ExecSyncRequest) (*runtimeapi.ExecSyncResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) Exec(context.Context, *runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) UpdateRuntimeConfig(_ context.Context, _ *runtimeapi.UpdateRuntimeConfigRequest) (*runtimeapi.UpdateRuntimeConfigResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ContainerStats(context.Context, *runtimeapi.ContainerStatsRequest) (*runtimeapi.ContainerStatsResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ListContainerStats(_ context.Context, request *runtimeapi.ListContainerStatsRequest) (*runtimeapi.ListContainerStatsResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) PodSandboxStats(_ context.Context, request *runtimeapi.PodSandboxStatsRequest) (*runtimeapi.PodSandboxStatsResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ListPodSandboxStats(_ context.Context, _ *runtimeapi.ListPodSandboxStatsRequest) (*runtimeapi.ListPodSandboxStatsResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ReopenContainerLog(_ context.Context, _ *runtimeapi.ReopenContainerLogRequest) (*runtimeapi.ReopenContainerLogResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) GetContainerEvents(*runtimeapi.GetEventsRequest, runtimeapi.RuntimeService_GetContainerEventsServer) error {
	return nil
}

func (r *RuntimeServer) ListMetricDescriptors(_ context.Context, _ *runtimeapi.ListMetricDescriptorsRequest) (*runtimeapi.ListMetricDescriptorsResponse, error) {
	return nil, nil
}

func (r *RuntimeServer) ListPodSandboxMetrics(_ context.Context, request *runtimeapi.ListPodSandboxMetricsRequest) (*runtimeapi.ListPodSandboxMetricsResponse, error) {
	return nil, nil
}

// Image Service

func (r *ImageServer) InjectError(f string, err error) {}
func (r *ImageServer) ListImages(_ context.Context, request *runtimeapi.ListImagesRequest) (*runtimeapi.ListImagesResponse, error) {
	return nil, nil
}
func (r *ImageServer) ImageStatus(_ context.Context, _ *runtimeapi.ImageStatusRequest) (*runtimeapi.ImageStatusResponse, error) {
	return nil, nil
}
func (r *ImageServer) PullImage(_ context.Context, request *runtimeapi.PullImageRequest) (*runtimeapi.PullImageResponse, error) {
	return nil, nil
}
func (r *ImageServer) RemoveImage(_ context.Context, image *runtimeapi.RemoveImageRequest) (*runtimeapi.RemoveImageResponse, error) {
	return nil, nil
}
func (r *ImageServer) ImageFsInfo(_ context.Context, request *runtimeapi.ImageFsInfoRequest) (*runtimeapi.ImageFsInfoResponse, error) {
	return nil, nil
}
