package pjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("unexpected value obtained; got %q; want %q", a, b)
	}
}

func TestGetNonLeaf(t *testing.T) {
	jsonContent := `{
					"User": {
						"name": "Portobello McSmith",
						"email": "p@mcsmith.com",
						"passwd": "mushrooms"
					},
					"Users": [
								{
									"name": "Portobello McSmith",
									"email": "p@mcsmith.com",
									"password": "mushrooms"
								},
								{
									"name": "Cucumber Daniels",
									"email": "c@daniels.com",
									"password": "crunchy"
								},
								{
									"name": "Luke Shaw",
									"email": "l@shaw.com",
									"password": "apple"
								}
					]
					}`

	jo := NewJsonObject(jsonContent)
	user1, _ := json.Marshal(jo.GetJsonObject("User").GetStringMap())
	user2, _ := json.Marshal(map[string]interface{}{
		"name":   "Portobello McSmith",
		"email":  "p@mcsmith.com",
		"passwd": "mushrooms",
	})
	assertEqual(t, string(user1), string(user2))

	users := jo.GetJsonObjectSlice("Users")
	passwords := []string{"mushrooms", "crunchy", "apple"}
	for i, item := range users {
		assertEqual(t, item.GetString("password"), passwords[i])
	}
}

func TestMarshal(t *testing.T) {
	jsonContent := `{"a":"123", "b":[{"b1":"456"}]}`
	jo := NewJsonObject(jsonContent)
	fmt.Println("jo Marshal: ", jo.Marshal())
	assertEqual(t, jo.Marshal(), `{"a":"123","b":[{"b1":"456"}]}`)
}
