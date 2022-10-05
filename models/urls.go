package models

import (
	"github.com/scylladb/gocqlx/v2/table"
	"time"
)

var (
	UrlsTable = table.New(table.Metadata{
		Name: "urls",
		Columns: []string{
			"created_at",
			"short_code",
			"url",
		},
		PartKey: []string{
			"short_code",
		},
		SortKey: []string{},
	})
)

type URL struct {
	Url       string    `db:"url" json:"url"`
	ShortCode string    `db:"short_code" json:"short-code"`
	CreatedAt time.Time `db:"created_at"`
}
