package fileloader

// Load is the function to load binary from path
func Load(filepath string) ([]byte, error) {
	return Asset(filepath)
}
