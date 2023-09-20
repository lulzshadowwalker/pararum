package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lulzshadowwalker/pararum/pkg/model"
)

type NoteRepo struct {
	Db *sql.DB
}

func (r *NoteRepo) Create(n model.Note) (err error) {
	_, err = r.Db.Exec(`
	INSERT INTO notes(title, body) VALUES(?, ?);
`, n.Title, n.Body)

	if err != nil {
		return fmt.Errorf("error inserting note into db %w", err)
	}

	return nil
}

func (r *NoteRepo) Update(id int, n model.Note) (err error) {
  result, err := r.Db.Exec(`
		UPDATE notes
		SET
		title = ?, 
		body = ?
		WHERE id = ?;
	`, n.Title, n.Body, id)

  rowsAffected, err := result.RowsAffected()
  if err != nil {
    return fmt.Errorf("error reading rows affected %w", err)
  }

	if errors.Is(err, sql.ErrNoRows) || rowsAffected == 0 {
		return ErrNotFound
	} else if err != nil {
		return fmt.Errorf("error udpating note %w", err)
	}

	return nil
}

func (r *NoteRepo) Delete(id int) (err error) {
	_, err = r.Db.Exec("DELETE FROM notes WHERE id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}

		return fmt.Errorf("error deleting note %w", err)
	}

	return nil
}

func (r *NoteRepo) Index() ([]model.Note, error) {
	rows, err := r.Db.Query("SELECT id, title, body FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var n model.Note
		rows.Scan(&n.Id, &n.Title, &n.Body)

		notes = append(notes, n)
	}

	return notes, err
}

func (r *NoteRepo) Show(id int) (note model.Note, err error) {
	err = r.Db.QueryRow("SELECT id, title, body FROM notes WHERE id = ?", id).Scan(&note.Id, &note.Title, &note.Body)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return note, ErrNotFound
		}

		log.Printf("error fetching note with id (%d) %q", id, err)
		return note, err
	}

	return
}

var ErrNotFound = errors.New("row not found")
