# Go Music

This project targets building of a commandline music player that is able to download songs from youtube by providing a search interface in CLI itself. It can be used by various programmers while coding in their IDE. Folks will be able to play music in the terminal of IDE.

## Dependencies
The project relies on the following dependencies:
- [`ffmpeg`](https://ffmpeg.org/download.html): This is required to download the songs from youtube. You can download it from the link provided.

Above dependencies are required to be installed in the system to run the project.

## Install and Use
### Using `go` command
You need to have Go installed in your system to run this project. If you don't have Go installed, you can download it from [here](https://golang.org/dl/).
To install the project using `go`, you can use the following command in your terminal:
```bash
go install github.com/SRV332003/go_music@latest
```
Then to run the project, you can use the following command from any directory in your terminal:
```bash
go_music
```
### Using pre-built binary releases
You can download the pre-built binary releases from the [releases](https://github.com/SRV332003/go_music/releases/tag/v1.0.5) page. After downloading the binary, you can run it from your terminal using the following command:
```bash
# For linux
./go_music
```
OR
```bash
# For windows
./go_music.exe
```
Provided that you are in the same directory as the binary file.
Also, you can add the binary to your PATH to run it from any directory in your terminal.

PS: The binary releases are only available for linux and windows. If you are using MacOS, you can use the `go` command to install the project.

Note: I couldn't test the binary releases on MacOS, Windows and Linux except **`ubuntu`**. If you face any issues, reach out to me by creating an [issue...](https://github.com/SRV332003/go_music/issues). I will get back to you ASAP.

## Usage Guide for Go Music App
Go to the [Usage Guide](https://github.com/SRV332003/go_music/blob/main/Usage.md) to get started with the Go Music App. This guide will help you navigate through the app using keyboard shortcuts and commands. Get ready to control your music playback with ease.

## Project Setup

Golang version 1.21.5 is required to setup the project.([Download Go](https://go.dev/dl/))  

Apart from Go, the project requires `libasound2-dev` installed in the environment if you are using **linux**. Using the following command:
```bash
sudo apt install pkg-config
sudo apt-get install libasound2-dev  
```
If you face some error about dependencies in windows, kindly create an [issue...](https://github.com/SRV332003/go_music/issues)  



Make sure to install go modules:
```bash
go mod download
```

To run the project, you can use the following command:
```bash
go run main.go
```
OR
```bash
make run
```
Feel free to fork and create cool features. I will be happy to merge it.

## Contribute

If you are interested in contributing to it, you can created issues, setup this project in your system and create pull requests as well (mentioning proper explaination of changes).







