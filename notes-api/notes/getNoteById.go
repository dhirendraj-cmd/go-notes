package notes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notes-api/notes-api/models"
)


func GetNoteById(db *sql.DB) http.HandlerFunc {
	fmt.Println("Get ALL notes.....")

	return func(w http.ResponseWriter, r *http.Request)  {

		id := r.URL.Query().Get("id")

		var note  models.Notes

		err := db.QueryRow("SELECT id, title, description FROM notes WHERE id=$1", id,).Scan(&note.ID, &note.Title, &note.Description)

		if err !=  nil{
			if err == sql.ErrNoRows {
				http.Error(w, "Note not found", http.StatusNotFound)
				return
			}

			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Println(err)
			return

		}


		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err:=json.NewEncoder(w).Encode(note); err!=nil{
			http.Error(w, "Fatal Error encoding note", http.StatusInternalServerError)
			log.Println(err)
			return
		}



	}
}

