package notes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notes-api/notes-api/models"
)

func GetAllNotesofUser(db *sql.DB) http.HandlerFunc {
	fmt.Println("Get ALL notes.....")

	return func(w http.ResponseWriter, r *http.Request)  {

		
		// get all notes
		rows, err := db.Query("SELECT * FROM notes")
		
		if err !=  nil{
			http.Error(w, "Error occured while getting all data from db", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		
		
		defer rows.Close()
		
		var allNotes []models.Notes

		// iterate over rows and store each note in allnotes
		for rows.Next(){
			var singleNote models.Notes
			
			// scan the rows
			if err := rows.Scan(&singleNote.ID, &singleNote.Title, &singleNote.Description); err != nil{
				http.Error(w, "Erorr scanning notes!!!", http.StatusInternalServerError)
				log.Println(err)
				return
			}

			// append eact note to array of allnotes
			allNotes = append(allNotes, singleNote)
		}

		// check if error occurs while iterating over rows
		if err := rows.Err(); err != nil{
			http.Error(w, "Error reading rows in Database", http.StatusInternalServerError)
			log.Println(err)
			return
		}


		// set content-type header to appication/json
		w.Header().Set("Content-Type", "application/json")

		// set status code to 200
		w.WriteHeader(http.StatusOK)

		// finally encode the notes to json and write it to response body
		if err := json.NewEncoder(w).Encode(allNotes); err!=nil{
			http.Error(w, "Error enoding notes", http.StatusInternalServerError)
			log.Println(err)
			return
		}

	}

}


// localhost:8080/notes/create
// localhost:8080/notes/all