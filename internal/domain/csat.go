package domain

type CheckQuestionStatistics struct {
	Title string `json:"title"`
	Count int32  `json:"count"`
}

type QuestionStatistics struct {
	Title         string                    `json:"title"`
	IsAdditional  bool                      `json:"isAdditional"`
	ScoresCount   uint32                    `json:"scoresCount"`
	AverageScore  float32                   `json:"averageScore"`
	CheckVariants []CheckQuestionStatistics `json:"checkVariants"`
}

type Variant struct {
	Id    uint32 `json:"id"`
	Title string `json:"title"`
}

type AdditionalQuestion struct {
	Uuid      string    `json:"uuid"`
	Title     string    `json:"title"`
	CheckVars []Variant `json:"checkVars"`
}

type Question struct {
	Uuid               string             `json:"uuid"`
	Title              string             `json:"title"`
	AdditionalQuestion AdditionalQuestion `json:"additionalQuestion"`
}

type AddQuestionStatistics struct {
	Uuid         string `json:"uuid"`
	IsAdditional bool   `json:"isAdditional"`
	Score        int32  `json:"score"`
}

type AddQuestion struct {
	Page          string   `json:"page"`
	Title         string   `json:"title"`
	CheckVariants []string `json:"checkVariants"`
}
