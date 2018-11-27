package web_test

import (
	"testing"

	"github.com/getfider/fider/app/pkg/mock"

	. "github.com/getfider/fider/app/pkg/assert"

	"github.com/getfider/fider/app/pkg/web"
)

func TestMarshalJSON(t *testing.T) {
	RegisterT(t)

	type model struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	bytes, err := web.MarshalJSON(&model{"Jon", 23}, nil)

	Expect(err).IsNil()
	Expect(string(bytes)).Equals(`{"name":"Jon","age":23}`)
}

func TestMarshalJSON_Conditional(t *testing.T) {
	RegisterT(t)

	type model struct {
		Name string `json:"name"`
		Age  int    `json:"age" role:"admin"`
	}

	bytes, err := web.MarshalJSON(&model{"Jon", 23}, mock.AryaStark)

	Expect(err).IsNil()
	Expect(string(bytes)).Equals(`{"name":"Jon"}`)
}
