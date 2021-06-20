package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("init env")
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	fmt.Println(".env loaded successfully")
}

func Get(key string) string {
	return os.Getenv(key)
}

func Set(key string, value string) {
	os.Setenv(key, value)
}
