package errors

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

const (
	ErrorReasonArticleNotFound = "ARTICLE_NOT_FOUND"
	ErrorReasonArticleInvalid  = "ARTICLE_INVALID"
	ErrorReasonEmptyTag        = "EMPTY_TAG"
	ErrorReasonInvalidId       = "INVALID_ID"
	ErrorReasonEmptyTitle      = "EMPTY_TITLE"
	ErrorReasonEmptySummary    = "EMPTY_SUMMARY"
	ErrorReasonEmptyContent    = "EMPTY_CONTENT_URL"
	ErrorReasonArticleNil      = "ARTICLE_NIL"
	ErrorReasonArticleID       = "ARTICLE_ID_INVALID"
)

var (
	ErrorArticleNotFound     = errors.NotFound(ErrorReasonArticleNotFound, "文章不存在")
	ErrorArticleInvalid      = errors.BadRequest("ARTICLE_INVALID", "文章格式错误")
	ErrorArticleEmptyTag     = errors.BadRequest(ErrorReasonEmptyTag, "标签不能为空")
	ErrorArticleInvalidId    = errors.BadRequest(ErrorReasonInvalidId, "无效的文章ID")
	ErrorArticleEmptyTitle   = errors.BadRequest(ErrorReasonEmptyTitle, "文章标题不能为空")
	ErrorArticleEmptySummary = errors.BadRequest(ErrorReasonEmptySummary, "文章摘要不能为空")
	ErrorArticleEmptyContent = errors.BadRequest(ErrorReasonEmptyContent, "文章内容URL不能为空")
	ErrorArticleNil          = errors.BadRequest(ErrorReasonArticleNil, "文章对象不能为空")
	ErrorArticleID           = errors.BadRequest(ErrorReasonArticleID, "ID错误!")
)

// WrapArticleError 包装文章相关错误
func WrapArticleError(err error, articleId ...int64) error {
	if err == nil {
		return ErrorArticleNil
	}

	metadata := map[string]string{}
	if len(articleId) > 0 {
		metadata["article_id"] = fmt.Sprintf("%d", articleId[0])
	}

	return errors.FromError(err).WithMetadata(metadata)
}
