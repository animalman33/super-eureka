# super-eureka

Scaleable chat application - not for production use

## Database setup for local development

Use the docker commands following to start the DB automatically as part of the dev compose file

create a .env file and add the following variables

POSTGRES_PASSWORD=''
POSTGRES_USER=''
POSTGRES_DB=''

values can be whatever you want

## To run develop containers

`docker compose -f docker/docker-compose.dev.yaml up --watch`

## To stop develop containers

`docker compose -f docker/docker-compose.dev.yaml down --rmi local --remove-orphans`
 
If the database needs to be reset

`docker compose -f docker/docker-compose.dev.yaml down --rmi local --remove-orphans -v`