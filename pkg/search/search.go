package search

import (
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
	rows, err := d.db.Query("SELECT id, title, text, tags FROM documents WHERE id = ?", id)
	if err != nil {
		return document, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&document.Id, &document.Title, &document.Text, &document.Tags); err != nil {
			return document, err
		}
	}

	return document, nil
}

func (d *DbEntity) SearchDocuments(query string) ([]Document, error) {
	query = "%" + strings.ToLower(query) + "%"
	rows, err := d.db.Query("SELECT id, title, text, tags FROM documents WHERE title LIKE ? OR text LIKE ? OR tags LIKE ?", query, query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []Document
	for rows.Next() {
		var doc Document
		if err := rows.Scan(&doc.Id, &doc.Title, &doc.Text, &doc.Tags); err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (d *DbEntity) UpdateDocumentTag(document Document) (Document, error) {
	_, err := d.db.Exec("UPDATE documents SET tags = ? WHERE id = ?", document.Tags, document.Id)
	return document, err
}
