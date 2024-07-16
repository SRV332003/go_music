# Usage Guide for Go Music App

Welcome to the Go Music App! This guide will help you navigate through the app using keyboard shortcuts and commands. Get ready to control your music playback with ease.

## Keyboard Shortcuts
This app provides single-key shortcuts to control the music playback and navigate through the app. Here are the keyboard shortcuts you can use:

### Playback Control
- **`Space`**: Toggle pause/play for the current track.
- **`Right Arrow`**: Skip forward 10 seconds in the current track.
- **`Left Arrow`**: Skip backward 10 seconds in the current track.
- **`Up Arrow`**: Increase the volume.
- **`Down Arrow`**: Decrease the volume.
- **`N`**: Play the next track.
- **`L`**: Toggle loop for the current track.

### App Navigation
- **`C`**: Clear the screen.
- **`P`**: Toggle pause/play for the current track (same as Space).
- **`S`**: Advanced search. Allows you to search locally downloaded song interactively and play them.
- **`R`**: Play a random song.
- **`Q`**: Quit the app.
- **`H`**: Display help information about the app and its features.

### Volume Control
- **`+`**: Increase the volume.
- **`-`**: Decrease the volume.

## Command Inputs
Apart from the keyboard shortcuts, you can also use written commands to control the app. 
You can enter commands by typing `:` followed by the command. The app name will become red when you are in command mode.

Here are the commands you can use:

### General Commands
- **`ls`**: List all the downloaded songs in default music directory along with their `song_id`s.
- **`play <song_id>`**: Play the song with the given number from the list of downloaded songs seen after running `ls` command.
- **`exit`**: Exit command mode. You can also press `Enter`or `Return` to exit command mode.
- **`skip <duration-sec>`**: Skip by the given duration in seconds. Accepts both positive and negative values.

### YouTube Download Commands
- **`: <query>`**: Search for a song on YouTube, download it, and play it.
- **`~ <video-link>`**: Download a song from the given YouTube video link and play it.

### Folder Navigation
- **`cd <absolute-folder-path>`**: Navigate to the folder where the downloaded songs are stored. 

    This will remain the same for all the sessions. 
    By default, the folder is set to the `music` folder in the user home directory.

# 
That's all for now. 
More features incoming soon! Enjoy your music! 🎶
```
Got an idea for a feature? Feel free to contribute to the project on GitHub.
```