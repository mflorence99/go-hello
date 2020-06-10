package main
	
import (
    "bufio"
    "fmt"
    "net/http"
    "os"
)

func main() {

    fmt.Printf("Http's host=%s port=%s\n\n", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))

    resp, err := http.Get(fmt.Sprintf("http://%s:%s/world", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT")))
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    	
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan(); i++ {
        fmt.Println(scanner.Text())
    }

}
