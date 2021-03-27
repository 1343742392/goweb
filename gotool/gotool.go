package gotool

import (
	"fmt"
	"reflect"
	"strings"
	"goweb/arrays"
	"goweb/regexpmap"
)

func GetImportStart(code string) int {
	r, err := regexpmap.GetRegexp("import\\(\n")
	if err != nil {
		fmt.Printf(err.Error())
		return -1
	}
	pos := r.FindStringIndex(code)
	if pos == nil {
		return -1
	}
	return pos[1]
}

func DeleteNotes(code string) string {
	fmt.Printf("run")
	r, _ := regexpmap.GetRegexp("\\/\\*[^\\*\\/]*\\*\\/")
	code = r.ReplaceAllString(code, "")

	r2, _ := regexpmap.GetRegexp("(//[^\\n]*\\n)|(//.*)")
	code = r2.ReplaceAllString(code, "")
	return code
}

func GetStructFuncNames(obj interface{}) []string {
	re := reflect.TypeOf(obj)
	ms := re.NumMethod()
	res := []string{}
	for i := 0; i < ms; i++ {
		res = append(res, re.Method(i).Name)
	}
	return res
}

/*
input:
	package gotool

	import(
		...
	)

	....
ouput:
	import(
		...
	)
*/
func GetImport(code string) string {
	r, _ := regexpmap.GetRegexp("import\\s*\\([^\\)]*\\)")
	pos := r.FindStringIndex(code)
	if pos == nil {
		return ""
	}
	return code[pos[0]:pos[1]]
}

func SetImport(code, importStr string) string {
	r, _ := regexpmap.GetRegexp("import\\s*\\([^\\)]*\\)")
	return r.ReplaceAllString(code, importStr)
}

func GetModules(code string) []string {
	simports := GetImport(code)
	//保留括号内的
	importstart := strings.Index(simports, "(")
	simports = simports[importstart+1 : len(simports)]
	importend := strings.LastIndex(simports, ")")
	simports = simports[0:importend]
	//去掉注释
	simports = DeleteNotes(simports)
	//去掉空格和tab
	r, _ := regexpmap.GetRegexp("[ \t\r]")
	simports = r.ReplaceAllString(simports, "")
	//转数组
	imports := strings.Split(simports, "\n")
	//去空
	imports = arrays.DelStrNull(imports)
	return imports
}

/*通过反射的方式来调用参数 可以动态调用可变参数函数
func Format(a ...interface{}) string{
    fmt.Println(a)
    return "format return"
}


ret2 := Call(Format, []interface{}{"hello",44})
*/

func Call(f interface{}, args *[]interface{}) []reflect.Value {
	fun := reflect.ValueOf(f)
	in := make([]reflect.Value, len(*args))
	for k, param := range *args {
		in[k] = reflect.ValueOf(param)
	}
	
	r := fun.Call(in)
	fmt.Printf("%d",  in[0].Elem());
	fmt.Printf("%s",  in[1].Elem());
	return r
}
