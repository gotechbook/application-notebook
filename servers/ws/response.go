package ws

import "encoding/json"

/************************  响应数据  **************************/
type Head struct {
	Seq      string      `json:"seq"`      // 消息的Id
	Cmd      string      `json:"cmd"`      // 消息的cmd 动作
	Response *WsResponse `json:"response"` // 消息体
}

type WsResponse struct {
	Code    uint32      `json:"code"`
	CodeMsg string      `json:"codeMsg"`
	Data    interface{} `json:"data"` // 数据 json
}

func NewResponse(code uint32, codeMsg string, data interface{}) *WsResponse {
	return &WsResponse{Code: code, CodeMsg: codeMsg, Data: data}
}

// 设置返回消息
func NewResponseHead(seq string, cmd string, code uint32, codeMsg string, data interface{}) *Head {
	response := NewResponse(code, codeMsg, data)
	return &Head{Seq: seq, Cmd: cmd, Response: response}
}

func (h *Head) String() (headStr string) {
	headBytes, _ := json.Marshal(h)
	headStr = string(headBytes)
	return
}
