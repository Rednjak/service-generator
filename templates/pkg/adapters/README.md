## Description

Adapters are different ways in which the Application is communicating with the external world. This could be databases, other http clients, pub/sub etc.

## Guidance

Adapters can import inner layers. Usually, they will operate on the types found in the Application and Domain e.g. we retrieve an object from the database and convert it to a domain/application model.

Currently the folder contains a file which indicates on which resource it operates and what the underlying adapter is. In our case it's currently MySql. These files implement the domain repository interfaces.
