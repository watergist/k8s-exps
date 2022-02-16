package app

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Started Application")
	time.Sleep(time.Hour)
	os.Open("ysafsd.json")
	fmt.Println("Exited Application")
}
