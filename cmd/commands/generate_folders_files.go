package commands

import (
	"codeCrafter/cmd"
	"codeCrafter/logger"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Event struct {
		CreatingPackageFiles struct {
			State string   `yaml:"state"`
			Path  []string `yaml:"flow"`
		} `yaml:"creating_package_files"`
		LinkingClass []string `yaml:"linking_class"`
	} `yaml:"event"`
}

var generateFoldersAndFiles = &cobra.Command{
	Use:   "genFoFI",
	Short: "Generate folder and files declared in codeCrafter.yml",
	Run:   genFolderAndFiles,
}

func genFolderAndFiles(cmd *cobra.Command, args []string) {
	projectDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		logger.ErrorLogger.Printf("Error getting current working directory: %v", err)
		return
	}

	configPath := filepath.Join(projectDir, "codeCrafter.yml")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %v\n", err)
		logger.ErrorLogger.Printf("Error reading YAML file: %v", err)
		return
	}
	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("Error parsing YAML file %v", err)
		logger.ErrorLogger.Printf("Error parsing Yaml file %v", err)
		return
	}
	if config.Event.CreatingPackageFiles.State != "on" {
		fmt.Println("Structure generation is disabled. Set 'on' to enable it.")
		return
	}
	for _, path := range config.Event.CreatingPackageFiles.Path {
		createStructure(path)
	}
	logger.InfoLogger.Printf("Folder and files created.")
	fmt.Printf("Folder and files created successfully.\n")
}

func createStructure(path string) {
	if strings.Contains(path, "[") && strings.Contains(path, "]") {
		parts := strings.Split(path, "[")
		dir := strings.TrimSpace(parts[0])
		filesPart := strings.TrimRight(parts[1], "]")

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}

		files := strings.Split(filesPart, ",")
		for _, file := range files {
			file = strings.TrimSpace(file)
			filePath := filepath.Join(dir, file)
			_, err := os.Create(filePath)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", filePath, err)
				return
			} else {
				fmt.Printf("Created file %s\n", filePath)
			}
		}
	} else {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		} else {
			fmt.Printf("Directory %s exists or created\n", path)
		}
	}
}

func init() {
	cmd.RootCmd.AddCommand(generateFoldersAndFiles)
}
