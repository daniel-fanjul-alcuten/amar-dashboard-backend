package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"
)

func main() {

	log.SetFlags(0)
	p := flag.String("p", "", "pid")
	u := flag.String("u", "", "uid")
	flag.Parse()

	now := time.Now()
	fetcher := MyStuffFetcher{Pid: *p, Uid: *u}
	str, err := fetcher.Fetch()
	if err != nil {
		log.Fatal(err)
	}

	myStuff, err := fetcher.Parse(now, str)
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	e := json.NewEncoder(w)
	if err = e.Encode(myStuff); err != nil {
		log.Fatal(err)
	}
}
