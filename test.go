package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	dockerBuild := exec.Command("docker", "build", "-t", "my-image", ".")
	dockerBuild.Stderr = os.Stderr

	_, err := dockerBuild.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	// otherwise, print the output from running the command
	cmd := exec.CommandContext(ctx, "docker", "run", "my-image")
	cmd.Stderr = os.Stderr

	out1, err1 := cmd.Output()
	if err1 != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err1)
	}
	fmt.Println("Output: ", string(out1))

}
