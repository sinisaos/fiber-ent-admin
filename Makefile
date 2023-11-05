server:
	go run main.go

generate:
	go generate ./ent

tests:
	go test -v -cover ./service/...
	