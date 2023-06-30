package docker

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type BackstopImageNotInstalled bool
type BackstopImageInstalled bool

func CheckForImage() tea.Msg {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// Image name to check
	imageName := "backstopjs/backstopjs:latest"

	// Check if the image exists
	imageExists, err := checkImageExists(cli, imageName)
	if err != nil {
		panic(err)
	}

	if !imageExists {
		return BackstopImageNotInstalled(true)
	}
	return BackstopImageInstalled(true)
}

func checkImageExists(cli *client.Client, imageName string) (bool, error) {
	// Retrieve a list of images on the machine
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return false, err
	}

	// Check if the specified image is present in the list
	for _, image := range images {
		for _, tag := range image.RepoTags {
			if tag == imageName {
				return true, nil
			}
		}
	}

	return false, nil
}
