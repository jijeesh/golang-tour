package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}
