package models

// Структура для одного игрока (то, что будем отдавать в JSON)
type Player struct {
	ID       int     `json:"id"`
	FullName string  `json:"full_name"`
	City     string  `json:"city"`
	Height   int     `json:"height_cm"`
	Weight   float64 `json:"weight_kg"`
}
