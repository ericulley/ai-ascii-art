package data

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type AsciiRecord struct {
	Name string
	Art  string
}

/*
 *  Save function
 */
func SaveArtToDB(ascii AsciiRecord) error {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(`INSERT INTO ascii (id, name, art) VALUES (?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(nil, ascii.Name, ascii.Art)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println("Successfully inserted record into db: ", res)
	return nil
}

/*
 *  List functions
 */

func ListAllArt() error {
	fmt.Println("In ListAllArt")
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM ascii`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var art string
		if err := rows.Scan(&id, &name, &art); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, art)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func ListArtWithLimit(limit int) error {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM ascii LIMIT ?`, limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var art string
		if err := rows.Scan(&id, &name, &art); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, art)
	}
	return nil 
}

func ListArtNames() error {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT id, name FROM ascii`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	return nil
}

func ListArtNamesWithLimit(limit int) error {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT id, name FROM ascii LIMIT ?`, limit)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	return nil
}

/*
 *  Delete functions
 */

 func DeleteArtById(id int) error {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(`DELETE FROM ascii WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if rows, err := res.RowsAffected(); err == nil && rows > 0 {
		fmt.Println("Successfully deleted record from db")
	} else if err != nil {
		fmt.Println("Error deleting record from db: ", err)
	} else if rows == 0 {
		fmt.Println("No record found with id: ", id)
	}
	return nil
 }

 func DeleteArtByName(name string) error {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(`DELETE FROM ascii WHERE name = ?`)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(name)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if rows, err := res.RowsAffected(); err == nil && rows > 0 {
		fmt.Println("Successfully deleted record from db")
	} else if err != nil {
		fmt.Println("Error deleting record from db: ", err)
	} else if rows == 0 {
		fmt.Println("No record found with name: ", name)
	}
	return nil
 }

 /*
 *  Update functions
 */
func UpdateArt(fromNameOrId string, toName string) error {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var res sql.Result
	fromId, err := strconv.Atoi(fromNameOrId)
    if err != nil {
		// value is a string
        stmt, err := db.Prepare(`UPDATE ascii SET name = ? WHERE name = ?`)
		if err != nil {
			log.Fatal(err)
		}
		res, err = stmt.Exec(toName, fromNameOrId)
		if err != nil {
			return err
		}
		defer stmt.Close()
    } else {
		// value is an integer
        stmt, err := db.Prepare(`UPDATE ascii SET name = ? WHERE id = ?`)
		if err != nil {
			log.Fatal(err)
		}
		res, err = stmt.Exec(toName, fromId)
		if err != nil {
			return err
		}
		defer stmt.Close()
    }
	if rows, err := res.RowsAffected(); err == nil && rows > 0 {
		fmt.Println("Successfully updated record in db")
	} else if err != nil {
		fmt.Println("Error updating record in db: ", err)
	} else if rows == 0 {
		fmt.Println("No record found with name or id: ", fromNameOrId)
	}
	return nil
}

/*
 *  Art function (generates random ascii art from db)
 */
func Art() string {
	db, err := sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt, err := db.Prepare(`SELECT art FROM ascii ORDER BY RANDOM() LIMIT 1`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var art string
	err = stmt.QueryRow().Scan(&art);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(art)
	return art
}