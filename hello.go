package main
	
import (
    "fmt"
    "os"
)

func main() {
    fmt.Printf("config=%s secret=%s\n", os.Getenv("HELLO_CONFIG"), os.Getenv("HELLO_CONFIG"))
}
