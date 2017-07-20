package handler

import (
	"../network"
	"../node"
	"../token"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strings"
)

func splitAddr(addr string) (string, string) {
	arr := strings.Split(addr, ":")
	ip := arr[0]
	port := arr[1]
	return ip, port
}

func Claim(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lt := ps.ByName("ltoken") //local token
	rt := ps.ByName("rtoken") //remote token
	p := ps.ByName("port")    //port

	log.Println(r.RemoteAddr, "wants to shake", lt, "providing", p, "proposing", rt)
	ip, _ := splitAddr(r.RemoteAddr)

	name := r.FormValue("name")

	n := node.Node{
		Name:        name,
		Ip:          ip,
		LocalToken:  lt,
		RemoteToken: rt,
		Port:        p,
	}

	network.Local.PushNode(n)

	log.Println("Shake performed:", n.LocalToken, "assigned to", n.Name, "("+n.Ip+":"+n.Port+")")
}

func Tap(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := token.Rand()

	network.Local.PushAccessToken(t)

	log.Println("Token given:", t)

	io.WriteString(w, t)
}

func DoTap(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ip := ps.ByName("ip")
	port := ps.ByName("port")

	n := node.Node{
		Ip:   ip,
		Port: port,
	}

	n.Tap()
	log.Println("Remote Token obtained:", n.RemoteToken)

	t := token.Rand()

	log.Println("Trying to claim and give", t)

	n.Claim()

	network.Local.PushNode(n)
}
