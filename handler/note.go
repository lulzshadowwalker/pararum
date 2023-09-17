package handler

import (
	"io"
	"net/http"
)

type Note struct{}

func (n *Note) Create(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "note created successfully ✨\n")
}

func (n *Note) Update(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "note updated successfully ✨\n")
}

func (n *Note) Delete(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "note deteted successfully ✨\n")
}

func (n *Note) Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "all notes ✨\n")
}

func (n *Note) Show(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "note of (id) ✨\n")
}
