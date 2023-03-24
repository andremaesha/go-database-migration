package domain

type Film struct {
	FilmId      uint64 `gorm:"primarykey"`
	Title       string
	Description string
	ReleaseYear string
	Fulltext    string
}
