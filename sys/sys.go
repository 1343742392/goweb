package sys

import(
	"os/exec"
	"runtime"
)

func Command(windcommand string, liunxcommand string)(string, error){
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "-c", windcommand)
		data, err := cmd.Output()
		if err != nil {
			return "",err;
		}
		return string(data), nil;
	}else{
		cmd := exec.Command("/bin/sh", "-c", liunxcommand)
		data, err := cmd.Output()
		if err != nil {
			return "",err;
		}
		return string(data), nil;
	}
}