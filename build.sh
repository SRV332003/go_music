#GOOS=windows GOARCH=amd64 go build -o builds/go_music.exe .

GOOS=linux GOARCH=amd64 go build -o builds/go_music .
sudo cp ./builds/go_music /usr/local/bin/dhvani