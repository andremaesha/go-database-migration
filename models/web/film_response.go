package web

type FilmResponse struct {
	Message     string `json:"message"`
	Code        int    `json:"code"`
	IdFilm      int    `json:"id_film"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear string `json:"release_year"`
	FullText    string `json:"full_text"`
}
