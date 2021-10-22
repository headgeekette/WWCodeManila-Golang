# SQL and databases with Go

## Session schedules

| Date                                                     |
| -------------------------------------------------------- | 
| Saturday, October 23, 2021, 2:00 pm - 4:00 pm GMT+08:00  |

## Build the mariadb Dockerfile

`$ docker build -t go-mariadb .`

## Start `mariadb` server instance

`$ docker run -p 127.0.0.1:3306:3306  --name some-mariadb -e MARIADB_ROOT_PASSWORD=my-secret-pw -d go-mariadb:latest`

or

```
$ docker network create some-network # Create the network
$ docker run --net some-network --name some-mariadb -e MARIADB_ROOT_PASSWORD=my-secret-pw -d mariadb:tag
```

