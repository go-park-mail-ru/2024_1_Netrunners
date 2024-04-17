package database

import (
	"time"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

func NewMockActor() domain.ActorData {
	return domain.ActorData{
		Uuid:       "1",
		Name:       "Danya",
		Avatar:     "http://avatar",
		Birthday:   time.Now(),
		Career:     "career",
		Height:     192,
		BirthPlace: "Angarsk",
		Genres:     "Riddim",
		Spouse:     "Дабстеп",
		Films: []domain.FilmLink{
			{Uuid: "1", Title: "Fast n Furious 1"},
			{Uuid: "2", Title: "Fast n Furious 2"},
			{Uuid: "3", Title: "Fast n Furious 3"},
		},
	}
}

func NewMockActorPreview() []domain.ActorPreview {
	return []domain.ActorPreview{
		{Uuid: "1", Name: "Fast n Furious 1", Avatar: "avatar"},
		{Uuid: "2", Name: "Fast n Furious 2", Avatar: "avatar"},
		{Uuid: "3", Name: "Fast n Furious 3", Avatar: "avatar"},
	}
}

func NewMockFilmData() domain.FilmData {
	return domain.FilmData{
		Uuid:     "1",
		Title:    "Fast n Furious",
		Preview:  "avatar",
		Director: "Danya",
		Data:     "information",
		AgeLimit: 0,
		Duration: 240,
	}
}

func NewMockFilmDataToAdd() domain.FilmDataToAdd {
	return domain.FilmDataToAdd{
		Title:    "Fast n Furious",
		Preview:  "avatar",
		Director: "Danya",
		Data:     "information",
		AgeLimit: 18,
		Duration: 240,
		Actors: []domain.ActorData{
			{Uuid: "1", Name: "Fast n Furious 1", Avatar: "avatar", Birthday: time.Now(), Career: "", Height: 100,
				BirthPlace: "", Genres: "", Spouse: "", Films: []domain.FilmLink{
					{Uuid: "1", Title: "Fast n Furious 1"},
				},
			},
		},
	}
}

func NewMockFilmPreview() domain.FilmPreview {
	return domain.FilmPreview{
		Uuid:         "1",
		Preview:      "avatar",
		Title:        "Fast n Furious",
		Director:     "Danya",
		AverageScore: 0,
		ScoresCount:  10,
		Duration:     240,
	}
}

func NewMockFilmPreviews() []domain.FilmPreview {
	return []domain.FilmPreview{
		{
			Uuid:         "1",
			Preview:      "avatar",
			Title:        "Fast n Furious",
			Director:     "Danya",
			AverageScore: 0,
			ScoresCount:  10,
			Duration:     240,
		},
		{
			Uuid:         "2",
			Preview:      "avatar",
			Title:        "Fast n Furious 2",
			Director:     "Danya",
			AverageScore: 0,
			ScoresCount:  10,
			Duration:     120,
		},
	}
}

func NewMockFilmActors() []domain.ActorPreview {
	return []domain.ActorPreview{
		{Uuid: "1", Name: "Fast n Furious 1", Avatar: "avatar"},
		{Uuid: "2", Name: "Fast n Furious 2", Avatar: "avatar"},
		{Uuid: "3", Name: "Fast n Furious 3", Avatar: "avatar"},
	}
}

func NewMockFilmComments() []domain.Comment {
	return []domain.Comment{
		{Uuid: "1", Author: "Fast n Furious 1", Text: "comment1", Score: 1, AddedAt: time.Now()},
		{Uuid: "2", Author: "Fast n Furious 2", Text: "comment1", Score: 1, AddedAt: time.Now()},
		{Uuid: "3", Author: "Fast n Furious 3", Text: "comment1", Score: 1, AddedAt: time.Now()},
	}
}

func NewMockUser() domain.User {
	return domain.User{
		Uuid:         "1",
		Email:        "cakethefake@gmail.com",
		Avatar:       "",
		Name:         "Danya",
		Password:     "123456789",
		IsAdmin:      true,
		RegisteredAt: time.Now(),
		Birthday:     time.Now(),
	}
}

func NewMockUserSignUp() domain.UserSignUp {
	return domain.UserSignUp{
		Email:    "cakethefake@gmail.com",
		Name:     "Danya",
		Password: "123456789",
	}
}
