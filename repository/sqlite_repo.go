package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"posts_api/config"
	"posts_api/entity"

	_ "github.com/mattn/go-sqlite3"
)

var (
	dbFile = "posts.db"
	dbPath = config.AbsPathFromProjRoot(dbFile)
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(createNew bool) (*SQLiteRepo, error) {
	if createNew {
		err := os.Remove(dbPath)
		if !errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("failed to remove a file %s: %s", dbPath, err)
		}
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error openning a database: %s", err)
	}

	sqlStmt := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (id integer not null primary key, title text, txt text)`,
		tableName,
	)

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error executing a sql statement; %s", err)
	}

	return &SQLiteRepo{db: db}, nil
}

func (r *SQLiteRepo) Save(post *entity.Post) (*entity.Post, error) {
	res, err := r.db.Exec(fmt.Sprintf("insert into %s (id, title, txt) values (?, ?, ?)", tableName), post.ID, post.Title, post.Text)
	if err != nil {
		return nil, fmt.Errorf("error executing a query: %s", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %s", err)
	}
	return post, nil
}

func (r *SQLiteRepo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	rows, err := r.db.Query(fmt.Sprintf("select * from %s", tableName))
	if err != nil {
		return nil, fmt.Errorf("error selecting all from posts: %s", err)
	}

	for rows.Next() {
		var post entity.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Text); err != nil {
			return nil, fmt.Errorf("error scanning rows: %s", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows: %s", err)
	}
	return posts, nil
}

func (r *SQLiteRepo) Delete(post *entity.Post) (int64, error) {
	res, err := r.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName), post.ID)
	if err != nil {
		return 0, fmt.Errorf("err executing a delete: %s", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("err get last id: %s", err)
	}
	return id, nil
}

func (r *SQLiteRepo) Truncate() error {
	_, err := r.db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	if err != nil {
		return fmt.Errorf("err executing a truncate: %s", err)
	}
	return nil
}

func (r *SQLiteRepo) CloseDB() error {
	if err := r.db.Close(); err != nil {
		return fmt.Errorf("error closing the db connecting %s", err)
	}
	return nil
}
