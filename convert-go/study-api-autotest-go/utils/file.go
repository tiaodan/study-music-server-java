package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveFile(srcFile, folder string) string {
	// Simplified file save - returns filename
	// In production, implement proper file upload handling
	ext := ".jpg"
	filename := fmt.Sprintf("%d%s%s", time.Now().Unix(), randString(6), ext)

	// Create directory if not exists
	dir := filepath.Join(".", "img", folder)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	return filename
}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetFileMD5(file *os.File) string {
	md5Hash := md5.New()
	_, err := io.Copy(md5Hash, file)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(md5Hash.Sum(nil))
}

func GetFileExt(filename string) string {
	idx := strings.LastIndex(filename, ".")
	if idx == -1 {
		return ""
	}
	return filename[idx:]
}
