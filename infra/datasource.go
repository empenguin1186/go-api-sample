package infra

import (
	"database/sql"
	"fmt"
	"log"
)

type DataSource struct {
	Db *sql.DB
	Tx *sql.Tx
}

func NewDataSource(db *sql.DB, tx *sql.Tx) *DataSource {
	return &DataSource{Db: db, Tx: tx}
}

func (d *DataSource) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.Db.Query(query, args)
}

func (d *DataSource) QueryWithTransaction(query string, args ...interface{}) error {
	d.Begin()
	defer d.Commit()

	stmt, err := d.Tx.Prepare(query)
	if err != nil {
		log.Printf("failed to create sql statement. error = %v", err)
		return fmt.Errorf("failed to create sql statement. caused by %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args)
	if err != nil {
		log.Printf("failed to execute sql statement. error = %v", err)
		return fmt.Errorf("failed to execute sql statement. caused by %v", err)
	}
	return nil
}

func (d *DataSource) Begin() (err error) {
	log.Println("transaction begin.")
	if d.Tx != nil {
		log.Println("Tx is not null")
	}
	if d.Tx, err = d.Db.Begin(); err != nil {
		log.Printf("failed to begin transaction. error = %v", err)
		return err
	}
	return nil
}

func (d *DataSource) Commit() error {
	if err := d.Tx.Commit(); err != nil {
		log.Printf("failed to commit transaction. error = %v", err)
		return err
	}
	d.Tx = nil
	log.Println("transaction committed.")
	return nil
}

func (d *DataSource) Rollback() error {
	log.Println("begin rollback.")
	if err := d.Tx.Rollback(); err != nil {
		log.Printf("failed to rollback transation. erorr = %v", err)
		d.Rollback()
	}
	d.Db.Close()
	log.Println("end rollback.")
	return nil
}

func (d *DataSource) Close() {
	if d.Tx != nil {
		log.Printf("transaction is not committed.")
		d.Rollback()
	}
	d.Db.Close()
	log.Println("database closed.")
}
