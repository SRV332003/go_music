package recom

import (
	"github.com/SRV332003/go_music/filemanager"
)

func PlayRandom() filemanager.Song {
	song := filemanager.GetRandom()
	return song
}
