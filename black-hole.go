package blackhole

import (
	"net"
	"net/http"
	"os"
)

type BlackHole struct {
}

func (b *BlackHole) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
	rw.Header().Set("Connection", "close")
	hj, ok := rw.(http.Hijacker)
	if !ok {
		os.Stderr.WriteString("can't hijack rw")
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		os.Stderr.WriteString("can't hijack conn")
		return
	}
	tcpCon, ok := conn.(*net.TCPConn)
	if ok {
		tcpCon.SetLinger(0)
	} else {
		os.Stderr.WriteString("can't set linger")
	}
	if err := conn.Close(); err != nil {
		os.Stderr.WriteString("close conn err")
	}

}
