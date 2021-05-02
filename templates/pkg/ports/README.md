## Description

Ports are the entry points to the Application. They could be http/grpc server, cli or any other way in which you would want to to talk to your Application. They can access the inner layers and execute the Application use cases.

## Guidance

Ports usually contain an http server with handlers that are transforming incoming models to any application models, calling the Application use cases and responding back to the clients.

In the package we defined the http handlers, routes and the payload types. We can additionally extend the handlers and routes with specific middleware important for use case execution.

In case we use OpenAPI we could put here the generated types and handlers
