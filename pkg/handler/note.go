package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lulzshadowwalker/pararum/pkg/model"
	"github.com/lulzshadowwalker/pararum/pkg/repo"
)

type Note struct {
	Repo repo.NoteRepo
}

func (n *Note) Create(w http.ResponseWriter, r *http.Request) {
	title, body := r.PostFormValue("title"), r.PostFormValue("body")
	var validationMessages []string
	if title == "" {
		validationMessages = append(validationMessages, "title parameter is required")
	}
	if body == "" {
		validationMessages = append(validationMessages, "body parameter is required")
	}
	if len(validationMessages) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors(validationMessages...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(res))
		return
	}

	note := model.Note{
		Title: title,
		Body:  body,
	}

	if err := n.Repo.Create(note); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (n *Note) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors("path parameter [id] is required")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(res))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors("path parameter [id] has to be an integer")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(res))
		return
	}

	title, body := r.PostFormValue("title"), r.PostFormValue("body")
	var validationMessages []string
	if title == "" {
		validationMessages = append(validationMessages, "title parameter is required")
	}
	if body == "" {
		validationMessages = append(validationMessages, "body parameter is required")
	}
	if len(validationMessages) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors(validationMessages...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(res))
		return
	}

	note := model.Note{
		Title: title,
		Body:  body,
	}

	if err := n.Repo.Update(id, note); err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (n *Note) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors("path parameter [id] is required")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(res))
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors("path parameter [id] has to be an integer")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(res))
		return
	}

	if err := n.Repo.Delete(id); err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return

		}

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Printf("error encoding json %q", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (n *Note) Index(w http.ResponseWriter, r *http.Request) {
	notes, err := n.Repo.Index()
	if err != nil {
		log.Printf("error reading from db %q", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	res := struct {
		Status bool         `json:"status"`
		Data   []model.Note `json:"data"`
	}{
		true,
		notes,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Printf("error encoding json %q", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func (n *Note) Show(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors("path parameter [id] is required")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(res))
		w.WriteHeader(http.StatusBadRequest)
		json, err := json.Marshal(res)
		if err != nil {
			log.Printf("error encoding json %q", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(json)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, err := validationErrors("query parameter [id] has to be an integer")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(res))
		w.WriteHeader(http.StatusBadRequest)
		json, err := json.Marshal(res)
		if err != nil {
			log.Printf("error encoding json %q", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(json)
		return
	}

	note, err := n.Repo.Show(id)
	if err != nil {
    if errors.Is(err, repo.ErrNotFound) {
      w.WriteHeader(http.StatusNotFound)
    }

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(note)
	if err != nil {
		log.Printf("error encoding json %q", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(json)
}

func validationErrors(messages ...string) (string, error) {
	res := struct {
		Status   bool     `json:"status"`
		Messages []string `json:"messages"`
	}{
		false,
		messages,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Printf("error encoding json %q", err)
		return "", err
	}

	return string(json), nil
}
