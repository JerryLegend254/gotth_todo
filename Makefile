build:
	@tailwindcss -i views/css/styles.css -o public/styles.css
	@templ generate view
	@go build -o bin/gotth main.go

test:
	@go test -v ./...

run: build
	@./bin/gotth

tailwind:
	@tailwindcss -i views/css/styles.css -o public/styles.css --watch

templ:
	@templ generate -watch -proxy=http://localhost:8080

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
