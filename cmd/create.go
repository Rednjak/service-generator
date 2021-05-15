package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var Template embed.FS
var moduleName = ""

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project with a predefined layout",
	Long: `
	Create a new project with a predefined layout. 
	You need to provide the name of the project as an argument
	and the module name as a flag`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		create(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Golang module name")
	createCmd.MarkFlagRequired("module")
}

func create(serviceName string) {
	currentPath, err := os.Getwd()
	checkErr(err)

	// Create the service directory
	directory := fmt.Sprintf("%s/%s", currentPath, serviceName)
	err = os.Mkdir(directory, 0755)
	checkErr(err)

	// Walk through the embeded files and create them in the service directory
	err = fs.WalkDir(Template, "templates", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if de.IsDir() {
			return nil
		}

		fileContent, err := fs.ReadFile(Template, de.Name())
		if err != nil {
			panic(err)
		}

		filePath := fmt.Sprintf("%s/%s", directory, path)
		createDirectoryIfMissing(filePath, de)

		matched := matchGoFiles(de.Name())
		if matched {
			// Update the module_placeholder with the provided module name
			fileContent = []byte(strings.Replace(string(fileContent), "module_placeholder", moduleName, -1))

			// We needed to rename go.mod and go.sum to be able to embed them
			if ok := strings.Contains(filePath, "go.mod.template"); ok {
				filePath = strings.Replace(filePath, "go.mod.template", "go.mod", 1)
			}
			if ok := strings.Contains(filePath, "go.sum.template"); ok {
				filePath = strings.Replace(filePath, "go.sum.template", "go.sum", 1)
			}
		}

		createAndWriteToFile(filePath, fileContent)
		return nil
	})
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func matchGoFiles(fileName string) bool {
	if fileName == "go.mod.template" || fileName == "go.sum.template" {
		return true
	}

	matched, err := filepath.Match("*.go", fileName)
	if err != nil {
		panic(err)
	}

	return matched
}

func createDirectoryIfMissing(filePath string, de fs.DirEntry) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(strings.Replace(filePath, de.Name(), "", 1), 0755)
	}
}

func createAndWriteToFile(filePath string, content []byte) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	f.Write(content)
	f.Close()
}
