FROM mariadb:10.6.4

# ENV MY_NAME="John Doe"
ENV MARIADB_ROOT_PASSWORD="rootpasswd"
ENV MARIADB_DATABASE="sakila"
ENV MARIADB_USER="gouser"
ENV MARIADB_PASSWORD="gopasswd"

COPY ./sakila-db/01-sakila-schema.sql /docker-entrypoint-initdb.d/
COPY ./sakila-db/02-sakila-data.sql /docker-entrypoint-initdb.d/

