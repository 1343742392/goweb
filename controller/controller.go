package controller

import(
	"web/filetool"
	"fmt"
)

type IC interface {
	Init(name, projectPath string)
}


type Controller struct {
	ProjectPath string
	moduleName string
}


func (c *Controller) Init(name, projectPath string){
	c.moduleName = name;
	c.ProjectPath = projectPath;
}


func (c *Controller) View(name string)([]byte){
    data,err := filetool.GetFileBytes(c.ProjectPath + "/src/modules/" + c.moduleName + "/view/" + name + ".html")
	if(err != nil){
		fmt.Printf("view file read fail");
		return []byte{};
	}
	return data;
}