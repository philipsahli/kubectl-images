package main

import (
	"fmt"
	"log"

	"github.com/philipsahli/client-go-wrapper/pkg/wrapper"
)

func main() {
	fmt.Println("kubectl-images: ")

	pods, err := wrapper.GetPods()
	if err != nil {
		log.Fatalf("Error in GetPods: %s", err)
	}

	for _, pod := range pods {
		for _, c := range pod.Spec.Containers {
			fmt.Println(c.Image)
		}
		for _, c := range pod.Spec.InitContainers {
			fmt.Println(c.Image)
		}

	}
}
