up:
	docker-compose up -d

down:
	docker-compose down

env:
	cp ./.env.example ./.env

migrations:
	go run main.go MakeMigrations

run:
	go run main.go createPdf

server:
	go run main.go httpserver

