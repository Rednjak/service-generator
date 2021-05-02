## Description

Domain is the core of the application. It contains all of the domain objects and the business rules.
It also defines how we will communicate with databases in an agnostic way. This allows us to quickly switch between databases.

Domain itself is not aware of the outer layers!

Currently the domain is named in a generic way "resource" and should be renamed upon creating the service and identifying the domain models. Most likely there will be multiple domain models identified and for each of the there should be a file.

## Guidance

Each of the domain models should keep their fields private to ensure their state can't be modified in an unexpected way(Always keep a valid state in the memory!). Have a constructor to initialize a struct and verify that the correct values are provided.

The domain model should be database agnostic and should focus on the business rules.
