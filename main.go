package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	// Define flags for the mode, input file, and output file
	mode := flag.String("mode", "", "Conversion mode: 'json2yaml' or 'yaml2json'")
	inputFile := flag.String("input", "", "Path to the input file (JSON or YAML)")
	outputFile := flag.String("output", "", "Path to the output file (YAML or JSON)")

	// Parse the command-line flags
	flag.Parse()

	// Check if the mode, input, and output flags are provided
	if *mode == "" || *inputFile == "" || *outputFile == "" {
		log.Fatalf("Usage: %s --mode=<json2yaml|yaml2json> --input=<input_file> --output=<output_file>\n", os.Args[0])
	}

	// Perform the appropriate conversion based on the mode
	switch *mode {
	case "json2yaml":
		convertJSONToYAML(*inputFile, *outputFile)
	case "yaml2json":
		convertYAMLToJSON(*inputFile, *outputFile)
	default:
		log.Fatalf("Invalid mode: %s. Use 'json2yaml' or 'yaml2json'.", *mode)
	}
}

func convertJSONToYAML(inputFile, outputFile string) {
	// Read JSON file
	jsonData, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal JSON into a generic map
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Marshal data into YAML
	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatalf("Failed to convert to YAML: %v", err)
	}

	// Write YAML file
	err = os.WriteFile(outputFile, yamlData, 0644)
	if err != nil {
		log.Fatalf("Failed to write YAML file: %v", err)
	}

	fmt.Printf("Successfully converted %s to %s\n", inputFile, outputFile)
}

func convertYAMLToJSON(inputFile, outputFile string) {
	// Read YAML file
	yamlData, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// Unmarshal YAML into a generic map
	var data map[string]interface{}
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}

	// Marshal data into JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to convert to JSON: %v", err)
	}

	// Write JSON file
	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}

	fmt.Printf("Successfully converted %s to %s\n", inputFile, outputFile)
}
