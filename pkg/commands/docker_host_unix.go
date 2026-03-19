//go:build !windows

package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	defaultDockerHost = "unix:///run/podman/podman.sock"
)

var osStat = os.Stat

func podmanSocketCandidates() []string {
	candidates := []string{}

	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		candidates = append(candidates, filepath.Join(runtimeDir, "podman", "podman.sock"))
	}

	candidates = append(candidates,
		filepath.Join("/run/user", fmt.Sprint(os.Getuid()), "podman", "podman.sock"),
		"/run/podman/podman.sock",
	)

	return candidates
}

func resolveDefaultDockerHost() string {
	for _, candidate := range podmanSocketCandidates() {
		if _, err := osStat(candidate); err == nil {
			return "unix://" + candidate
		}
	}

	return defaultDockerHost
}
