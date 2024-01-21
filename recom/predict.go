package recom

import (
	"go_music/filemanager"
)

func PlayRandom() filemanager.Song {
	song := filemanager.GetRandom()
	return song
}
