## Description

The application logic is the glue between the ports and domain layer. In general it executes the use cases that we have defined. It shouldn't know who is calling it and what kind of a database(Adapter) is being used when executing some logic.

## Guidance

We split the use cases into Commands and Queries based on the CQRS principles(Command Query Responsibility Segregation). The basic principle is that Commands should makes changes in the system while Queries are only used to read the data from the system itself.

## When not to use CQRS

In case you application is saving and returning the same kind of data it doesn't make sense to over-engineer the Application layer. Keep it simple until you identify the need for introducing CQRS.

A good example is also a simple CRUD application. CQRS would most likely add too much complexity to a CRUD app. In case the CRUD app evolves it would make sense to re-evaluate if CQRS would make sense.
