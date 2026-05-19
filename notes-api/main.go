package main

import (
	"fmt"
	"log"
	"net/http"
	"notes-api/notes-api/database"
	"notes-api/notes-api/notes"
)



func main(){

	fmt.Print("Notes API!!!")

	db := database.Connection()
	defer db.Close()
	

	// handler function creation
	http.HandleFunc("/notes/create", notes.CreateNote(db))
	http.HandleFunc("/notes/all", notes.GetAllNotesofUser(db))

	// get by id
	http.HandleFunc("/notes/get", notes.GetNoteById(db))

	// delete by id
	http.HandleFunc("/notes/delete", notes.DeleteNotesById(db))


	if err := http.ListenAndServe(":8080", nil); err!=nil{
		log.Fatal("Error while starting HTTP Server: ", err)
		fmt.Println("Error while starting HTTP Server: ", err)
	}
}

