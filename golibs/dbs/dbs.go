package dbs


import (
	"database/sql"
    "fmt"
	"github.com/mattn/go-sqlite3" //sqlite3 support
)

func Query(ind string) (outd string) {
	db, err := sql.Open("sqlite3,","./main.db")
	checkErr(err)
	
	rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var kuid string
        var pubkey string
        var keyid string
        err = rows.Scan(&kuid, &pubkey, &keyid)
        checkErr(err)
        fmt.Println(kuid)
        fmt.Println(pubkey)
        fmt.Println(keyid)
    }
    db.Close()
}

func Insert(kuid, pubkey, keyid string) {
	stmt, err := db.Prepare("INSERT INTO users(kuid, pubkey, keyid) values(?,?,?)")
    checkErr(err)

    res, err := stmt.Exec(kuid, pubkey, keyid)
    checkErr(err)
    db.Close()
}

func Update() {
	stmt, err = db.Prepare("update users set pubkey=? and keyid=? where kuid=?")
    checkErr(err)

    res, err = stmt.Exec(pubkey, keyid, kuid)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)
	db.Close()
}

func Delete () {
    stmt, err = db.Prepare("delete from users where kuid=?")
    checkErr(err)

    res, err = stmt.Exec(kuid)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)
    db.Close()
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}