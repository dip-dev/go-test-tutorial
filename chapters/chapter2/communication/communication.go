package communication

// String ...
type String struct{}

// New ...
func New(str string) *String {
	return &String{}
}

// Greeting ..
func (s *String) Greeting() string {
	return "Hello!!"
}
