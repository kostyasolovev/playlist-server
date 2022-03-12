package models

import "fmt"

// методы, которые обязательно должны быть у ответа.
type Playlist interface {
	Error() string
	Print() []string
}

// форма ответа от удаленного сервера.
type Answer struct {
	Err   string
	Items []string // например названия видео файлов
}

// реализация интерфейса Playlists.
func (ans *Answer) Error() string { return ans.Err }

func (ans *Answer) Print() []string { return ans.Items }

func MockRPC(id string) Playlist {
	ans := &Answer{
		Items: make([]string, 10),
	}

	for i := 0; i < 10; i++ {
		ans.Items[i] = fmt.Sprintf("video %d", i)
	}

	return ans
}
