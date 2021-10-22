# The Sakila sample database

This session is using the Sakila database provided by MySQL. You can find the full documentation from the MySQL Sakila website.

<https://dev.mysql.com/doc/sakila/en/sakila-structure.html>

You can use either a MySQL or MariaDB system for this. The scripts to create the database can be found in `sakila-db/`. Run the schema script first, then the data script.

### The database Docker container

Using the Docker container is optional. If you already have MySQL or MariaDB locally installed, you can use that instead.

The container uses MariaDB. MariaDB works the same as MySQL.

To build the container:

`docker build -t go-mariadb .`

To start the container:

`$ docker run -p 127.0.0.1:3306:3306  --name <alias-for-container> -e MARIADB_ROOT_PASSWORD=<password> -d go-mariadb:latest`

Account credentials can be found in the `Dockerfile`.
