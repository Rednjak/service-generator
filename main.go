package main

import (
	"embed"

	"github.com/Rednjak/service-generator/cmd"
)

//go:embed templates
//go:embed templates/.editorconfig templates/.env templates/.envrc templates/.gitignore
//go:embed templates/docs/.keep templates/migrations/.keep
var templates embed.FS

func main() {
	// Inject the templates into the create function
	// We can't import files from a different package when we are in a package
	cmd.Template = templates
	cmd.Execute()
}
