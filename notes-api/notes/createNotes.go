package notes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notes-api/notes-api/models"
)


func CreateNote(db *sql.DB) http.HandlerFunc {
	fmt.Println("Creating Notes.... ")

	return func(w http.ResponseWriter, r *http.Request) {

		// crete variable to strore data from request body
		var notes models.Notes

		// decode data from request
		err := json.NewDecoder(r.Body).Decode(&notes)

		if err != nil{
			http.Error(w, "Error to decode the request", http.StatusBadRequest)
			log.Println(err)
			return
			// return 
		}

		// execute sql query to insert data
		_, err = db.Exec("INSERT INTO notes (title, description) VALUES ($1, $2)", notes.Title, notes.Description)

		if err != nil{
			http.Error(w, "Error while inserting in notes table.... ", http.StatusInternalServerError)
			log.Println(err)
			return
			// return 
		}

		// sqlStmt := `INSERT INTO notes (title, description) VALUES($1, $2) RETURNING id`
		// id := 0
		// err = db.QueryRow(sqlStmt, notes.Title, notes.Description).Scan(&id)

		// if err != nil{
		// 	log.Println(err)
		// return
		// 	// return 
		// }
		
		// fmt.Println("New record ID is:", id)
		
		// response msg with created 201 status
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Note Created in DB"))

	}


}