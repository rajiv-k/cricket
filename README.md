# cricket

![Release Version](https://img.shields.io/badge/Release-v0.4.20-blue.svg)

<img src="./assets/cricket-1f997.png" alt="cricket logo" width="150" height="150"/>

`cricket` is a simple, _terribly lazy_, work-in-progress Open Container Initiative-based implementation of Kubernetes [Container Runtime Interface](https://kubernetes.io/docs/concepts/architecture/cri/).
When asked to do anything by its clients, it just "nods" its imaginary head and silently ignores it (_insert crickets noise here_). It stores some state which helps it to lie in the future. This is by design.

It's **NOT** meant to be used for anything serious at all.


## But why?
tl;dr: because we can!

Besides, it helps one to see what is happening under the hood and gain a deeper understanding of the layers beneath the kubelet.

I created this mostly for fun ~~and no profit~~.

## Build

```sh
$ make cricket
```


## Run

```sh
$ build/cricket
```

Once the server is running, it can be interacted with by the [`crictl`](https://github.com/kubernetes-sigs/cri-tools/blob/master/docs/crictl.md) command-line tool.

```sh
$ crictl --runtime-endpoint unix:///tmp/cricket.sock version
Version:  v0.4.20
RuntimeName:  cricket
RuntimeVersion:  v0.4.20
RuntimeApiVersion:  v1alpha1
```

## Supported commands

- [x] Version
- [ ] RunPodSandbox
- [ ] StopPodSandbox
- [ ] RemovePodSandbox
- [ ] PodSandboxStatus
- [ ] ListPodSandbox
- [ ] CreateContainer
- [ ] StartContainer
- [ ] StopContainer
- [ ] RemoveContainer
- [ ] ListContainers
- [ ] ContainerStatus
- [ ] UpdateContainerResources
- [ ] ReopenContainerLog
- [ ] ExecSync
- [ ] Exec
- [ ] Attach
- [ ] PortForward
- [ ] ContainerStats
- [ ] ListContainerStats
- [ ] PodSandboxStats
- [ ] ListPodSandboxStats
- [ ] UpdateRuntimeConfig
- [ ] Status
- [ ] CheckpointContainer
- [ ] GetContainerEvents
- [ ] ListMetricDescriptors
- [ ] ListPodSandboxMetrics
- [ ] RuntimeConfig
- [ ] ListImages
- [ ] ImageStatus
- [ ] PullImage
- [ ] RemoveImage
- [ ] ImageFsInfo
