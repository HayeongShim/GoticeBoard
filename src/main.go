package main

import (
	"blog"
	"fmt"
)

func main() {
	user := blog.NewBlogger(0, "gogo@naver.com", "1234")
	blog := blog.NewBlog(0, "https://gogo.tistory.com", user)
	post := blog.NewPost(0, "What is Lorem Ipsum?", "Lorem Ipsum is simply dummy text of ...")

	/*posts := []*Post{
		blog.NewPost(0, "What is Lorem Ipsum?", "Lorem Ipsum is simply dummy text of ..."),
		blog.NewPost(1, "Why do we use it", "It is a long established fact that ..."),
	}*/

	//for _, post := range posts {
	user.Create(blog, post)
	//}

	fmt.Println(user.Read(blog, 0), user.Read(blog, 1))
}
