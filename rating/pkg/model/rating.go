package model

type RecordType string

type RecordID string

type UserID string

type RatingValue int

const (
	RecordTypeMovie = RecordType("movie")
)

type Rating struct {
	ID         string     `json:"id"`
	RecordID   RecordID   `json:"recordId"`
	RecordType RecordType `json:"recordType"`
	UserID     UserID     `json:"userId"`
	Value      float64    `json:"value"`
}
