package pjson

import (
	"encoding/json"
	"fmt"
)

type JsonObject struct {
	p *interface{}
}

func (jo *JsonObject) getObject(params ...string) interface{} {
	if jo.p == nil {
		return nil
	}

	myObj := *jo.p
	if len(params) == 0 {
		return myObj
	}
	if myMap, ok := myObj.(map[string]interface{}); ok {
		return myMap[params[0]]
	}

	return nil
}

func (jo *JsonObject) GetString(params ...string) string {
	myObj := jo.getObject(params...)
	myStr, ok := myObj.(string)
	if !ok {
		myFloat, ok := myObj.(float64)
		if !ok {
			return ""
		}
		myStr = fmt.Sprint(myFloat)
	}

	return myStr
}

func (jo *JsonObject) GetStringMap(params ...string) map[string]interface{} {
	myObj := jo.getObject(params...)
	myMap, ok := myObj.(map[string]interface{})
	if !ok {
		myMap = make(map[string]interface{})
	}
	return myMap
}

func (jo *JsonObject) GetJsonObject(params ...string) *JsonObject {
	myObj := jo.getObject(params...)
	newJo := &JsonObject{
		&myObj,
	}
	return newJo
}

func (jo *JsonObject) GetJsonObjectSlice(params ...string) []*JsonObject {
	myObj := jo.getObject(params...)
	var newJoSlice []*JsonObject
	if mySlice, ok := myObj.([]interface{}); ok {
		for k := range mySlice {
			newJoSlice = append(newJoSlice, &JsonObject{&mySlice[k]})
		}
	}

	if mySlice, ok := myObj.([]map[string]interface{}); ok {
		for k := range mySlice {
			tmpMap := interface{}(mySlice[k])
			newJoSlice = append(newJoSlice, &JsonObject{&tmpMap})
		}
	}

	return newJoSlice
}

func (jo *JsonObject) Marshal(params ...string) string {
	myObj := jo.getObject(params...)
	bytes, err := json.Marshal(myObj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func NewJsonObject(obj interface{}) *JsonObject {
	newObj := interface{}(nil)
	switch obj.(type) {
	case string:
		_ = json.Unmarshal([]byte(obj.(string)), &newObj)
	case map[string]interface{}:
		newObj = obj
	}
	newJo := &JsonObject{
		&newObj,
	}

	return newJo
}
