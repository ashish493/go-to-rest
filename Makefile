build:
	go build -o bin/main main.go


run:
	go run main.go

test:
	go test tests/ -v

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 main.go

	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 main.go


release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

	