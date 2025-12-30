package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// COURSE 5: FILE HANDLING AND I/O
// Topics covered:
// 1. Reading files
// 2. Writing files
// 3. Appending to files
// 4. Reading line by line
// 5. File information
// 6. Directory operations
// 7. Copying files
// 8. Working with paths
// 9. Buffered I/O

// ============ 1. READ ENTIRE FILE ============
func readFileContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ============ 2. WRITE FILE ============
func writeToFile(filename string, content string) error {
	// Create or truncate file
	return os.WriteFile(filename, []byte(content), 0644)
}

// ============ 3. APPEND TO FILE ============
func appendToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// ============ 4. READ LINE BY LINE ============
func readLineByLine(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// ============ 5. READ WITH BUFFER ============
func readWithBuffer(filename string, bufferSize int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, bufferSize)
	var result strings.Builder

	for {
		chunk, err := reader.ReadString('\n')
		result.WriteString(chunk)

		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}

	return result.String(), nil
}

// ============ 6. FILE INFORMATION ============
func getFileInfo(filename string) error {
	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	fmt.Printf("Filename: %s\n", info.Name())
	fmt.Printf("Size: %d bytes\n", info.Size())
	fmt.Printf("Modified: %v\n", info.ModTime())
	fmt.Printf("Is Directory: %v\n", info.IsDir())
	fmt.Printf("Permissions: %v\n", info.Mode())

	return nil
}

// ============ 7. CHECK IF FILE EXISTS ============
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// ============ 8. LIST DIRECTORY CONTENTS ============
func listDirectory(dirPath string) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name())
	}

	return files, nil
}

// ============ 9. CREATE DIRECTORY ============
func createDirectory(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

// ============ 10. COPY FILE ============
func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

// ============ 11. DELETE FILE ============
func deleteFile(filename string) error {
	return os.Remove(filename)
}

// ============ 12. WORK WITH PATHS ============
func pathOperations(filePath string) {
	fmt.Printf("Full path: %s\n", filePath)
	fmt.Printf("Directory: %s\n", filepath.Dir(filePath))
	fmt.Printf("Filename: %s\n", filepath.Base(filePath))
	fmt.Printf("Extension: %s\n", filepath.Ext(filePath))

	// Join paths correctly for the OS
	newPath := filepath.Join(".", "data", "file.txt")
	fmt.Printf("Joined path: %s\n", newPath)
}

// ============ 13. CSV-LIKE FILE OPERATIONS ============
func parseCSVFile(filename string) ([][]string, error) {
	lines, err := readLineByLine(filename)
	if err != nil {
		return nil, err
	}

	var records [][]string
	for _, line := range lines {
		fields := strings.Split(line, ",")
		records = append(records, fields)
	}

	return records, nil
}

// ============ COURSE FIVE MAIN FUNCTION ============
func courseFive() {
	fmt.Println("=== FILE HANDLING AND I/O ===\n")

	tempDir := "./temp"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir) // Cleanup after demo

	// ============ 1. WRITE FILE ============
	fmt.Println("1. WRITE FILE")
	fmt.Println("---")

	testFile := filepath.Join(tempDir, "test.txt")
	content := "Hello, Go!\nThis is a test file.\nWelcome to file handling!"

	err := writeToFile(testFile, content)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	} else {
		fmt.Printf("✓ File written: %s\n\n", testFile)
	}

	// ============ 2. READ ENTIRE FILE ============
	fmt.Println("2. READ ENTIRE FILE")
	fmt.Println("---")

	data, err := readFileContents(testFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	} else {
		fmt.Printf("File contents:\n%s\n\n", data)
	}

	// ============ 3. READ LINE BY LINE ============
	fmt.Println("3. READ LINE BY LINE")
	fmt.Println("---")

	lines, err := readLineByLine(testFile)
	if err != nil {
		fmt.Printf("Error reading lines: %v\n", err)
	} else {
		fmt.Println("Lines:")
		for i, line := range lines {
			fmt.Printf("  Line %d: %s\n", i+1, line)
		}
		fmt.Println()
	}

	// ============ 4. APPEND TO FILE ============
	fmt.Println("4. APPEND TO FILE")
	fmt.Println("---")

	appendContent := "\nAppended line 1\nAppended line 2"
	err = appendToFile(testFile, appendContent)
	if err != nil {
		fmt.Printf("Error appending: %v\n", err)
	} else {
		fmt.Printf("✓ Content appended\n")
		updatedData, _ := readFileContents(testFile)
		fmt.Printf("Updated contents:\n%s\n\n", updatedData)
	}

	// ============ 5. FILE INFORMATION ============
	fmt.Println("5. FILE INFORMATION")
	fmt.Println("---")

	err = getFileInfo(testFile)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
	}
	fmt.Println()

	// ============ 6. CHECK IF FILE EXISTS ============
	fmt.Println("6. CHECK IF FILE EXISTS")
	fmt.Println("---")

	exists := fileExists(testFile)
	fmt.Printf("File exists: %v\n", exists)

	notExists := fileExists("nonexistent.txt")
	fmt.Printf("Nonexistent file exists: %v\n\n", notExists)

	// ============ 7. CREATE DIRECTORY ============
	fmt.Println("7. CREATE DIRECTORY")
	fmt.Println("---")

	newDir := filepath.Join(tempDir, "subdir", "nested")
	err = createDirectory(newDir)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
	} else {
		fmt.Printf("✓ Directory created: %s\n\n", newDir)
	}

	// ============ 8. LIST DIRECTORY ============
	fmt.Println("8. LIST DIRECTORY CONTENTS")
	fmt.Println("---")

	files, err := listDirectory(tempDir)
	if err != nil {
		fmt.Printf("Error listing directory: %v\n", err)
	} else {
		fmt.Printf("Contents of %s:\n", tempDir)
		for _, file := range files {
			fmt.Printf("  - %s\n", file)
		}
		fmt.Println()
	}

	// ============ 9. COPY FILE ============
	fmt.Println("9. COPY FILE")
	fmt.Println("---")

	copiedFile := filepath.Join(tempDir, "test_copy.txt")
	err = copyFile(testFile, copiedFile)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
	} else {
		fmt.Printf("✓ File copied from %s to %s\n", testFile, copiedFile)

		exists = fileExists(copiedFile)
		fmt.Printf("Copied file exists: %v\n\n", exists)
	}

	// ============ 10. PATH OPERATIONS ============
	fmt.Println("10. PATH OPERATIONS")
	fmt.Println("---")

	examplePath := "/home/user/documents/report.pdf"
	pathOperations(examplePath)
	fmt.Println()

	// ============ 11. CSV-LIKE FILE ============
	fmt.Println("11. PARSE CSV-LIKE FILE")
	fmt.Println("---")

	csvFile := filepath.Join(tempDir, "data.csv")
	csvContent := `Name,Age,City
Alice,30,New York
Bob,25,Los Angeles
Charlie,35,Chicago`

	writeToFile(csvFile, csvContent)

	records, err := parseCSVFile(csvFile)
	if err != nil {
		fmt.Printf("Error parsing CSV: %v\n", err)
	} else {
		fmt.Println("CSV Data:")
		for i, record := range records {
			fmt.Printf("  Row %d: %v\n", i+1, record)
		}
		fmt.Println()
	}

	// ============ 12. DELETE FILE ============
	fmt.Println("12. DELETE FILE")
	fmt.Println("---")

	err = deleteFile(copiedFile)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
	} else {
		fmt.Printf("✓ File deleted: %s\n", copiedFile)
		exists = fileExists(copiedFile)
		fmt.Printf("File exists after deletion: %v\n\n", exists)
	}

	fmt.Println("=== END OF FILE HANDLING ===")
}

// KEY TAKEAWAYS:
// 1. os.ReadFile() reads entire file into memory (simple, not for huge files)
// 2. os.WriteFile() creates/truncates and writes to file
// 3. Use os.OpenFile with flags for more control (append, etc.)
// 4. Always defer file.Close() to prevent resource leaks
// 5. bufio.Scanner is great for reading line-by-line
// 6. io.Copy() is efficient for copying between readers/writers
// 7. filepath package handles paths correctly for your OS
// 8. Check errors! File operations often fail
// 9. Use os.Stat() to get file information and check existence
// 10. os.ReadDir() for listing directory (not os.ReadFile!)
// 11. Be careful with file permissions (0644 for files, 0755 for dirs)
// 12. Delete files carefully - they're gone permanently!
// 13. Use buffered I/O for better performance with large files
// 14. Error handling is crucial in file operations
// 15. Always clean up temporary files and directories
