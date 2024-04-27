package domain

type CheckQuestionStatistics struct {
	Title string
	Count int32
}

type QuestionStatistics struct {
	Title         string
	IsAdditional  bool
	ScoresCount   uint32
	AverageScore  float32
	CheckVariants []CheckQuestionStatistics
}

type Variant struct {
	Id    uint32
	Title string
}

type AdditionalQuestion struct {
	Uuid      string
	Title     string
	CheckVars []Variant
}

type Question struct {
	Uuid               string
	Title              string
	AdditionalQuestion AdditionalQuestion
}

type AddQuestionStatistics struct {
	Uuid         string
	IsAdditional bool
	Score        int32
	// CheckVariant int32
}

type AddQuestion struct {
	Page          string
	Title         string
	CheckVariants []string
}
