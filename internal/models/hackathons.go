package models

type Hackathon struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Money          int    `json:"money" db:"money"`
	ParticipantMin *int   `json:"participant_min" db:"participant_min"`
	ParticipantMax *int   `json:"participant_max" db:"participant_max"`
	ImageURL       string `json:"image_url" db:"image_url"`
}
