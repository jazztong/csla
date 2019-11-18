package fileloader

import "testing"

func TestLoad(t *testing.T) {
	tests := []string{
		"template/aws-api-lambda-golang/.gitignore.tmpl",
		"template/aws-api-lambda-golang/main.go.tmpl",
		"template/aws-api-lambda-golang/Makefile.tmpl",
	}
	for _, tt := range tests {
		t.Run("TestLoad", func(t *testing.T) {
			_, err := Load(tt)
			if err != nil {
				t.Errorf("Load() error = %v", err)
				return
			}
		})
	}
}
