package repository

import (
	"database/sql"
	"entity"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbFile = "./posts.db"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo() (*SQLiteRepo, error) {
	os.Remove(dbFile)
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("error openning a database: %s", err)
	}
	defer db.Close()

	sqlStmt := fmt.Sprintf(`
	create table %s (id integer not null primary key, title text, txt text);
	delete from %s;
	`, tableName, tableName)

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error executing a sql statement; %s", err)
	}

	return &SQLiteRepo{db: db}, nil
}

func (s *SQLiteRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("error openning a database: %s", err)
	}
	defer db.Close()

	res, err := db.Exec(fmt.Sprintf("insert into %s (id, title, txt) values (?, ?, ?)", tableName), post.ID, post.Title, post.Text)
	if err != nil {
		return nil, fmt.Errorf("error executing a query: %s", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %s", err)
	}
	return post, nil
}

func (s *SQLiteRepo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("error openning a database: %s", err)
	}
	defer db.Close()

	var posts []entity.Post
	rows, err := db.Query(fmt.Sprintf("select * from %s", tableName))
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
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return 0, fmt.Errorf("error openning a database: %s", err)
	}
	defer db.Close()

	res, err := db.Exec(fmt.Sprintf("delete from %s where id = ?", tableName), post.ID)
	if err != nil {
		return 0, fmt.Errorf("err executing a delete: %s", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("err get last id: %s", err)
	}
	return id, nil
}
