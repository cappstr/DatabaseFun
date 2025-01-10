# DatabseFun

project outlining good practises when connecting to a database in Go

## Tools
Docker (docker-compose)
Postgres
go migrate - https://github.com/golang-migrate/migrate
  - reads migrations from a source and applies them to a database in the correct order
  - see makefile for a couple of commands related for this
envrc
  - create environment variables for each project without cluttering your terminal rc file
  - direnv allow
