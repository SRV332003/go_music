package recom

import (
	"dhvani/filemanager"
	"dhvani/player"
)

func PlayRandom() {
	song := filemanager.GetRandom()
	player.Play(song.Path)
}
