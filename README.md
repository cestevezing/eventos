# Eventos API REST

## Requerimientos

Asegúrate de tener instalados los siguientes requisitos antes de ejecutar la API:

- PostgreSQL instalado en tu máquina o configurado en Docker.
- Go instalado en tu sistema.

## Configuración del proyecto

1. Crea una base de datos en PostgreSQL para su aplicación.

    - Si está utilizando Docker, podría ejecutar la siguiente línea:
        ```
        docker run --name eventos -e POSTGRES_USER=cestevezing -e POSTGRES_PASSWORD=mysecret -e POSTGRES_DB=eventos -p 5432:5432 -d postgres
        ```
    - Si tiene instalado PostgresSQL solo tendría que crear la base de datos:
        ```
        createdb eventos
        ```
2. Configura las credenciales de acceso a la base de datos en el archivo de configuración `.env` de la aplicación.
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=cestevezing
    DB_PASSWORD=mysecret
    DB_NAME=eventos
    ```
3. Migración de la Base de Datos. Ejecuta las migraciones de Gorm para crear las tablas necesarias en la base de datos.
    ```
    go run cmd/migrations/main.go
    ```
4. Carga de Datos por Defecto. Ejecuta el siguiente comando para cargar datos por defecto en la base de datos.
    ```
    go run cmd/loaddata/main.go
    ```
5. Antes de ejecutar el proyecto asegúrese de ejecutar la siguiente línea, para que pueda consultar la documentación:
    ```
    swag init -g cmd/events/main.go
    ```


## Ejecución de la API

Para ejecutar la API, utiliza el siguiente comando:

```
go run cmd/events/main.go
```  

La API estará disponible en `http://localhost:3001`.

## Documentación en Swagger

Accede a la documentación de la API en Swagger en la siguiente URL:

http://localhost:3001/swagger/index.html

    