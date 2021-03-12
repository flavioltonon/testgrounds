package databases

import (
	"database/sql"
	"fmt"
	"testgrounds/configmap"
)

func NewSQL(name string) (*sql.DB, error) {
	switch name {
	default:
		return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			configmap.String("DB_USERNAME"), configmap.String("DB_PASSWORD"), configmap.String("DB_HOST"), configmap.String("DB_NAME")))
	}
}
