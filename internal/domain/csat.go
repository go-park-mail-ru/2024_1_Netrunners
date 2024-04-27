package domain

type CheckQuestionStatistics struct {
	Title string
	Count int32
}

type QuestionStatistics struct {
	Type          string // decimal / check
	Title         string
	ScoresCount   uint32
	AverageScore  float32
	CheckVariants []CheckQuestionStatistics
}

type Question struct {
	Uuid      string
	Title     string
	Type      string // decimal / check
	CheckVars []string
}

type AddQuestionStatistics struct {
	Uuid         string
	Type         string // decimal / check
	Score        int32
	CheckVariant int32
}

type AddQuestion struct {
	Page          string
	Title         string
	Type          string // decimal / check
	CheckVariants []string
}
