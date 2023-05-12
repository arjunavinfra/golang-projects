
type Director struct {
	Fristname string `json:"firstname"`
	LastName  string `json:"lastname"`
}
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"lsbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
