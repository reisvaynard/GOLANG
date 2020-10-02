package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Users struct {
	Id         int
	Username   string
	Password   string
	Name       string
	Email      string
	Priviledge string
}

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		login(w, r)
		return
	} else if r.URL.Path == "/login" {
		login(w, r)
		return
		/* } else if r.URL.Path == "/homepage" {
		homepage(w, r)
		return */
	} else if r.URL.Path == "/new" {
		new(w, r)
		return
	} else if r.URL.Path == "/homeadmin" {
		homeadmin(w, r)
		return
	} else if r.URL.Path == "/home" {
		home(w, r)
		return
	} else if r.URL.Path == "/insert" {
		insert(w, r)
		return
	} else if r.URL.Path == "/delete" {
		delete(w, r)
		return
	} else if r.URL.Path == "/editadmin" {
		editadmin(w, r)
		return
	} else if r.URL.Path == "/updateadmin" {
		updateadmin(w, r)
		return
	} else if r.URL.Path == "/edit" {
		edit(w, r)
		return
	} else if r.URL.Path == "/update" {
		update(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":6699", mux)
	if err != nil {
		log.Fatal("Error running service:", err)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./template/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		users := cekUserPass(r.FormValue("username"), r.FormValue("password"))

		//fmt.Println("Username: ", users.Username)
		if users.Username != "" {
			fmt.Println("Name: ", users.Name)
			if users.Priviledge == "admin" {
				http.Redirect(w, r, "/homeadmin", 301)
			} else {
				http.Redirect(w, r, "/home", 301)
			}
		} else {
			fmt.Println("Username atau password tidak sesuai: ", r.FormValue("username"))
			http.Redirect(w, r, "/", 301)
		}

	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func conn() {
	_, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
}

func cekUserPass(username string, password string) Users {
	var users = Users{}
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	//selDB, err := db.Query("SELECT * FROM users WHERE username=?", username)
	stmt, err := db.Prepare("SELECT * FROM users WHERE username=? and password=?")
	selDB, err := stmt.Query(username, password)
	checkErr(err)
	for selDB.Next() {
		var id int
		var username, password, name, email, priviledge string
		err = selDB.Scan(&id, &username, &password, &name, &email, &priviledge)
		if err != nil {
			panic(err.Error())
		}
		users.Id = id
		users.Username = username
		users.Password = password
		users.Name = name
		users.Email = email
		users.Priviledge = priviledge
	}
	defer db.Close()

	/* 	err = db.QueryRow(`SELECT * FROM users WHERE username=?`, username).
	Scan(
		&users.Id,
		&users.Username,
		&users.Password,
		&users.Name,
		&users.Email,
		&users.Priviledge,
	) */
	return users
}

func cekUserId(idx int) Users {
	var users = Users{}
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	selDB, err := db.Query("SELECT * FROM users WHERE id=?", idx)
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		var id int
		var username, password, name, email, priviledge string
		err = selDB.Scan(&id, &username, &password, &name, &email, &priviledge)
		if err != nil {
			panic(err.Error())
		}
		users.Id = id
		users.Username = username
		users.Password = password
		users.Name = name
		users.Email = email
		users.Priviledge = priviledge
	}
	defer db.Close()
	return users

}

func new(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./template/insert.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
		checkErr(dberr)
		username := r.FormValue("username")
		password := r.FormValue("password")
		name := r.FormValue("name")
		email := r.FormValue("email")
		priviledge := r.FormValue("priviledge")
		stmt, err := db.Prepare("INSERT INTO users(username, password, name, email, priviledge) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		stmt.Exec(username, password, name, email, priviledge)
		log.Println("INSERT: username: " + username + " | priviledge: " + priviledge)
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}

}

func insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./template/insert.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
		checkErr(dberr)
		username := r.FormValue("username")
		password := r.FormValue("password")
		name := r.FormValue("name")
		email := r.FormValue("email")
		priviledge := r.FormValue("priviledge")
		stmt, err := db.Prepare("INSERT INTO users(username, password, name, email, priviledge) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		stmt.Exec(username, password, name, email, priviledge)
		log.Println("INSERT: username: " + username + " | priviledge: " + priviledge)
		defer db.Close()
		http.Redirect(w, r, "/homeadmin", 301)
	}

}

func delete(w http.ResponseWriter, r *http.Request) {
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	usr := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(usr)
	log.Println("DELETED: id: " + usr)
	defer db.Close()
	http.Redirect(w, r, "/homeadmin", 301)
}

func editadmin(w http.ResponseWriter, r *http.Request) {
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	nId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	users := cekUserId(nId)
	t, _ := template.ParseFiles("./template/editadmin.gtpl")
	t.Execute(w, users)
	defer db.Close()
}

func updateadmin(w http.ResponseWriter, r *http.Request) {
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		name := r.FormValue("name")
		email := r.FormValue("email")
		priviledge := r.FormValue("priviledge")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE users SET username=?, password=?, name=?, email=?, priviledge=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(username, password, name, email, priviledge, id)
		log.Println("UPDATE: Username: " + username + " | Priviledge: " + priviledge)

	}
	defer db.Close()
	http.Redirect(w, r, "/homeadmin", 301)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	nId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	users := cekUserId(nId)
	t, _ := template.ParseFiles("./template/edit.gtpl")
	t.Execute(w, users)
	defer db.Close()
}

func update(w http.ResponseWriter, r *http.Request) {
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		name := r.FormValue("name")
		email := r.FormValue("email")
		priviledge := r.FormValue("priviledge")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE users SET username=?, password=?, name=?, email=?, priviledge=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(username, password, name, email, priviledge, id)
		log.Println("UPDATE: Username: " + username + " | Priviledge: " + priviledge)

	}
	defer db.Close()
	http.Redirect(w, r, "/home", 301)
}

func homeadmin(w http.ResponseWriter, r *http.Request) {
	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	selDB, err := db.Query("SELECT * FROM users ORDER BY id DESC")
	checkErr(dberr)
	if err != nil {
		panic(err.Error())
	}
	var users = Users{}
	res := []Users{}
	for selDB.Next() {
		var id int
		var username, password, name, email, priviledge string
		err = selDB.Scan(&id, &username, &password, &name, &email, &priviledge)
		if err != nil {
			panic(err.Error())
		}
		users.Id = id
		users.Username = username
		users.Password = password
		users.Name = name
		users.Email = email
		users.Priviledge = priviledge
		res = append(res, users)
	}
	t, _ := template.ParseFiles("./template/homeadmin.gtpl")
	fmt.Println(res)
	t.Execute(w, res)
	defer db.Close()
}

func home(w http.ResponseWriter, r *http.Request) {

	db, dberr := sql.Open("sqlite3", "./hacktiv8usr.db")
	checkErr(dberr)
	selDB, err := db.Query("SELECT * FROM users ORDER BY id DESC")
	checkErr(dberr)
	if err != nil {
		panic(err.Error())
	}
	var users = Users{}
	res := []Users{}
	for selDB.Next() {
		var id int
		var username, password, name, email, priviledge string
		err = selDB.Scan(&id, &username, &password, &name, &email, &priviledge)
		if err != nil {
			panic(err.Error())
		}
		users.Id = id
		users.Username = username
		users.Password = password
		users.Name = name
		users.Email = email
		users.Priviledge = priviledge
		res = append(res, users)
	}
	t, _ := template.ParseFiles("./template/home.gtpl")
	fmt.Println(res)
	t.Execute(w, res)
	defer db.Close()

}
