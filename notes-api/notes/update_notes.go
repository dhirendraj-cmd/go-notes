package notes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notes-api/notes-api/models"
	"strconv"
)


func UpdateNotesById(db *sql.DB) http.HandlerFunc{
	fmt.Println("Update notes by id!!")

	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")

		if idStr == ""{
			http.Error(w, "ID is required to pass", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil{
			http.Error(w, "invalid ID", http.StatusBadRequest)
			return 
		}

		var uNotes models.Notes

		// decode data from request
		err1 := json.NewDecoder(r.Body).Decode(&uNotes)

		if err1!=nil{
			http.Error(w, "Error to decode the request", http.StatusBadRequest)
			log.Println(err1)
			return
		}

		sqlStmt := "UPDATE notes SET title=$2, description=$3 WHERE id=$1;"
		res, err := db.Exec(sqlStmt, id, uNotes.Title, uNotes.Description)

		if err!=nil{
			http.Error(w, "Error while updating record!", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		cnt, err := res.RowsAffected()
		if err != nil{
			http.Error(w, "Zero Rows Affected", http.StatusInternalServerError)
			log.Println(err)
			return 
		}

		if cnt == 0 {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Record updated with id: %d", id),
		})

	}
}
