package internal

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
	"unicode"
)

func CountLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, nil
}

func CountLinesFromStdin() int {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

func CalculateChecksum(filename, algorithm string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var h hash.Hash
	switch algorithm {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	default:
		return "", fmt.Errorf("unsupported checksum algorithm")
	}

	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func ChecksumFromStdin(algorithm string) string {
	var h hash.Hash
	switch algorithm {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	default:
		return "unsupported checksum algorithm"
	}

	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		return "error: failed to calculate checksum"
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func CheckFileValid(file string) error {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return fmt.Errorf("No such file '%s'\n", file)
	}

	if info.IsDir() {
		return fmt.Errorf("Expected file got directory '%s'\n", file)
	}

	if isBinaryFile(file) {
		return fmt.Errorf("Cannot do linecount or checksum for binary file '%s'\n", file)
	}
	return nil
}

func isBinaryFile(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 8000)
	n, _ := reader.Read(buffer)

	for i := 0; i < n; i++ {
		if !unicode.IsPrint(rune(buffer[i])) && buffer[i] != '\n' && buffer[i] != '\r' && buffer[i] != '\t' {
			return true
		}
	}
	return false
}
