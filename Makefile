run:
	go run *.go

main:
	go run server.go

test:
	go run server_test.go

gen:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

server:
	go run server.go

serverLocal:
	go run serverLocal.go

backup:
	cd graph/model && ls && cp models_gen.go backup

dockerUp:
	docker run --name postgres -e POSTGRES_PASSWORD=Eauu0244 -p 5432:5432 -d postgres


dockerClear:
	docker kill $$(docker ps -q)
	docker rm $$(docker ps -a -q)

deploy:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main
	zip deployment.zip main

deploywindows:
	GOOS=windows go build -o main
	gzip main -k -f -9
	mv main.gz main.zip

client:
	genqlient graph/genqlient.yaml

t:
	go test -v ./test/...