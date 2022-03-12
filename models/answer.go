package models

import "fmt"

// методы, которые обязательно должны быть у ответа.
type Playlist interface {
	Error() string
	Print() []string
}

// форма ответа от удаленного сервера.
type Answer struct {
	Err   error
	Items []string // например названия видео файлов
}

// реализация интерфейса Playlists.
func (ans *Answer) Error() string { return ans.Err.Error() }

func (ans *Answer) Print() []string {
	buf := make([]string, 0, len(ans.Items))
	for _, v := range ans.Items {
		buf = append(buf, v)
	}

	return buf
}

func MockRPC(id string) Playlist {
	ans := &Answer{
		Items: make([]string, 10),
	}

	for i := 0; i < 10; i++ {
		ans.Items[i] = fmt.Sprintf("video %d", i)
	}

	return ans
}
