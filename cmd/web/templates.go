package main

import "github.com/cherradiyacyn/alex-edwards-app/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
