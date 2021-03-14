Service generator's purpose is to create an initial structure of Golang services that are being used in a micro-service architecture.
The project structure is following the principles of the Clean architecture. The generated service should contain the logic for a
specific domain within the system, but it's also possible to use multiple domain within one service.

## Motivation

This project was build in my free time while studying Golang and reading about DDD and Clean architecture.
Since the use case of the created service can't be known up front I imagined that all the
resource names(methods, services, storage etc.) would be renamed after being created.

I would be happy to receive any feedback on my approach. All contributions and ideas are welcomed.

## Installation and usage

Execute `go get github.com/Rednjak/service-generator` to install the package locally. I you don't have the the GO bin folder in your path please do so by
adding `export PATH=$PATH:$HOME/go/bin` to your .zshrc, bashrc etc. In case you have a custom GOPATH please update the provided example accordingly.

After completing the above instructions you should be able to create a binary file by moving into the project folder `cd $HOME/go/src/github.com/Rednjak/service-generator` and executing `go install`.

Since you added the `bin` folder to your PATH you should be able to execute `service-generator` command. Follow the CLI instructions to create a project.

`service-generator my-project-name -m module-name`

### Considerations

In case you don't need any files you can simply remove them. Usually the deployments folder would not be needed since it's probably it's own project where you build all the dependency of your services. The same goes for the migrations folder.

### Examples

You can execute `./scripts/restart_server.sh` to start the docker container for the provided MySQL instance and the server itself.
Running `./scripts/restart_server.sh stop` will stop the database container upon exiting the server.

Keep in mind that currently there is no logic in the services and therefore you won't get any response from the server if you decide to test the API.
