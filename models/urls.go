package models

import "github.com/scylladb/gocqlx/v2/table"

var urlsMetadata = table.Metadata{
	Name:    "urls",
	Columns: []string{"short_code", "url", "created_at"},
	PartKey: []string{"short_code"},
	SortKey: []string{"created_at"},
}

var UrlsTable = table.New(urlsMetadata)

type URL struct {
	ShortCode string
	Url       string
	CreatedAt string
}
