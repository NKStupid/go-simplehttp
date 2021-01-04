package main

import (
	"fmt"
	"log"
	"net"
// 	"os"
 	"net/http"
	"io/ioutil"
)

func main() {
	log.Print("simplehttp: Enter main()")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// printing request headers/params
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "\n===> local IP: %q\n\n",GetOutboundIP())
	fmt.Fprintf(w, "Begin to fetch Backend data.")
	MakeRequest(w)
}

func MakeRequest(w http.ResponseWriter) {
	resp, err := http.Get("wise-proto-nlb-df9fa7d3abe4fef3.elb.us-west-2.amazonaws.com")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	str2 := string(body)
	fmt.Fprintf(w, str2)

}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}
