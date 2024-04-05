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
			{"1", "Fast n Furious 1"},
			{"2", "Fast n Furious 2"},
			{"3", "Fast n Furious 3"},
		},
	}
}

func NewMockActorPreview() []domain.ActorPreview {
	return []domain.ActorPreview{
		{"1", "Fast n Furious 1", "avatar"},
		{"2", "Fast n Furious 2", "avatar"},
		{"3", "Fast n Furious 3", "avatar"},
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
			{"1", "Fast n Furious 1", "avatar", time.Now(), "", 100,
				"", "", "", []domain.FilmLink{{"1", "Fast n Furious 1"}}},
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
		{"1", "Fast n Furious 1", "avatar"},
		{"2", "Fast n Furious 2", "avatar"},
		{"3", "Fast n Furious 3", "avatar"},
	}
}

func NewMockFilmComments() []domain.Comment {
	return []domain.Comment{
		{"1", "1", "Fast n Furious 1", "comment1", 1, time.Now()},
		{"2", "1", "Fast n Furious 2", "comment2", 1, time.Now()},
		{"3", "1", "Fast n Furious 3", "comment3", 1, time.Now()},
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
