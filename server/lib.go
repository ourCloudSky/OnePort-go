package server

import (
	"encoding/base64"
	"encoding/base32"
	"net"
	"net/http"
	"golang.org/x/net/websocket"
	"net/url"
)

func encode(dat []byte, typ string) string {

	switch typ {
		case "base64":
			return base64.StdEncoding.EncodeToString(dat)
		case "base32":
			return base32.StdEncoding.EncodeToString(dat)
		default:
			panic("Error encoding data: Unknown type: " + typ)
	}

}

func decode(dat string, typ string) []byte {

	switch typ {
		case "base64":
			rslt, err := base64.StdEncoding.DecodeString(typ)
			if err != nil {
				panic(err)
			}
			return rslt
		case "base32":
			rslt, err := base32.StdEncoding.DecodeString(typ)
			if err != nil {
				panic(err)
			}
			return rslt
		default:
			panic("Error decoding data: Unknown type: " + typ)
	}

}

func parseUrl(uri string) (scheme string, host string) {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	return u.Scheme, u.Host
}

func makeConnection(uri string) net.Conn {
	
	scheme, host := parseUrl(uri)

	conn, err := net.Dial(scheme, host)
	
	if err != nil {
		panic(err)
	}
	
	return conn

}

func handleWsRequest(path string, handler func(*websocket.Conn)) {
	http.Handle(path, websocket.Handler(handler))
}