package main

import (
	"fmt"
	"leetcode/pjson"
	"regexp"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func UnknowJson(pattern string) {
	//tj1 := `{"name": "qiqi", "password": "mali1"}`
	//tj2 := `[{"name": "qiqi", "password": "mali1"},{"name": "qiqi2", "password": "mali12"}]`
	tj3 := `{
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

	//jo1 := pjson.NewJsonObject(tj1)
	//jo2 := pjson.NewJsonObject(tj2)
	jo3 := pjson.NewJsonObject(tj3)

	joObj := jo3.GetJsonObject("User").GetStringMap()
	if len(joObj) != 0 {
		for k, _ := range joObj {
			if ok, _ := regexp.MatchString(pattern, k); ok {
				joObj[k] = "REDACTED"
			}
		}
		//fmt.Println(jo1.Marshal())
	}

	joArr := jo3.GetJsonObjectSlice("Users")
	for _, joRow := range joArr {
		joValue := joRow.GetStringMap()
		for k, _ := range joValue {
			if ok, _ := regexp.MatchString(pattern, k); ok {
				joValue[k] = "REDACTED"
			}
		}
	}

	fmt.Println("result: ", jo3.Marshal())
}

func main() {
	//coins := []int{1, 2, 5, 10}
	//amount := 98
	//result := coinChange(coins, amount)
	UnknowJson("passw(d|ord)")

}
