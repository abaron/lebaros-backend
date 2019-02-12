package common

type Response struct {
	Code     int16       `json:"code"`
	Status   string      `json:"status"`
	Messages string      `json:"messages"`
	Data     interface{} `json:"data"`
}
