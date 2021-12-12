package interface_utils

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

type IInterfaceUtils interface {
	
}

type InterfaceUtils struct {
	
}


func NewInterfaceUtils() *InterfaceUtils {
	
	return &InterfaceUtils{}
}

func GetReturnedDataFromInterface(innInterfaceData interface{}) (outInterfaceData interface{}) {
	switch v := innInterfaceData.(type) {

	case *gin.Context:
		// v is an int here, so e.g. v + 1 is possible.
		message := fmt.Sprintf("Detected *gin.Context: %v", v)
		outInterfaceData = innInterfaceData.(*gin.Context)
		log.Println(message)
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		message := fmt.Sprintf("Detected Integer: %v", v)
		outInterfaceData = innInterfaceData.(int)
		log.Println(message)
	case int16:
		// v is an int here, so e.g. v + 1 is possible.
		message := fmt.Sprintf("Detected Integer base 16: %v", v)
		outInterfaceData = innInterfaceData.(int16)
		log.Println(message)
	case int32:
		// v is an int here, so e.g. v + 1 is possible.
		message := fmt.Sprintf("Detected Integer base 32: %v", v)
		outInterfaceData = innInterfaceData.(int32)
		log.Println(message)

	case int64:
		// v is an int here, so e.g. v + 1 is possible.
		message := fmt.Sprintf("Detected Integer base 64: %v", v)
		outInterfaceData = innInterfaceData.(int64)
		log.Println(message)
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		message := fmt.Sprintf("Detected Float64: %v", v)
		log.Println(message)
		outInterfaceData = innInterfaceData.(float64)

	case float32:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		message := fmt.Sprintf("Detected Float64: %v", v)
		log.Println(message)
		outInterfaceData = innInterfaceData.(float32)

	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		fmt.Printf("Detected String: %v", v)
		outInterfaceData = innInterfaceData.(string)

	default:
		// And here I'm feeling dumb. ;)
		message := fmt.Sprintf("I don't know, ask stackoverflow.Or Add Support to the interface type you supplied %v", reflect.TypeOf(innInterfaceData))
		log.Println(message)

	}

	return outInterfaceData

}


// func (InterfaceUtils) ConvertValue(v interface{}) (driver.Value, error) {
// 	logger := getLogger()
// 	if logger != nil {
// 		logger.Log("ConvertValue", "Float64", "value", v)
// 	}
// 	switch x := v.(type) {
// 	case int8:
// 		return float64(x), nil
// 	case int16:
// 		return float64(x), nil
// 	case int32:
// 		return float64(x), nil
// 	case uint16:
// 		return float64(x), nil
// 	case uint32:
// 		return float64(x), nil
// 	case int64:
// 		return float64(x), nil
// 	case uint64:
// 		return float64(x), nil
// 	case float32:
// 		return float64(x), nil
// 	case float64:
// 		return x, nil
// 	case string:
// 		if x == "" {
// 			return 0, nil
// 		}
// 		return strconv.ParseFloat(x, 64)
// 	case Number:
// 		if x == "" {
// 			return 0, nil
// 		}
// 		return strconv.ParseFloat(string(x), 64)
// 	case *Number:
// 		if x == nil || *x == "" {
// 			return 0, nil
// 		}
// 		return strconv.ParseFloat(string(*x), 64)
// 	default:
// 		return nil, fmt.Errorf("unknown type %T", v)
// 	}
// }
