package domain

import "time"

type FilmData struct {
	Uuid         string
	Preview      string
	Title        string
	Director     string
	AverageScore float32
	ScoresCount  int
	Duration     int
	Date         time.Time
	Data         string
	AgeLimit     uint8
}

type FilmDataToAdd struct {
	Title       string
	Preview     string
	Director    string
	Data        string
	AgeLimit    uint8
	Duration    int
	PublishedAt time.Time
	Actors      []ActorData
}

type FilmPreview struct {
	Uuid         string  `json:"uuid"`
	Preview      string  `json:"preview_data"`
	Title        string  `json:"title"`
	Director     string  `json:"author"`
	AverageScore float32 `json:"average_score"`
	ScoresCount  int     `json:"scores_count"`
	Duration     int     `json:"duration"`
}

type FilmLink struct {
	Uuid  string `json:"uuid"`
	Title string `json:"title"`
}

func NewMockFilmData() FilmData {
	return FilmData{
		Uuid:     "1",
		Title:    "Fast n Furious",
		Preview:  "avatar",
		Director: "Danya",
		Data:     "information",
		AgeLimit: 0,
		Duration: 240,
	}
}

func NewMockFilmDataToAdd() FilmDataToAdd {
	return FilmDataToAdd{
		Title:    "Fast n Furious",
		Preview:  "avatar",
		Director: "Danya",
		Data:     "information",
		AgeLimit: 18,
		Duration: 240,
		Actors: []ActorData{
			{"1", "Fast n Furious 1", "avatar", time.Now(), "", 100,
				"", "", "", []FilmLink{{"1", "Fast n Furious 1"}}},
		},
	}
}

func NewMockFilmPreview() FilmPreview {
	return FilmPreview{
		Uuid:         "1",
		Preview:      "avatar",
		Title:        "Fast n Furious",
		Director:     "Danya",
		AverageScore: 0,
		ScoresCount:  10,
		Duration:     240,
	}
}

func NewMockFilmPreviews() []FilmPreview {
	return []FilmPreview{
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

func NewMockFilmActors() []ActorPreview {
	return []ActorPreview{
		{"1", "Fast n Furious 1", "avatar"},
		{"2", "Fast n Furious 2", "avatar"},
		{"3", "Fast n Furious 3", "avatar"},
	}
}

func NewMockFilmComments() []Comment {
	return []Comment{
		{"1", "1", "Fast n Furious 1", "comment1", 1, time.Now()},
		{"2", "1", "Fast n Furious 2", "comment2", 1, time.Now()},
		{"3", "1", "Fast n Furious 3", "comment3", 1, time.Now()},
	}
}
