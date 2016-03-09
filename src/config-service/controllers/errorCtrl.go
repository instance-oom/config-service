package controllers

//ErrorController : handle error like 404 500 etc.
type ErrorController struct {
	BaseController
}

//Error404 : handle 404 error
func (e *ErrorController) Error404() {
	e.Data["json"] = map[string]string{
		"code":  "404",
		"error": "page not found.",
	}
	e.Ctx.Output.SetStatus(404)
	e.ServeJSON()
	e.StopRun()
}
