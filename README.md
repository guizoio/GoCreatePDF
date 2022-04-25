# GoCreatePDF
Project create file PDF in GO
## :books: Project description
<!--ts-->
* Main objective study
* Create PDF file
  * Endpoint to download PDF file
* Save PDF file information to database
<!--te-->

## :wrench: Tecnologias utilizadas
* Goland
* Postgres
* Docker-Compose

## :rocket: Running the project
### `First step` create .env file.
```sh
make env
```

### `Second step` run docker-compose.
```sh
make up
```
to stop docker-compose run the command:`make down`.

### `Third step` run the command to create the database tables.
```sh
make migrations
```

### `Fourth step` to create test PDF file.
```sh
make run
```
PDF file is generated in project root with fake data.

### `Fifth step` run the HTTP server.
```sh
make server
```