.PHONY: dependency install

dependency: 
	go get github.com/tools/godep

install:
	godep restore
	go install

runlocal:
	ENVIRONMENT=local DB_HOST=172.17.11.101 DB_PORT=5432 DB_NAME=toyo DB_USER=takasing DB_PASS=oifwslu4SGcBiMzL go run main.go
