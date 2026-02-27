package docker

import (
	"github.com/moby/moby/client"
)

func GetDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv)
}
