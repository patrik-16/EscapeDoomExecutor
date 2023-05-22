package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {

	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m

}

func setupForExecution(input *Request) string {

	//make directory
	if _, err := os.Stat(input.PlayerSessionId); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(input.PlayerSessionId, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	//make Language Docker File
	//TODO make it customizable for all case
	copy("java.Dockerfile", input.PlayerSessionId+"/"+".Dockerfile")

	filename := input.PlayerSessionId + "/" + "app.java"

	var _, err = os.Stat(filename)
	//make code File
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Fprintf(file, input.Code)
		defer file.Close()
	} else {
		fmt.Println("File already exists!", filename)
		return ""
	}

	//call Docker
	curr := executeDocker(input.PlayerSessionId+"/.Dockerfile", input.PlayerSessionId)

	err23 := os.RemoveAll(input.PlayerSessionId + "/")
	if err23 != nil {
		fmt.Println("lel not deleted")
		return ""
	}

	return curr
}

func executeDocker(dockerFileName string, name string) string {

	fmt.Println("Welcome")
	dockerBuild := exec.Command("docker", "build", "-t", name, "-f", dockerFileName, ".")
	dockerBuild.Stderr = os.Stderr

	_, err := dockerBuild.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	// otherwise, print the output from running the command
	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", name)
	cmd.Stderr = os.Stderr

	out1, err1 := cmd.Output()
	if err1 != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err1)
	}
	fmt.Println("Output: ", string(out1))

	dockerDeleteImage := exec.Command("docker", "rmi", name)
	dockerDeleteImage.Stderr = os.Stderr

	_, err234 := dockerDeleteImage.Output()
	if err234 != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}

	return string(out1)
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
