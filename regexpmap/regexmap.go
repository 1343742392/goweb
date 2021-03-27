package regexpmap

import(
	"regexp"
)
	


var regexps = make(map[string]*regexp.Regexp); 

func GetRegexp(expr string)( *regexp.Regexp, error){
	reg,has := regexps[expr];
	if(has){
		return reg, nil;
	}else{
		newReg ,err := regexp.Compile(expr);
		if(err != nil){
			return nil, err;
		}
		regexps[expr] = newReg;
		return newReg, nil;
	}
}