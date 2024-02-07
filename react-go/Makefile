.PHONY: deps app server build docker-run

deps:
	# installing air for go hot reloading
	cd ./server && \
	go install github.com/cosmtrek/air@latest

app:
	cd ./app && \
	npm run dev 

server:
	cd ./server && \
	air -c .air.toml

build:
	@echo "building docker image" && \
	docker build -t awesome-templates-react-go . 

docker-run:
	docker run -p 3000:3000 -e ENVIRONMENT=production awesome-templates-react-go