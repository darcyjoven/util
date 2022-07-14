package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewDB(db, connect string) (*DB, error) {
	d, err := sqlx.Connect(db, connect)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &DB{
		db: d,
	}, nil
}

func (db *DB) Close() {
	db.db.Close()
}

func (db *DB) GetData(query string, args ...any) [][]any {
	data := make([][]any, 0, 64)
	rows, err := db.db.Queryx(query, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		temp, err := rows.SliceScan()
		if err != nil {
			log.Println(err)
			return nil
		}
		data = append(data, temp)
	}
	return data
}

func (db *DB) Fetch(query string, args ...any) []any {

	row := db.db.QueryRowx(query, args...)
	temp, err := row.SliceScan()
	if err != nil {
		log.Println(err)
		return nil
	}
	return temp
}

func (db *DB) Foreach(query string, args ...any) *Rows {
	rows, err := db.db.Queryx(query, args...)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &Rows{
		rows: rows,
	}
}

func (db *DB) Exec(execute string, args ...any) (int64, error) {
	result, err := db.db.Exec(execute, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	rowaffect, _ := result.RowsAffected()
	return rowaffect, err
}

func (r *Rows) Next() bool {

	return r.rows.Next()
}

func (r *Rows) SliceScan() ([]interface{}, error) {
	return r.rows.SliceScan()
}
func (r *Rows) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}
