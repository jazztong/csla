package generator

import (
	"fmt"
	"github.com/jazztong/csla/cross"
	"github.com/jazztong/csla/provider/fileloader"
	"io/ioutil"
	"os"
	"path"
)

type awsAPILambdaGolang struct {
	TemplatePrefix string
}

func (t *awsAPILambdaGolang) Generate(req Request) {
	// Create folder
	if err := os.Mkdir(req.Name, os.ModePerm); err != nil {
		panic(err)
	}
	for _, v := range t.Files() {
		if v.Generate(req) {
			binary, err := fileloader.Load(path.Join(t.TemplatePrefix, fmt.Sprintf("%s.tmpl", v.Name)))
			if err != nil {
				cross.Error(err.Error())
				panic(err)
			}
			err = ioutil.WriteFile(path.Join(req.Name, v.Name), binary, os.ModeDevice)
			if err != nil {
				panic(err)
			}
		}
	}
	cross.Print("Init project %s successful", req.Name)
}

// NewAwsAPILambdaGolang create new generator for aws api lambda golang
func NewAwsAPILambdaGolang() Generator {
	return &awsAPILambdaGolang{TemplatePrefix: "template/aws-api-lambda-golang"}
}

func (t *awsAPILambdaGolang) Files() []File {
	return []File{
		File{Name: ".gitignore", Generate: func(req Request) bool { return true }},
		File{Name: "main.go", Generate: func(req Request) bool { return true }},
		File{Name: "Makefile", Generate: func(req Request) bool { return true }},
	}
}
