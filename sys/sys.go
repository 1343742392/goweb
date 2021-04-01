package sys

import(
	"os/exec"
	"runtime"
)

func Command(command string)(string, error){
	if runtime.GOOS == "windows" {
		cmd := exec.Command("/bin/sh", "-c", command)
		data, err := cmd.Output()
		if err != nil {
			return "",err;
		}
		return string(data), nil;
	}else{
		cmd := exec.Command("/bin/sh", "-c", command)
		data, err := cmd.Output()
		if err != nil {
			return "",err;
		}
		return string(data), nil;
	}
}