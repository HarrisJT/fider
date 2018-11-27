package web

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/getfider/fider/app/models"
)

// MarshalJSON returns a JSON object from given input
func MarshalJSON(i interface{}, user *models.User) ([]byte, error) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		requiredRoles := strings.Split(field.Tag.Get("role"), ",")
		if len(requiredRoles) > 0 {
			panic(v.FieldByName(field.Name))
			v.FieldByName(field.Name).SetInt(0)
		}
	}
	return json.Marshal(i)
}
