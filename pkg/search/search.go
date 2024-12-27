package search

import (
	"database/sql"
	"strings"
)

func (d *DbEntity) AddDocument(document Document) (Document, error) {
	if res, err := d.db.Exec("INSERT INTO documents (title, text, tags) VALUES (?, ?, ?)", document.Title, document.Text, document.Tags); err != nil {
		return document, err
	} else if document.Id, err = res.LastInsertId(); err != nil {
		return document, err
	} else {
		return document, err
	}
}

func (d *DbEntity) GetDocument(id int64) (document Document, err error) {
	var rows *sql.Rows
	rows, err = d.db.Query("SELECT id, title, text, tags FROM documents WHERE id = ?", id)
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			return
		} else {
			return
		}
	}(rows)

	for rows.Next() {
		if err = rows.Scan(&document.Id, &document.Title, &document.Text, &document.Tags); err != nil {
			return
		}
	}

	return
}

func (d *DbEntity) GetDocumentByTitle(title string) (document Document, err error) {
	var rows *sql.Rows
	rows, err = d.db.Query("SELECT id, title, text, tags FROM documents WHERE title = ?", title)
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			return
		} else {
			return
		}
	}(rows)

	for rows.Next() {
		if err = rows.Scan(&document.Id, &document.Title, &document.Text, &document.Tags); err != nil {
			return
		}
	}

	return
}

func (d *DbEntity) GetDocuments() (documents []Document, err error) {
	var rows *sql.Rows
	rows, err = d.db.Query("SELECT id, title, text, tags FROM documents")
	if err != nil {
		return
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			return
		} else {
			return
		}
	}(rows)

	for rows.Next() {
		var doc Document
		if err := rows.Scan(&doc.Id, &doc.Title, &doc.Text, &doc.Tags); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return
}

func (d *DbEntity) SearchDocuments(query string) (documents []Document, err error) {
	var rows *sql.Rows
	query = "%" + strings.ToLower(query) + "%"
	rows, err = d.db.Query("SELECT id, title, text, tags FROM documents WHERE title LIKE ? OR text LIKE ? OR tags LIKE ?", query, query, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			return
		} else {
			return
		}
	}(rows)

	for rows.Next() {
		var doc Document
		if err = rows.Scan(&doc.Id, &doc.Title, &doc.Text, &doc.Tags); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return
}

func (d *DbEntity) UpdateDocumentTag(document Document) (Document, error) {
	_, err := d.db.Exec("UPDATE documents SET title = ?, tags = ? WHERE id = ?", document.Title, document.Tags, document.Id)
	return document, err
}
