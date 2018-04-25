package res

import "encoding/json"

type proxy struct {
	ID   int    `json:"id"`   // 代理编号
	Host string `json:"host"` // 代理 ip:port
}

func NewProxy(id int, host string) *proxy {
	return &proxy{
		ID:   id,
		Host: host,
	}
}
func (u proxy) GetID() int {
	return u.ID
}

func (u proxy) Info() []byte {
	info, _ := json.Marshal(u)
	return info
}

func (u *proxy) Type() Type {
	return Proxy
}
