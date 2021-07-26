package usecase

import (
	"blog/domain/blog"
	"log"
)

type BlogUseCase struct {
	repo            blog.BlogRepository
	shortURLGateway ShortURLGateway
}

func NewBlogUseCase(repo blog.BlogRepository, shortURLGateway ShortURLGateway) *BlogUseCase {
	return &BlogUseCase{repo: repo, shortURLGateway: shortURLGateway}
}

func (c BlogUseCase) FindBlog(id int64) (*blog.Blog, error) {
	found, err := c.repo.Find(id)
	if err != nil {
		log.Printf("err: %s", err)
		return nil, err
	}
	return found, err
}

func (c BlogUseCase) ShareBlog(id int64) (string, error) {
	found, err := c.repo.Find(id)
	if err != nil {
		log.Printf("err: %s", err)
		return "", err
	}
	if found == nil {
		return "", nil
	}

	shortcut, err := c.shortURLGateway.CreateShortURL(found.GetURL())
	if err != nil {
		log.Printf("short URL gateway err: %s", err)
		return "", err
	}
	return shortcut, nil
}
