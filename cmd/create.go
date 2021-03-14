package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var sourcePath = fmt.Sprintf("%s/go/src/%s", os.Getenv("HOME"), "github.com/Rednjak/service-generator")
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

	// Copy the service tempalte
	err = copy.Copy(sourcePath+"/templates", directory)
	checkErr(err)

	// Update the module placeholders with the provided module name
	err = filepath.Walk(directory, visit)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func match(fileName string) bool {
	if fileName == "go.mod" {
		return true
	}

	matched, err := filepath.Match("*.go", fileName)
	if err != nil {
		panic(err)
	}

	return matched
}

func visit(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil
	}

	matched := match(fi.Name())
	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		newContents := strings.Replace(string(read), "github.com/repo/module", moduleName, -1)
		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
