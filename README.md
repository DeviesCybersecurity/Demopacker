# Demopacker

Demopacker is intended to be a template for packaging different files into a Docker image to be run as part of trainings, demos, etc.

The purpose of using Docker is that this can be run in any environment or operating system which has Docker installed, making it an excellent choice in performing labs, assignments, and demos without having to install any additional files.

The intention is that you create your package, select which environment it should be run in (See [here](#environments)) and any additional tools you need. You then build your Docker image, which should then be immediately ready to be run as part of your demo!

**THIS IS A TEMPLATE, DO NOT ADD CUSTOM PACKAGES HERE**

# Instructions
1. Use this Git repository as a Template for your own project.
2. Modify `config.yml` to fit your specific needs.
3. Add your specific demo files to the 'Package' folder'
4. Ensure Golang version 1.19 or later is installed.
5. Run the project using `go run packager.go`
6. You can now push the docker image, or inspect it using `docker run -it <image name>`

# Environments
The environments are specific Docker base images which you can choose to build your image on top of.

## C++
The C++ (`cpp`) environment contains tools for compiling and debugging C++ code. It contains:
- gcc
- clang
- clang-tools
- cmake
- gdb

## Golang
The Golang (`golang`)

# Potential Improvements
- More environments, ideas:
    - Web servers
    - Python
    - ...?
- Add parameters for things such as docker tagging, etc.
- Create a packaging step which performs additional commands such as building and copying binaries.
