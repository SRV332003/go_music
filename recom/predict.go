package recom

import (
	"dhvani/filemanager"
)

func PlayRandom() filemanager.Song {
	song := filemanager.GetRandom()
	return song
}
