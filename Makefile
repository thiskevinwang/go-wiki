.PHONY dev:
dev: # run webserver
	@go run wiki.go

.PHONY up:
up: # build and start webserver
	@go build wiki.go
	./wiki