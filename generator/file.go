package generator

// File represent the small unit of template
type File struct {
	Name     string
	Generate func(req Request) bool
}
