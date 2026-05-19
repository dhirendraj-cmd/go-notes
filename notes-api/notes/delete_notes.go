package notes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// "notes-api/notes-api/models"
)


func DeleteNotesById(db *sql.DB) http.HandlerFunc{
	fmt.Println("Delete notes ")

	return func(w http.ResponseWriter, r *http.Request) {
		id:=r.URL.Query().Get("id")

		// var delNote models.Notes

		delStmt := "DELETE FROM notes WHERE id=$1;"
		res, err := db.Exec(delStmt, id)

		if err!=nil{
			if err == sql.ErrNoRows {
				http.Error(w, "Note not found", http.StatusNotFound)
				return
			}

			http.Error(w, "Db Error!", http.StatusInternalServerError)
			log.Println(err)
			return 
		}

		
		cnt, err := res.RowsAffected()
		
		if err != nil{
			http.Error(w, "Zero Rows Affected", http.StatusInternalServerError)
			log.Println(err)
			return 
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Println("Deleted id: ", id)
		fmt.Printf("%v rows affected\n", cnt)

		
	}

}