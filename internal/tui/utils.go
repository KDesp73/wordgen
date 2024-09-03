package tui

import (
	"encoding/json"
	"fmt"
	"os"
)


func eumod(a, b int) int {
    if b == 0 {
        fmt.Fprintf(os.Stderr, "Error: Division by zero is undefined.\n");
        return 0;
    }
    
    var r = a % b
    if (r < 0) {
		if b > 0 {
			r += b
		} else {
			r += -b
		}
    }
    return r;
}

func saveToFile(filename string, data []string) error {
	// Convert the string array to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write JSON data to the file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadFromFile(filename string) ([]string, error) {
	// Read the JSON data from the file
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Convert JSON data back to a string array
	var data []string
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
