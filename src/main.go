package main

import {
	"blog"
	"fmt"
}

type Post struct {
	id int
	title string
	content string
}

func NewPost (id int, title string, content string) *Post {
	return &Post{id, title, content}
}

type Blog struct {
	id int
	url string
	owner *Blogger
	posts map[int]*Post
}

func NewBlog (id int, url string, owner *Blogger) *Blog {
	blog := Blog{id, url, owner, make(map[int]*Post)}
	owner.blogs[blog.id] = &blog
	return &blog
}

type Blogger struct {
	id int
	id string
	password string
	blogs map[int]*Blog
}

func (user *Blogger) Create(blog *Blog, post *Post) {
	if _, ok := user.blogs[blog.owner.id]; ok {
		blog.posts[post.id] = post
	}
}

func (user *Blogger) Read(blog *Blog, id int) *Post {
	return blog.posts[id]
}

func (user *Blogger) Update(blog *Blog, id int, title string, content string) {
	if _, ok := user.blogs[blog.owner.id]; ok {
		post := blog.posts[id]
		post.title, post.content = title, content
	}
}

func (user *Blogger) Delete(blog *Blog, id int) {
	if _, ok := user.blogs[blog.owner.id]; ok {
		delete(blog.posts, id)
	}
}

func main() {
	user := blog.NewBlogger(0, "gogo@naver.com", "1234")
	blog := blog.New(0, "https://gogo.tistory.com", user)

	posts := []*Post {
		blog.NewPost(0, "What is Lorem Ipsum?", "Lorem Ipsum is simply dummy text of ...");
		blog.NewPost(1, "Why do we use it", "It is a long established fact that ...");
	}

	for _, post := range posts {
		user.Create(blog, post)
	}

	fmt.Println(user.Read(blog, 0), user.Read(blog, 1))
}
