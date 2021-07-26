package blog

import (
	"fmt"
)

type Blog struct {
	Id    int64
	Title string
	Body  string
}

func (b Blog) GetURL() string {
	return fmt.Sprintf("https://www.blog.com/blogs/%d", b.Id)
}
