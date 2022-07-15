package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Events struct {
	Idevent         int
	Name            string
	Date_event      string
	Venue           string
	Description     string
	Idevents_type   int
	Prize_fund      string
	Organizers      string
	Target_audience string
	Link            string
}

func index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:2468586425@tcp(localhost:3306)/practice")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	t, err := template.ParseFiles("static/index.html", "static/promo.html")

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("select * from practice.events")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	products := []Events{}

	for rows.Next() {
		p := Events{}
		err := rows.Scan(&p.Idevent, &p.Name, &p.Date_event, &p.Venue, &p.Description, &p.Idevents_type, &p.Prize_fund, &p.Organizers, &p.Target_audience, &p.Link)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	t.ExecuteTemplate(w, "index", products)
}

func search(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:2468586425@tcp(localhost:3306)/practice")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	t, err := template.ParseFiles("static/search.html","static/promo.html")

	if err != nil {
		panic(err.Error())
	}

	search_text := r.FormValue("search")
	selected_value := r.FormValue("select_box")
	fmt.Println(selected_value)
	if selected_value != "0" {
		rows, err := db.Query(("select * from practice.events where name LIKE ? && idevents_type = ?"), "%"+search_text+"%", selected_value)

		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		products := []Events{}

		for rows.Next() {
			p := Events{}
			err := rows.Scan(&p.Idevent, &p.Name, &p.Date_event, &p.Venue, &p.Description, &p.Idevents_type, &p.Prize_fund, &p.Organizers, &p.Target_audience, &p.Link)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products = append(products, p)

		}
		t.ExecuteTemplate(w, "search", products)
	} else {
		rows, err := db.Query(("select * from practice.events where name LIKE ?"), "%"+search_text+"%")

		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		products := []Events{}

		for rows.Next() {
			p := Events{}
			err := rows.Scan(&p.Idevent, &p.Name, &p.Date_event, &p.Venue, &p.Description, &p.Idevents_type, &p.Prize_fund, &p.Organizers, &p.Target_audience, &p.Link)
			if err != nil {
				fmt.Println(err)
				continue
			}
			products = append(products, p)

		}
		t.ExecuteTemplate(w, "search", products)
	}

}

func handleFunc() {

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)

}

func main() {

	fmt.Println("Подключено к MySQL")
	handleFunc()
}
