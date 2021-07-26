package outbound

import (
	"blog/domain/blog"
	"database/sql"
	"fmt"
	"log"
)

type BlogRepositoryImpl struct {
	db *sql.DB
}

func NewBlogRepositoryImpl(db *sql.DB) *BlogRepositoryImpl {
	return &BlogRepositoryImpl{db: db}
}

func (b BlogRepositoryImpl) Find(id int64) (*blog.Blog, error) {
	found := blog.Blog{}
	err := b.db.QueryRow("select id, title, body from blogs where id = ?", id).Scan(
		&found.Id,
		&found.Title,
		&found.Body)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Printf("db find blog err: %q", err)
		return nil, err
	}
	return &found, nil
}

func (b BlogRepositoryImpl) Save(_ *blog.Blog) error {
	return fmt.Errorf("unimplement")
}
