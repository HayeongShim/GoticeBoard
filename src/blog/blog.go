package blog

type Post struct {
	id      int
	title   string
	content string
}

func NewPost(id int, title string, content string) *Post {
	return &Post{id, title, content}
}

type Blog struct {
	id    int
	url   string
	owner *Blogger
	posts map[int]*Post
}

func NewBlog(id int, url string, owner *Blogger) *Blog {
	blog := Blog{id, url, owner, make(map[int]*Post)}
	owner.blogs[blog.id] = &blog
	return &blog
}

type Blogger struct {
	id       int
	email    string
	password string
	blogs    map[int]*Blog
}

// user := blog.NewBlogger(0, "gogo@naver.com", "1234")
func NewBlogger(id int, email string, password string) *Blogger {
	blogger := Blogger{id, email, password, make(map[int]*Blog)}
	return &blogger
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
