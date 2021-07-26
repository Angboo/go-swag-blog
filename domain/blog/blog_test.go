/*
 * Copyright (c) 2019-2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package blog_test

import (
	"blog/domain/blog"
	"blog/domain/blog/blogmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlog_GetURL(t *testing.T) {
	controller := gomock.NewController(t)
	repository := blogmock.NewMockBlogRepository(controller)
	blog := blog.Blog{
		Id:    0,
		Title: "hello",
		Body:  "world",
	}
	repository.EXPECT().Save(gomock.Eq(&blog))

	b := blog
	// b.Title += " "

	// assert.True(t, &b == &blog)
	assert.Same(t, &blog, &b)
	repository.Save(&b)
}
