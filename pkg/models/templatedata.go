package models

type TemplateData struct {
	FloatMap       map[string]float32
	StringMap      map[string]string
	IntMap         map[int]int
	Data           map[string]interface{}
	CSRFToken      string
	FlashMessage   string
	WarningMessage string
	ErrorMessage   string
}
