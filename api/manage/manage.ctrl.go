package manage

import (
	"capulus/package/app"
	"capulus/package/w"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetHosts(c *gin.Context) {
	appG := app.Gin{C: c}
	HostService := Host{}
	hosts, err := HostService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, w.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, w.SUCCESS, hosts)
}

type AddHostForm struct {
	Host string `form:"host" valid:"Required;MinSize(15)"`
	Time string `form:"time" valid:"Required;MaxSize(30)"`
}

func AddHost(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddHostForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != w.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	HostService := Host{
		Host: form.Host,
		Time: form.Time,
	}

	err := HostService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, w.SUCCESS, nil)
		return
	}

	appG.Response(http.StatusOK, w.SUCCESS, nil)
}

func DeleteHost(c *gin.Context) {
	var appG = app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Required(id, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, w.INVALID_PARAMS, nil)
	}

	HostService := Host{ID: id}

	if err := HostService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, w.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, w.SUCCESS, nil)
}
