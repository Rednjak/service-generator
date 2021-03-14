### Idea

The service is designed with the clean architecture in mind.

- We have the app layer which handler the incoming requests, their validation, calling the service layer, serializing the response and error handling.
- The second layer is the service layer where we do our business logic. This layer also communicates with the repository layer.
- The third layer is the repository layer which holds the models and takes care of the communication with the database.

The repository layer doesn't know anything about the service layer and the service layer doesn't know anything about the app
layer. The direction of communication is going from outwards to inwards which ensures lose coupling.

On initial service generation the resources and methods will be named in a generic way. All of them need to be adjusted per
the domain in which this service will work e.g. User.
