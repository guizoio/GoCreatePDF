# GoCreatePDF
Project create file PDF in GO
## :books: Project description
<!--ts-->
* Main objective study
* Create PDF file
  * Endpoint to create PDF file
  * Endpoint to download PDF file
  * Save file in storage local [Minion](https://min.io/)
* Save PDF file information to database
<!--te-->

## :wrench: Tecnologias utilizadas
* Goland
* Postgres
* Minio
* Docker-Compose

## :rocket: Running the project
### `First step` run docker-compose.
```sh
make up
```
to stop docker-compose run the command:`make down`.
## > Soon after we will configure the local storage, [wiki](https://github.com/guizoio/GoCreatePDF/wiki) to help.


### `Second step` create .env file.
```sh
make env
```

### `Third step` run the command to create the database tables.
```sh
make migrations
```

### `Fourth step` run the HTTP server.
```sh
make server
```