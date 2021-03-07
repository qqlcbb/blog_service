package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"go-blog-service/pkg/errcode"
	"net/http"
	"strings"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRow int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page: GetPage(r.Ctx),
			PageSize: GetPage(r.Ctx),
			TotalRows: totalRow,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}

type ValidError struct {
	Key string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string  {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	fmt.Println(v)

	if err != nil {
		value := c.Value("trans")
		translator, _ := value.(ut.Translator)
		errors, ok := err.(val.ValidationErrors)
		if !ok {
			return false, append(errs, &ValidError{
				"",
				err.Error(),
			})
		}
		for key, value := range errors.Translate(translator) {
			errs = append(errs, &ValidError{
				key,
				value,
			})
		}
		return false, errs
	}

	return true, nil
}

