## Restrciting capabilties

[WIP]
I think runc provides an API for restricting capabilities assigned to a container, which in turn are used by docker, and kubernetes.
docker probably uses runC's [seccomp](https://github.com/opencontainers/runc/blob/7496a9682535c8ca143e981116f5b67463ad1d69/libcontainer/seccomp/seccomp_linux.go#L3)
