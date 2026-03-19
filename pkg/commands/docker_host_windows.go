//go:build windows

package commands

const (
	defaultDockerHost = "npipe:////./pipe/docker_engine"
)

func resolveDefaultDockerHost() string {
	return defaultDockerHost
}
