package articledto

type ArticleRequest struct {
	UserId      int    `form:"userId"`
	Title       string `form:"title"`
	Attache     string `form:"attache"`
	Description string `form:"description"`
}
