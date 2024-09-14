package main

import (
	"os"
	"os/exec"
	"testing"
)

// Helper function to create a temporary file with the given content
func createTempFile(content string) (string, error) {
	tmpFile, err := os.CreateTemp("", "testfile-*.json")
	if err != nil {
		return "", err
	}
	if _, err := tmpFile.WriteString(content); err != nil {
		return "", err
	}
	if err := tmpFile.Close(); err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}

func TestValidJSON(t *testing.T) {
	jsonContent := `{"name": "John Doe", "age": 30, "email": "john.doe@example.com"}`
	fileName, err := createTempFile(jsonContent)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(fileName) // Clean up the temp file

	cmd := exec.Command("go", "run", "main.go", fileName)
	if err := cmd.Run(); err != nil {
		t.Errorf("Expected valid JSON but got error: %v", err)
	}
}

func TestInvalidJSON(t *testing.T) {
	jsonContent := `{"name": "John Doe", "age": 30, "email": "john.doe@example.com"`
	fileName, err := createTempFile(jsonContent)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(fileName) // Clean up the temp file

	cmd := exec.Command("go", "run", "main.go", fileName)
	if err := cmd.Run(); err == nil {
		t.Error("Expected error for invalid JSON but got none")
	}
}
