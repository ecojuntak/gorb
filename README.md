## My Personal REST API Boilerplate

[![Build Status](https://travis-ci.org/ecojuntak/gorb.svg?branch=master)](https://travis-ci.org/ecojuntak/gorb)
[![Go Report Card](https://goreportcard.com/badge/github.com/ecojuntak/gorb)](https://goreportcard.com/report/github.com/ecojuntak/gorb)

Simple mini framework for creating REST API using GoLang.

## Installation
Clone the framework
```
git clone https://github.com/ecojuntak/gorb.git
```

## Configuration
Create your configuration file first
```
cp config.example.yaml config.yaml
```
### General Configuration
```
APP_NAME: GO-REST
PORT: 8000
HOST: localhost
```
For your database, you can use MySQL, Postgresql, and SQLite3.
### MySQL Configuration
```
DB_DRIVER: mysql
DB_HOST: 127.0.0.1
DB_PORT: 3306
DB_NAME: gorest
DB_USERNAME: username
DB_PASSWORD: secret
```
### Postgresql Configuration
```
DB_DRIVER: postgres
DB_HOST: 127.0.0.1
DB_PORT: 5342
DB_NAME: gorest
DB_USERNAME: username
DB_PASSWORD: secret
DB_POSTGRES_SSL_MODE: disable
```
### SQLite3 Configuration
For SQLite3, you need to create your database file first.
```
touch model/database.db
```
Then update your configuration file. You only need to set the DB_DRIVER and DB_NAME
```
DB_DRIVER: sqlite3
DB_NAME: model/database.db
```
## Migration and Seeder
To run the migration
```
go run *.go migrate
```
To run the seeder
```
go run *.go seed
```
## Start the server
To start the server
```
go run *.go start
```
## Contribution
Pull requests are welcome!

Cappy Hoding!
