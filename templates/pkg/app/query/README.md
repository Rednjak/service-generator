## Guidance

The Query handlers are only reading from the system. It's highly likely that our read model will be different than our domain model, which is why it's commong to create a query model. This way we separate
the concern between the models and we can easily adjust the query model without affecting our domain model.

We don't always need to add a model for each query. It always depends on the use case. It's also fine to
simply use the domain model to query the data. In case we see that defining another model makes sense we can do it.
