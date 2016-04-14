package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/daniel-fanjul-alcuten/amar-dashboard-backend"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	log.SetFlags(0)
	driver := flag.String("d", "mysql", "sql driver")
	host := flag.String("h", "garak:grapes@tcp(127.0.0.1:3306)/amar", "user:password@tcp(host:port)/dbname")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
	d := json.NewDecoder(r)
	var p amar.MyStuffPage
	if err := d.Decode(&p); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(*driver, *host)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = save(p, db); err != nil {
		log.Fatal(err)
	}
}

func save(p amar.MyStuffPage, db *sql.DB) (err error) {

	var tx *sql.Tx
	if tx, err = db.Begin(); err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	if _, err = tx.Exec("DELETE FROM `FETCH` WHERE TRUE"); err != nil {
		return
	}
	if _, err = tx.Exec("INSERT INTO `FETCH` (`LAST_FETCH`) VALUES (?)", p.Time); err != nil {
		return
	}

	if _, err = tx.Exec("UPDATE `ITEM` SET `ITEM_COUNT` = 0"); err != nil {
		return
	}
	for n, s := range p.Stuff {
		if _, err = tx.Exec("INSERT INTO `ITEM` (`ITEM_NAME`, `ITEM_LINK`, `ITEM_COUNT`) VALUES (?, ?, ?)"+
			" ON DUPLICATE KEY UPDATE `ITEM_COUNT` = `ITEM_COUNT` + ?", n, s.Link, s.Guild, s.Guild); err != nil {
			return
		}
	}

	return
}
