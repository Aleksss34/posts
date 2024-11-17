package service

import (
	"context"
	"posts/internal/config"
	"posts/internal/domain"
	"strings"
)

type repositoryPosts interface {
	Create(post domain.Post) error
	Select() ([]domain.Post, error)
}
type Posts struct {
	repo repositoryPosts
}

func NewPosts(repo repositoryPosts) *Posts {
	return &Posts{repo: repo}
}
func (p *Posts) Create(ctx context.Context, post domain.Post) error {
	sl_post := strings.Split(post.Post, " ")
	var new_sl []string
	for _, post_in_sl := range sl_post {
		for _, mate := range config.MATES {
			if post_in_sl == mate {
				post_in_sl = "###"
				break
			}

		}
		new_sl = append(new_sl, post_in_sl)

	}
	post.Post = strings.Join(new_sl, " ")
	return p.repo.Create(post)
}
func (p *Posts) Select() ([]domain.Post, error) { return p.repo.Select() }
