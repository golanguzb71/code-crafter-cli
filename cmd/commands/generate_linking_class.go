package commands

import (
	"codeCrafter/cmd"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var generateLinkingClasses = &cobra.Command{
	Use:   "gen",
	Short: "Generate Linking classes",
	Run:   genLiClass,
}

func genLiClass(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: No filename provided.")
		return
	}

	filename := args[0]

	projectDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}

	configPath := filepath.Join(projectDir, "codeCrafter.yml")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %v\n", err)
		return
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		fmt.Printf("Error parsing YAML file: %v\n", err)
		return
	}

	if len(config.Event.LinkingClass) == 0 {
		fmt.Println("Error: No linking class entries found.")
		return
	}

	mainDir, mainFileTemplate := parseLinkingClassEntry(config.Event.LinkingClass[0])
	className := filename

	mainFileName := className + strings.TrimPrefix(mainFileTemplate, " ")
	mainFilePath := filepath.Join(projectDir, mainDir, mainFileName)
	if err := ensureDirExists(filepath.Join(projectDir, mainDir)); err != nil {
		fmt.Printf("Error ensuring directory %s exists: %v\n", mainDir, err)
		return
	}
	if err := createFileWithSuffix(mainFilePath, className); err != nil {
		fmt.Printf("Error creating file %s: %v\n", mainFilePath, err)
		return
	}

	for _, relatedClass := range config.Event.LinkingClass {
		dir, template := parseLinkingClassEntry(relatedClass)
		if dir == mainDir {
			continue
		}

		fileName := className + strings.TrimPrefix(template, " ")
		relatedFilePath := filepath.Join(projectDir, dir, fileName)
		if err := ensureDirExists(filepath.Join(projectDir, dir)); err != nil {
			fmt.Printf("Error ensuring directory %s exists: %v\n", dir, err)
			continue
		}
		if err := createFileWithSuffix(relatedFilePath, className); err != nil {
			fmt.Printf("Error creating file %s: %v\n", relatedFilePath, err)
		}
	}

	fmt.Printf("Linking classes created successfully.\n")
}

func parseLinkingClassEntry(entry string) (dir, template string) {
	parts := strings.SplitN(entry, "|", 2)
	if len(parts) != 2 {
		fmt.Printf("Invalid linking class format: %s\n", entry)
		return "", ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func ensureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}

func createFileWithSuffix(filePath, className string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	content := fmt.Sprintf("//It is auto generated. Edit it as you want. Also join my telegram channel => https://t.me/kuyov_taraf")
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Printf("Created file %s\n", filePath)
	return nil
}

func init() {
	cmd.RootCmd.AddCommand(generateLinkingClasses)
}
