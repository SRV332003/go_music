package filemanager

import (
	"fmt"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
)

func Search(s string) []Song {

	mp := make(map[Song]int)

	for _, song := range files {

		n := strings.Count(strings.ToLower(song.Name), strings.ToLower(s))

		// fmt.Println(n, song.id, song.name, s)
		mp[song] = n
	}

	f := make([]Song, len(files))
	copy(f, files)

	for i := 0; i < 10; i++ {
		for j := range f {
			if j == 0 {
				continue
			}
			if mp[f[j]] < mp[f[j-1]] {
				temp := f[j]
				f[j] = f[j-1]
				f[j-1] = temp
			}
		}
	}

	// for j := range f {
	// 	fmt.Print(mp[f[j]], " ")
	// }

	var ans []Song
	for i := len(f) - 10; i < len(f); i++ {
		if mp[f[i]] > 0 {
			ans = append(ans, f[i])
		}

	}

	return ans
}

func AdvSearch() (Song, error) {

	idx, err := fuzzyfinder.Find(files, func(i int) string {
		return fmt.Sprintf("%s", files[i].Name)
	})

	if err != nil {
		return Song{}, err
	}

	return files[idx], nil

}
