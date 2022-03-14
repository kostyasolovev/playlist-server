package models

import "fmt"

// методы, которые обязательно должны быть у ответа.
type Playlist interface {
	GetErr() string
	GetItem() []string
}

// форма ответа от удаленного сервера.
type AnswerMock struct {
	Err   string
	Items []string // например названия видео файлов
}

// реализация интерфейса Playlists.
func (ans *AnswerMock) GetErr() string { return ans.Err }

func (ans *AnswerMock) GetItem() []string { return ans.Items }

func MockRPC(id string) Playlist {
	ans := &AnswerMock{
		Items: make([]string, 10),
	}

	for i := 0; i < 10; i++ {
		ans.Items[i] = fmt.Sprintf("video %d", i)
	}

	return ans
}
