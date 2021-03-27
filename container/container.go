package container

import(

)

var objs  = make(map[string]interface{});

func AddObj(name string, obj interface{}) bool{
	objs[name] = obj;
	return true;
}

func GetObj(name string)(interface{}, bool){
	obj, has := objs[name];
	return obj, has;
}