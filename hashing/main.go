package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"io"
	"time"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var hash string
	var rootCmd = &cobra.Command{
		Use:   "hash",
		Short: "",
		Long:  "",
	}
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Tests multiple hashing algorithms",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println()
			md5Hash(hash)
			sha1Hash(hash)
			sha2Hash(hash)
			fmt.Println()
			fmt.Print("Password Hashing")
			fmt.Println()
			bcryptHash(hash)
			argonHash(hash)
		},
	}
	runCmd.Flags().StringVarP(&hash, "hash", "s", "", "Specify a string to hash")
	rootCmd.AddCommand(runCmd)
	rootCmd.Execute()
}

func md5Hash(value string) {
	startTime := time.Now()
	hash := md5.New()
	io.WriteString(hash, value)
	endTime := time.Now()
	timeDiff := endTime.Sub(startTime)
	hashValue := hex.EncodeToString(hash.Sum(nil))
	fmt.Print("MD5 Hash: " + hashValue + " " + " Duration: ")
	fmt.Print(timeDiff)
	fmt.Println()
}

func sha1Hash(value string) {
	startTime := time.Now()
	hash := sha1.New()
	io.WriteString(hash, value)
	endTime := time.Now()
	timeDiff := endTime.Sub(startTime)
	hashValue := hex.EncodeToString(hash.Sum(nil))
	fmt.Print("SHA1 Hash: " + hashValue + " " + " Duration: ")
	fmt.Print(timeDiff)
	fmt.Println()
}

func sha2Hash(value string) {
	startTime := time.Now()
	hash := sha256.New()
	hash.Write([]byte(value))
	endTime := time.Now()
	timeDiff := endTime.Sub(startTime)
	hashValue := hex.EncodeToString(hash.Sum(nil))
	fmt.Print("SHA256 Hash: " + hashValue + " " + " Duration: ")
	fmt.Print(timeDiff)
	fmt.Println()
}

func bcryptHash(value string) {
	startTime := time.Now()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	endTime := time.Now()
	timeDiff := endTime.Sub(startTime)
	fmt.Print("Bcrypt Hash: " + string(bytes) + " " + " Duration: ")
	fmt.Print(timeDiff)
	fmt.Println()
}

func argonHash(value string) {
	var (
		memory      uint32 = 64 * 1024
		iterations  uint32 = 3
		parallelism uint8  = 2
		saltLength         = 16
		keyLength   uint32 = 32
	)
	startTime := time.Now()
	salt, _ := randomString(saltLength)
	bytes := argon2.IDKey([]byte(value), []byte(salt), iterations, memory, parallelism, keyLength)
	base64Salt := base64.RawStdEncoding.EncodeToString([]byte(salt))
	base64Hash := base64.RawStdEncoding.EncodeToString(bytes)
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, parallelism, base64Salt, base64Hash)
	endTime := time.Now()
	timeDiff := endTime.Sub(startTime)
	fmt.Print("Argon2 Hash: " + encodedHash + " " + " Duration: ")
	fmt.Print(timeDiff)
	fmt.Println()
}

func randomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	return string(bytes), nil
}
