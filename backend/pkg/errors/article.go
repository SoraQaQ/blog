package errors

import "github.com/go-kratos/kratos/v2/errors"

const (
	ErrorReasonArticleNotFound = "ARTICLE_NOT_FOUND"
	ErrorReasonArticleInvalid  = "ARTICLE_INVALID"
)

var (
	ErrorArticleNotFound = errors.NotFound(ErrorReasonArticleNotFound, "文章不存在")
	ErrorArticleInvalid  = errors.BadRequest("ARTICLE_INVALID", "文章格式错误")
)
