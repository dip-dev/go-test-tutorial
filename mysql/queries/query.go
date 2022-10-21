package queries

// Query ...
type Query struct{}

// New ...
func New() *Query {
	return new(Query)
}

// SelectPrefecture ..
func (q *Query) SelectPrefecture() string {
	return `
	SELECT
		id
		, name
		, area
		, land_area
	FROM
		m_prefecture
	WHERE
		name = :name`
}

// SelectPrefectures ..
func (q *Query) SelectPrefectures() string {
	return `
	SELECT
		id
		, name
		, area
		, land_area
	FROM
		m_prefecture
	WHERE
		area = :area
	ORDER BY
		ID ASC`
}
