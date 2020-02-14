package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

	"capulus/package/w"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, w.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, w.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, w.INVALID_PARAMS
	}

	return http.StatusOK, w.SUCCESS
}
