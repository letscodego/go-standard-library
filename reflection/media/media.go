package media

import "strings"

type Catalogable interface {
	NewMovie(title string, rating Rating, boxOffice float32)
	GetTitle() string
	GetRating() string
	GetBoxOffice() float32
	SetTitle(title string)
	SetRating(rating Rating)
	SetBoxOffice(boxOffice float32)
}

type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

type Rating string

const (
	R    = "R (Restricted)"
	G    = "G (General Audience)"
	PG   = "PG (Parental Guidance)"
	PG13 = "PG-13 (Parental Caution)"
	NC17 = "NC-17 (No Children under 17)"
)

func (m *Movie) NewMovie(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}

func (m *Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m *Movie) GetRating() string {
	return string(m.rating)
}

func (m *Movie) GetBoxOffice() float32 {
	return m.boxOffice
}

func (m *Movie) SetTitle(title string) {
	m.title = title
}

func (m *Movie) SetRating(rating Rating) {
	m.rating = rating
}

func (m *Movie) SetBoxOffice(boxOffice float32) {
	m.boxOffice = boxOffice
}
