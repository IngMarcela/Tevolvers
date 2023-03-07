# CRUD API-Rest en Golang
Este proyecto es una API-Rest implementada en Golang que permite realizar operaciones CRUD (Create, Read, Update, Delete) en una entidad de datos con los siguientes campos: id (autogenerado), name, lastname, datetime, address, gender y age.

## Tecnologías utilizadas

- Golang
- MySQL (motor de base de datos)
- GORM (ORM para interactuar con la base de datos)
- OpenAPI 3 Swagger (documentación autogenerada)
- Docker (para el despliegue)


## Funcionalidades
La API cuenta con los siguientes endpoints:

- GET /users: Devuelve una lista paginada de usuarios con los campos id, name, lastname, datetime, address, gender y age.
- POST /users: Crea un nuevo usuario en la base de datos con los campos name, lastname, datetime, address, gender y age.
- GET /users/{id}: Devuelve la información de un usuario específico según su id.
- PUT /users/{id}: Actualiza la información de un usuario específico según su id.
- DELETE /users/{id}: Elimina un usuario específico según su id.

## Documentación
La documentación de la API se encuentra disponible en formato OpenAPI 3 Swagger y puede ser consultada mediante el endpoint GET /swagger.

## Despliegue
La API está disponible en un contenedor Docker que puede ser desplegado en cualquier servidor compatible con Docker. Para desplegar la API se deben seguir los siguientes pasos:

1. Clonar este repositorio.

2. Ejecutar el siguiente comando para crear la imagen de Docker:
```
docker-compose up -d 
docker run -p 8080:8080 --env-file .env crud-api-golang
```
3. O la otra opcion es ejecutar el ambiente local
```
go build -o main . 
./main
```