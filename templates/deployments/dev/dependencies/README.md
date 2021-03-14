### Purpose

Usually we would have a separate dependency project where we would initialize everything that we need for our systems. This docker-compose should be used in case there is a service specific dependency that is needed.

Currently the file will build a database and execute all the migration files within the migrations folder. The initialization only happens when we first run the docker-compose up. In case we want to re-run all the migrations we first need to remove the volume. This is done with `docker-compose down -v`.
