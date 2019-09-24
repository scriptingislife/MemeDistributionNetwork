build:
	GOOS=windows GOARCH=amd64 go build
	GOOS=linux   GOARCH=amd64 go build

serve:
	cd memes && python3 http.server