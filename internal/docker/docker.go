package docker

import (
	"context"
	"io"
	"log"
	"os"
	"strings"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

type DockerService struct {
	cli *client.Client
}

func NewDockerService() (*DockerService, error) {
	opts := []client.Opt{
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	}

	if os.Getenv("DOCKER_HOST") == "" {
		home, _ := os.UserHomeDir()
		orbstackSocket := home + "/.orbstack/run/docker.sock"
		if _, err := os.Stat(orbstackSocket); err == nil {
			log.Printf("OrbStack detected, using socket: %s", orbstackSocket)
			opts = append(opts, client.WithHost("unix://"+orbstackSocket))
		}
	}

	cli, err := client.NewClientWithOpts(opts...)
	if err != nil {
		log.Printf("Failed to create Docker client: %v", err)
		return nil, err
	}
	return &DockerService{cli: cli}, nil
}

func (s *DockerService) GetContainers(ctx context.Context) ([]container.Summary, error) {
	result, err := s.cli.ContainerList(ctx, client.ContainerListOptions{All: true})
	if err != nil {
		log.Printf("Docker ContainerList error: %v", err)
		return nil, err
	}
	return result.Items, nil
}

func (s *DockerService) StartContainer(ctx context.Context, id string) error {
	_, err := s.cli.ContainerStart(ctx, id, client.ContainerStartOptions{})
	return err
}

func (s *DockerService) StopContainer(ctx context.Context, id string) error {
	_, err := s.cli.ContainerStop(ctx, id, client.ContainerStopOptions{})
	return err
}

func (s *DockerService) RestartContainer(ctx context.Context, id string) error {
	_, err := s.cli.ContainerRestart(ctx, id, client.ContainerRestartOptions{})
	return err
}

func (s *DockerService) GetContainerLogs(ctx context.Context, id string) (string, error) {
	options := client.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Tail:       "100",
	}

	out, err := s.cli.ContainerLogs(ctx, id, options)
	if err != nil {
		return "", err
	}
	defer out.Close()

	var b strings.Builder
	_, err = io.Copy(&b, out)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (s *DockerService) Close() error {
	return s.cli.Close()
}
