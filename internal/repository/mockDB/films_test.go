package mockdb

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestAddFilm(t *testing.T) {
	validCases := []struct {
		testName string
		film     domain.FilmPreview
	}{
		{
			"add new film",
			domain.FilmPreview{
				Uuid:     "dfgea4ra424r4fw",
				Title:    "Film1",
				Duration: 3600,
			},
		},
		{
			"add new film",
			domain.FilmPreview{
				Uuid:     "fnuf7842huirn23",
				Title:    "Film2",
				Duration: 7200,
			},
		},
	}
	invalidCases := []struct {
		testName string
		film     domain.FilmPreview
	}{
		{
			"add existed film",
			domain.FilmPreview{
				Uuid:     "dfgea4ra424r4fw",
				Title:    "Film1",
				Duration: 3600,
			},
		},
		{
			"add existed film",
			domain.FilmPreview{
				Uuid:  "dfgea4ra424r4fw",
				Title: "Film3",
			},
		},
		{
			"add existed film",
			domain.FilmPreview{
				Uuid:     "fnuf7842huirn23",
				Title:    "Film2",
				Duration: 7200,
			},
		},
	}

	db := InitFilmsMockDB()

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.AddFilm(currentCase.film)
			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.AddFilm(currentCase.film)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestRemoveFilm(t *testing.T) {
	data := []domain.FilmPreview{
		{
			Uuid:     "dfgea4ra424r4fw",
			Title:    "Film1",
			Duration: 3600,
		},
		{
			Uuid:     "fnuf7842huirn23",
			Title:    "Film2",
			Duration: 7200,
		},
	}

	validCases := []struct {
		testName string
		id       string
	}{
		{
			"remove existed film",
			"dfgea4ra424r4fw",
		},
		{
			"remove existed film",
			"fnuf7842huirn23",
		},
	}
	invalidCases := []struct {
		testName string
		id       string
	}{
		{
			"remove unexisted film",
			"asfgerg",
		},
		{
			"remove unexisted film",
			"34tqgerf",
		},
	}

	db := InitFilmsMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Uuid] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.RemoveFilm(currentCase.id)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.RemoveFilm(currentCase.id)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestGetFilmPreview(t *testing.T) {
	data := []domain.FilmPreview{
		{
			Uuid:     "dfgea4ra424r4fw",
			Title:    "Film1",
			Duration: 3600,
		},
		{
			Uuid:     "fnuf7842huirn23",
			Title:    "Film2",
			Duration: 7200,
		},
	}

	validCases := []struct {
		testName string
		expected domain.FilmPreview
	}{
		{
			"get existed film",
			domain.FilmPreview{
				Uuid:     "dfgea4ra424r4fw",
				Title:    "Film1",
				Duration: 3600,
			},
		},
		{
			"get existed film",
			domain.FilmPreview{
				Uuid:     "fnuf7842huirn23",
				Title:    "Film2",
				Duration: 7200,
			},
		},
	}
	invalidCases := []struct {
		testName string
		expected domain.FilmPreview
	}{
		{
			"get unexisted film",
			domain.FilmPreview{
				Uuid:     "dfgea4ra424r4fw",
				Title:    "Film1",
				Duration: 7200,
			},
		},
		{
			"get unexisted film",
			domain.FilmPreview{
				Uuid:     "fnuf7842huirn23",
				Title:    "Film3",
				Duration: 7200,
			},
		},
	}

	db := InitFilmsMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Uuid] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			filmPreview, err := db.GetFilmPreview(currentCase.expected.Uuid)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, true, reflect.DeepEqual(filmPreview, currentCase.expected))
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			filmPreview, err := db.GetFilmPreview(currentCase.expected.Uuid)

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, false, reflect.DeepEqual(filmPreview, currentCase.expected))
		})
	}
}

func TestGetAllFilmsPreviews(t *testing.T) {
	data := []domain.FilmPreview{
		{
			Uuid:     "dfgea4ra424r4fw",
			Title:    "Film1",
			Duration: 3600,
		},
		{
			Uuid:     "fnuf7842huirn23",
			Title:    "Film2",
			Duration: 7200,
		},
	}

	db := InitFilmsMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Uuid] = currentData
		db.mutex.Unlock()
	}

	films := db.GetAllFilmsPreviews()

	assert.Equal(t, true, reflect.DeepEqual(films, data))
}
