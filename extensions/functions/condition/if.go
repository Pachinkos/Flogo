package condition

import (
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"fmt"
)

func init() {
	_ = function.Register(&fnIf{})
}

type fnIf struct {
}

func (fnIf) Name() string {
	return "if"
}

func (fnIf) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt, data.TypeString, data.TypeString}, false
}

func (fnIf) Eval(params ...interface{}) (interface{}, error) {
	b, err := coerce.ToBool(params[0])
	
	if err != nil {
		return nil, fmt.Errorf("condition.if function first parameter [%+v] must be a boolean", params[0])
	}

	if b {
		return params[1], nil
	}else{
		return params[2], nil
	}
}
