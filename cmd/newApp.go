package cmd

import (
	"os"

	"github.com/jazztong/csla/cross"
	"github.com/jazztong/csla/generator"
	"github.com/spf13/cobra"
)

var newAppCmd = &cobra.Command{
	Use:   "newapp",
	Short: "Create new application --name=[name] --template=aws-api-lambda-golang",
	Long:  `All software has versions. This is csla`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			name = cmd.Flag(name).Value.String()
			temp = cmd.Flag(template).Value.String()
			gen  generator.Generator
		)
		if _, err := os.Stat(name); !os.IsNotExist(err) {
			cross.Error("Folder %s already exist, please run with empty folder.\n", name)
		}
		gen, ok := templates[temp]
		if !ok {
			cross.Error("Template %s currently not support, please stay tune or help me to create template.\n", temp)
		}
		gen.Generate(generator.Request{Name: name})
	},
}

// flags
const (
	name     = "name"
	template = "template"
)

var (
	templates = map[string]generator.Generator{
		"aws-api-lambda-golang": generator.NewAwsAPILambdaGolang(),
		"gcp-http-func-golang":  nil,
	}
)

func init() {
	// name flag
	newAppCmd.Flags().StringP(name, "n", "", "Name of the new application, also indicate the folder of new application, it should be not existed folder")
	if err := newAppCmd.MarkFlagRequired(name); err != nil {
		panic(err)
	}
	// template
	newAppCmd.Flags().StringP(template, "t", "", "project template use for the application, option available `aws-api-lambda-golang | gcp-http-func-golang`")
	if err := newAppCmd.MarkFlagRequired(template); err != nil {
		panic(err)
	}
}
