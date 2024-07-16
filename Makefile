build:
	GOOS=linux GOARCH=amd64 go build -o builds/go_music .

place:
	sudo cp ./builds/go_music /usr/local/bin/dhvani

deploy-linux: build place

deploy-windows:
	GOOS=windows GOARCH=amd64 go build -o builds/go_music.exe .
	copy ./builds/go_music.exe /usr/local/bin/dhvani

run:
	go run .