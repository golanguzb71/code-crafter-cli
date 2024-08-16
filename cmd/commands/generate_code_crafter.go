package commands

import (
	"codeCrafter/cmd"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var generateCrafter = &cobra.Command{
	Use:   "genCrafter",
	Short: "Generate codeCrafter.yml file for project structure",
	Run:   genCra,
}

func genCra(cmd *cobra.Command, args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current location %v", err)
	}
	filePath := filepath.Join(dir, "codeCrafter.yml")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	content := `
event:   
  creating_package_files:
    state: on
    flow:
      - config/[ApplicationConfig.java,DataLoaderConfig.java,OpenApiConfig.java]
      - controller
      - entity
      - exception
      - jwt/[JwtFilter.java,JwtProvider.java]
      - mapper
      - payload
      - repository
      - service/[MockService.java]
      - service/impl/[MockServiceImpl.java]
  linking_class:
    - entity| Entity.java
    - controller| Controller.java
    - mapper| Mapper.java
    - payload| Payload.java
    - repository| Repository.java
    - service| Service.java
    - service/impl| ServiceImpl.java
`
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to file: %v", err)
		return
	}
	fmt.Printf("codeCrafter.yml file created successfully in %s", filePath)
}

func init() {
	cmd.RootCmd.AddCommand(generateCrafter)
}
