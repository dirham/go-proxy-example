package entities

import (
	"time"
)

type Proxy struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`         // name of apis
	RequestPath string    `json:"request_path"` // request path for example /test point to upstream url
	RequestHost string    `json:"request_host"` // host header to send to upstram url
	UpstreamUrl string    `json:"upstream_url"` // target url to proxying
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}
