//go:generate mockgen -package blogmock -destination blogmock/blog_repository.go . BlogRepository
package blog

type BlogRepository interface {
	Find(id int64) (*Blog, error)
	Save(blog *Blog) error
}
