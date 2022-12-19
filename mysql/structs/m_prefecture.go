package structs

// MPrefecture ..
type MPrefecture struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Area     string `db:"area"`
	LandArea int    `db:"land_area"`
}
