server:
	go run src/main.go

migration-create:
	go run src/main.go migration create

migration-update:
	go run src/main.go migration update
	
docker-compose:
	docker-compose up -d