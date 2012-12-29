package main

import (
	"os"
	"io"
	"io/ioutil"
	"net/http"
)

func gnoteHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(GNOTE_DIR+"/notebook")
	if err != nil {
		io.WriteString(w, "gnote can't open note direct.")
		return
	}
	locals := make(map[string]interface{})
	notes := []string{}
	for _, fileInfo := range fileInfoArr {
		noteSubject := fileInfo.Name()
		noteSubjectLen := len(noteSubject)
		notes = append(notes, noteSubject[:noteSubjectLen])
	}

	locals["notes"] = notes
	renderTemplate(w, GNOTE_DIR+"/templates/note_list.html", locals)
}

func gnoteViewHandler(w http.ResponseWriter, r *http.Request) {
	noteSubject := r.FormValue("note")
	noteContent, err := ioutil.ReadFile(GNOTE_DIR+"/notebook/"+noteSubject)
	if err != nil {
		io.WriteString(w, "gnote can't load note file.")
		return
	}
	locals := make(map[string]interface{})
	locals["subject"] = noteSubject
	locals["content"] = string(noteContent)
	renderTemplate(w, GNOTE_DIR+"/templates/note_view.html", locals)
}

func gnoteAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderTemplate(w, GNOTE_DIR+"/templates/note_add.html", nil)
		return
	}
	if r.Method == "POST" {
		noteSubject := r.FormValue("subject")
		noteContent := r.FormValue("content")
		noteContent2 := []byte(noteContent)
		ioutil.WriteFile(GNOTE_DIR+"/notebook/"+noteSubject, noteContent2, 0)
		http.Redirect(w, r, "/gnote/view?note="+noteSubject, http.StatusFound)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func gnoteDeleteHandler(w http.ResponseWriter, r *http.Request) {
	noteSubject := r.FormValue("note")
	os.Remove(GNOTE_DIR+"/notebook/"+noteSubject)
	http.Redirect(w, r, "/", http.StatusFound)
}

func gnoteChangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		noteSubject := r.FormValue("note")
		noteContent, err := ioutil.ReadFile(GNOTE_DIR+"/notebook/"+noteSubject)
		if err != nil {
			io.WriteString(w, "gnote can't load note file.")
			return
		}
		locals := make(map[string]interface{})
		locals["subject"] = noteSubject
		locals["content"] = string(noteContent)
		renderTemplate(w, GNOTE_DIR+"/templates/note_change.html", locals)
		return
	}
	if r.Method == "POST" {
		noteSubject := r.FormValue("subject")
		noteContent := r.FormValue("content")
		noteContent2 := []byte(noteContent)
		ioutil.WriteFile(GNOTE_DIR+"/notebook/"+noteSubject, noteContent2, 0)
		http.Redirect(w, r, "/gnote/view?note="+noteSubject, http.StatusFound)
		return
	}
}