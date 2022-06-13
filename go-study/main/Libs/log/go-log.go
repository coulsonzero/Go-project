package log

import (
	"log"
	"time"
)

func main() {
	log.Printf("%s", time.Now())
	log.Println()
	// log.Fatal(err)
}
