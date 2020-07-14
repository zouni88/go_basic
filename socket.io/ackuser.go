package main

type AckUserInfo struct {

	Username string `json:"username"`
	NumUsers int `json:"numUsers"`
}

/*func main() {
	ack := AckUserInfo{Username: "cao",NumUsers: 123}
	res ,_ := json.Marshal(ack)
	println(string(res))
}*/