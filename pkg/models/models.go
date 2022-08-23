package models

type Movie struct {
	Id          string    `json:"id"`
	MovieName   string    `json:"name"`
	ReleaseDate string    `json:"relesedate"`
	Director    *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
