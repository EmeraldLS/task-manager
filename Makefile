run:
	go run main.go

tidy:
	go run tidy

build:
	go run build

test:
	go test ./manager/tests -v

.PHONY:
	run