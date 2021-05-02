## Guidance

The Command handlers are used for modifying the system state. It's not that common that we would need a separate application model. Usually we would use the domain model to change the system.

In the example handler we only use a repository, but we could also add another service that needs to be called or anything that we need to complete our use case.

Upon the project creation we should rename the resource to the proper action
