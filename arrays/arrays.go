package arrays


import (
    "reflect"
    //"fmt"
    //"strconv"   
)

var Test = "";
// Contains Returns the index position of the val in array
func Contains(array interface{}, val interface{}) (index int) {
    index = -1
    s := reflect.ValueOf(array)
    for i := 0; i < s.Len(); i++ {
        if reflect.DeepEqual(val, s.Index(i).Interface()) {
            index = i
            return
        }
    }

    return
}

func DelStrNull(array []string)[]string{
	res := []string{};
	for i := 0; i < len(array); i++ {
        if(array[i] != ""){
			res = append(res, array[i]);
        }
	}
	return res;
}