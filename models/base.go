package models

type MsgEnvelope struct {
	Result  bool        `json:"result"`
	ErrMsg  string      `json:"errMsg"`
	ErrCode int64       `json:"errCode"`
	Data    interface{} `json:"data"`
	NowTime string      `json:"nowtime"`
}
