package notes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	// "notes-api/notes-api/models"
)


func DeleteNotesById(db *sql.DB) http.HandlerFunc{
	fmt.Println("Delete notes ")

	return func(w http.ResponseWriter, r *http.Request) {
		idStr:=r.URL.Query().Get("id")
		
		if idStr == ""{
			http.Error(w, "id is required", http.StatusBadRequest)
			return 
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}


		sqlStmt := "DELETE FROM notes WHERE id=$1;"
		res, err := db.Exec(sqlStmt, id)

		if err!=nil{

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

		if cnt == 0 {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}
		
		w.WriteHeader(http.StatusNoContent)

		fmt.Println("Deleted id: ", id)
		fmt.Printf("%v rows affected\n", cnt)

		
	}

}