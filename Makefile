build-web:
	cd web; npm install; npm run build

build-linux:
	export PATH=$$PATH:/usr/local/bin
	env GOOS=linux GOARCH=amd64 go build ./cmd/webapp

run: build-linux build-web
	docker-compose up --build