package solutions

import (
	"os"
	"bufio"
	"strings"
	"path/filepath"
)


func FilterLogs(srcPath, destPath, keyword string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dest)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			if _, err := writer.WriteString(line + "\n"); err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

func CountWords(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func FindFiles(dirPath, ext string) ([]string, error) {
	var matches []string
	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ext {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			matches = append(matches, absPath)
		}
		return nil
	})
	return matches, err
}
