package domain

type FilmPreview struct {
	Id string `json:"uuid"`
	// Preview      []byte `json:"preview_data"`
	Preview      string `json:"preview_data"`
	Name         string `json:"name"`
	Author       string `json:"author"`
	AverageScore int    `json:"average_score"`
	ScoresCount  int    `json:"scores_count"`
	Duration     int    `json:"duration"`
}
