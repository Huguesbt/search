package search

import (
	"database/sql"
	"fmt"
	"strings"
)

func (d *DbEntity) AddDocument(document Document) (Document, error) {
	if res, err := d.db.Exec("INSERT INTO documents (title, text, tags, description, notes) VALUES (?, ?, ?, ?, ?)", document.Title, document.Text, document.Tags, document.Description, document.Notes); err != nil {
		return document, err
	} else if document.Id, err = res.LastInsertId(); err != nil {
		return document, err
	} else {
		return document, err
	}
}

func (d *DbEntity) GetDocument(id int64) (document Document, err error) {
	var rows *sql.Rows
	rows, err = d.db.Query("SELECT id, title, text, tags, description, notes FROM documents WHERE id = ?", id)
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

	if documents, err := buildDocuments(rows); err != nil {
		return document, err
	} else if len(documents) == 0 {
		return document, nil
	} else {
		return documents[0], err
	}
}

func (d *DbEntity) GetDocumentByTitle(title string) (document Document, err error) {
	var rows *sql.Rows
	rows, err = d.db.Query("SELECT id, title, text, tags, description, notes FROM documents WHERE title = ?", title)
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

	if documents, err := buildDocuments(rows); err != nil {
		return document, err
	} else if len(documents) == 0 {
		return document, nil
	} else {
		return documents[0], err
	}
}

func (d *DbEntity) GetDocuments() (documents []Document, err error) {
	var rows *sql.Rows
	rows, err = d.db.Query("SELECT id, title, text, tags, description, notes FROM documents")
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

	return buildDocuments(rows)
}

func (d *DbEntity) SearchDocuments(query string) (documents []Document, err error) {
	var rows *sql.Rows
	queryFormatted := fmt.Sprintf("%%%s%%", strings.ToLower(query))
	rows, err = d.db.Query("SELECT id, title, text, tags, description, notes FROM documents WHERE title LIKE ? OR text LIKE ? OR tags LIKE ?", queryFormatted, queryFormatted, queryFormatted)
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

	return buildDocuments(rows)
}

func buildDocuments(rows *sql.Rows) (documents []Document, err error) {
	for rows.Next() {
		var doc Document
		if err = rows.Scan(&doc.Id, &doc.Title, &doc.Text, &doc.Tags, &doc.Description, &doc.Notes); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func (d *DbEntity) UpdateDocumentTag(document Document) (Document, error) {
	_, err := d.db.Exec("UPDATE documents SET title = ?, tags = ?, description = ?, notes = ? WHERE id = ?", document.Title, document.Tags, document.Description, document.Notes, document.Id)
	return document, err
}
