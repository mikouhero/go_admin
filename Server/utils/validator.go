package utils

import (
	"errors"
	"fmt"
	"reflect"
)

type Rules map[string]string

type RulesMap map[string]Rules

var CustomizeMap = make(map[string]Rules)

//注册自定义规则方案
func RegisterRule(key string, rule Rules) (err error) {

	if CustomizeMap[key] != nil {
		return errors.New(fmt.Sprintf("%s 已注册无法重复注册"))
	} else {
		CustomizeMap[key] = rule
		return nil
	}
}

//非空 不能为其他类型的 0 值
func NotEmpty() string {
	return "notEmpty"
}

//小于入参 （<） 类型为string  array slice 比较长度  int float uint 比较大小
func Lt(mark string) string {
	return "lt=" + mark
}

// 小于等于入参(<=)
func Le(mark string) string {
	return "le=" + mark
}

// 等于(==)
func Eq(mark string) string {
	return "eq" + mark
}

// 不等于入参(==)
func Ne(mark string) string {
	return "ne=" + mark
}

// 大于等于入参(>=)
func Ge(mark string) string {
	return "ge=" + mark
}

// 大于入参(>)
func Gt(mark string) string {
	return "gt=" + mark
}

//校验方法 接受两个参数 入参实例 规则map
func Verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	//通过反射 获得st  的数据类型
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)

	//kind  st 对应的类型
	kd := val.Kind()

	if kd != reflect.Struct {
		return errors.New("需要结构体struct")
	}

	//返回结构体的字段个数
	num := val.NumField()

	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)

		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				fmt.Println(v)
				//switch {
				//case v == "notEmpty":
				//	if isBlank(val) {
				//		return errors.New(tagVal.Name + "值不能为空")
				//	}
				//case compareMap[strings.Split(v, "=")[0]]:
				//	if !compareVerify(val, v) {
				//		return errors.New(tagVal.Name + "长度或值不在合法范围," + v)
				//	}
				//}
			}
		}
	}
	return nil
}

// 长度和数字的检验方法 根据类型自动转换
func compareVerify(value reflect.Value, verifyStr string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array:
		return compare(value.Len(), verifyStr)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return compare(value.Uint(), verifyStr)
	case reflect.Float32, reflect.Float64:
		return compare(value.Float(), verifyStr)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compare(value.Int(), verifyStr)
	default:
		return false
	}
}

// 非空校验
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uintptr, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		return value.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()

	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func compare(value interface{}, verifyStr string) bool {
	return false
}
