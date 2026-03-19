//go:build !windows

package commands

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveDefaultDockerHostPrefersExistingPodmanSocket(t *testing.T) {
	originalRuntimeDir := os.Getenv("XDG_RUNTIME_DIR")
	originalStat := osStat
	defer func() {
		_ = os.Setenv("XDG_RUNTIME_DIR", originalRuntimeDir)
		osStat = originalStat
	}()

	_ = os.Setenv("XDG_RUNTIME_DIR", "/tmp/test-runtime")
	expectedPath := filepath.Join("/tmp/test-runtime", "podman", "podman.sock")
	osStat = func(name string) (os.FileInfo, error) {
		if name == expectedPath {
			return nil, nil
		}
		return nil, os.ErrNotExist
	}

	if actual := resolveDefaultDockerHost(); actual != "unix://"+expectedPath {
		t.Fatalf("expected %s but got %s", "unix://"+expectedPath, actual)
	}
}

func TestResolveDefaultDockerHostFallsBackToRootfulPodmanSocket(t *testing.T) {
	originalRuntimeDir := os.Getenv("XDG_RUNTIME_DIR")
	originalStat := osStat
	defer func() {
		_ = os.Setenv("XDG_RUNTIME_DIR", originalRuntimeDir)
		osStat = originalStat
	}()

	_ = os.Unsetenv("XDG_RUNTIME_DIR")
	osStat = func(name string) (os.FileInfo, error) {
		return nil, os.ErrNotExist
	}

	if actual := resolveDefaultDockerHost(); actual != defaultDockerHost {
		t.Fatalf("expected %s but got %s", defaultDockerHost, actual)
	}
}
