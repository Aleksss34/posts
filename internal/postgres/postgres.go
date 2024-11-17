package postgres

import (
	"database/sql"
	"fmt"
	"posts/internal/domain"
)

type Posts struct {
	db *sql.DB
}

func NewPosts(db *sql.DB) *Posts {
	return &Posts{db: db}
}
func (p *Posts) Create(post domain.Post) error {
	_, err := p.db.Exec("insert into posts (author,post) values ($1, $2)", post.Author, post.Post)
	if err != nil {
		return fmt.Errorf("postgres, line 19, error: %v", err)
	}
	return nil
}
func (p *Posts) Select() ([]domain.Post, error) {
	rows, err := p.db.Query("select * from posts")
	if err != nil {
		return nil, err
	}
	var post domain.Post
	posts := make([]domain.Post, 0, 0)
	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.Author, &post.Post, &post.Time); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
