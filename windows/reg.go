package windows

import(
	registry "golang.org/x/sys/windows/registry"
	"strings"
	"log"
)

/**
windows.GetReg(`HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Environment`, "GOPATH");
*/

func getKey(path string, back func(key registry.Key)(error) )(error)  {
	pos := strings.Index(path, "\\");
	head := path[0: pos];
	tail := path[pos + 1: len(path)];
	var khead = registry.LOCAL_MACHINE;
	switch head{
		case "HKEY_CLASSES_ROOT":
			khead = registry.CLASSES_ROOT;
			break;
		case "HKEY_CURRENT_USER":
			khead = registry.CURRENT_USER;
			break;
		case "HKEY_USERS":
			khead = registry.USERS;
			break;
		case "HKEY_CURRENT_CONFIG":
			khead = registry.CURRENT_CONFIG;
			break;
	}
	k, err := registry.OpenKey(khead, tail, registry.ALL_ACCESS)
	defer k.Close()
	if err != nil {
		log.Fatal(err)
		return err;
	}
    return back(k);
}
func GetReg(path string, key string) (string, error) {
	value := "";
	err := getKey(path, func (k registry.Key) error {
		s, _, err := k.GetStringValue(key)
		if err != nil {
			log.Fatal(err)
			return err;
		}
		value = s;
		return nil;
	});
	if(err != nil){
		return "",err;
	}
	return value, nil;
}

func SetReg(path string, key string, value string)(error)  {
	err := getKey(path, func(k registry.Key)error  {
		k.SetStringValue(key, value);
		return nil;
	});
	if(err != nil){
		return err;
	}
	return nil;
}