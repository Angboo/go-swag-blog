package usecase

import (
	"blog/app/usecase/gatewaymock"
	"blog/domain/blog"
	"blog/domain/blog/blogmock"
	"blog/util/vxlog"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlogUseCase_ShouldReturnShortURLWhenShareBlog(t *testing.T) {
	const id int64 = 1
	foundBlog := blog.Blog{Id: id, Title: "a title", Body: "a body."}
	const expectedShortURL = "https://s.url/xcndiejdd"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := blogmock.NewMockBlogRepository(ctrl)
	repo.EXPECT().Find(id).Return(&foundBlog, nil)
	shortURLGateway := gatewaymock.NewMockShortURLGateway(ctrl)
	shortURLGateway.EXPECT().CreateShortURL(foundBlog.GetURL()).Return(expectedShortURL, nil)
	uc := NewBlogUseCase(repo, shortURLGateway)

	shortURL, err := uc.ShareBlog(id)

	assertions := assert.New(t)
	assertions.NoError(err)
	assertions.Equal(expectedShortURL, shortURL)
}

func TestCaller(t *testing.T) {
	vxlog.Printf("hello %s!", "world")
}
