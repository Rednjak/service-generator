package cmd

import (
	"fmt"
	"os"

	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
)

var sourcePath = fmt.Sprintf("%s/go/src/%s", os.Getenv("HOME"), "github.com/Rednjak/service-generator")

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		create(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func create(serviceName string) {
	fmt.Println(serviceName)
	currentPath, err := os.Getwd()
	check(err)
	// if err != nil {
	// 	return fmt.Errorf("Issues getting current path. Trace: %w", err)
	// }

	// Create the service directory
	directory := fmt.Sprintf("%s/%s", currentPath, serviceName)
	err = os.Mkdir(directory, 0755)
	check(err)

	// fmt.Println(sourcePath)
	// fmt.Println(getExecPath())
	// fmt.Println(os.Getenv("HOME"))
	// fmt.Println(sourcePath)
	// createFiles(directory)

	err = copy.Copy(sourcePath+"/templates", directory)
	check(err)
}

// func createFiles(directory string) error {
// 	files, err := ioutil.ReadDir(sourcePath + "/templates")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		fmt.Println(file.Name())
// 		destination, err := os.Create(directory + "/" + file.Name())
// 		if err != nil {
// 			check(err)
// 		}

// 		source, err := os.Open(sourcePath + "/templates/" + file.Name())
// 		if err != nil {
// 			check(err)
// 		}
// 		defer source.Close()

// 		defer destination.Close()
// 		_, err = io.Copy(destination, source)
// 		check(err)
// 	}

// 	return nil
// }

// func getExecPath() string {
// 	fmt.Println("getExecPath")
// 	ex, err := os.Executable()
// 	check(err)
// 	fmt.Println(filepath.Dir(ex))
// 	return filepath.Dir(ex)
// }
