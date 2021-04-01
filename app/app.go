package app


import(
	"runtime"
	"strings"
)

type App struct{
	ProjectPath string
}

func New()*App{
	app := new(App);
	_, path, _, _ := runtime.Caller(1) 
	path = strings.ReplaceAll(path, `/`, `\`);
	app.ProjectPath = strings.Replace(path, "\\main\\main.go", "", 1);
	return app;
}
