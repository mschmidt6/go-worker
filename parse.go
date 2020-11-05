package main

import (
	"fmt"

	"github.com/containers/image/v5/docker"
	"github.com/containers/image/v5/docker/reference"
	"github.com/containers/image/v5/types"
)

func parseImage(name string) (types.ImageReference, error) {
	imgRef, err := docker.ParseReference(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(imgRef.DockerReference().(reference.NamedTagged).Tag())
	fmt.Println(imgRef.Transport().Name())
	fmt.Println(imgRef.DockerReference().Name())
	fmt.Println(imgRef.DockerReference().String())
	return imgRef, nil
}
