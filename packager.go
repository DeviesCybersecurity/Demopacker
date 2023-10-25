package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PackageName   string   `yaml:"package_name"`
	Environment   string   `yaml:"environment"`
	ExtraTools    []string `yaml:"extra_tools"`
	ExtraCommands []string `yaml:"extra_commands"`
}

func generateDockerfile(config Config) {
	var dockerfileContent strings.Builder

	dockerfileContent.WriteString(fmt.Sprintf("FROM demopacker-base-image:%s\n\n", config.Environment))

	// apt install the extra tools configured
	if len(config.ExtraTools) > 0 {
		dockerfileContent.WriteString("RUN apt-get update && apt-get install -y \\\n\t")
		dockerfileContent.WriteString(strings.Join(config.ExtraTools, " \\\n\t"))
		dockerfileContent.WriteString("\n\n")
	}

	// run extra commands configured
	for _, cmd := range config.ExtraCommands {
		dockerfileContent.WriteString(fmt.Sprintf("RUN %s\n", cmd))
	}
	dockerfileContent.WriteString("\n")

	// copy the 'Package' folder into /app
	dockerfileContent.WriteString("COPY Package/ /app/\n\n")

	// set the working directory to /app
	dockerfileContent.WriteString("WORKDIR /app\n")

	ioutil.WriteFile("Dockerfile", []byte(dockerfileContent.String()), 0644)
}

func main() {
	// Read the YAML file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	// Unmarshal the YAML data into the Config struct
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	// Build the base image for the environment
	cmd := exec.Command(
		"docker", "build",
		"-f", fmt.Sprintf("./Environments/Dockerfile.%s", config.Environment),
		"-t", fmt.Sprintf("demopacker-base-image:%s", config.Environment),
		".")
	// This is needed to see the output from the docker build command.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	generateDockerfile(config)

	// Build the final docker image for release
	cmd = exec.Command(
		"docker", "build",
		"-t", config.PackageName,
		".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
