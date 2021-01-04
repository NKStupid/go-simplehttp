package main

import (
	"fmt"
	"log"
	"net"
	"os"
 	"net/http"
	"io/ioutil"
)

func main() {
	log.Print("simplehttp: Enter main()")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

// printing request headers/params
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "\n===> local IP: %q\n\n",GetOutboundIP())
	fmt.Fprintf(w, "Begin to fetch Backend data.")
}

func MakeRequest(w http.ResponseWriter) {
	resp, err := http.Get("https://www.packtpub.com/product/the-complete-node-js-developer-course-3rd-edition-video/9781789955071")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, body)

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
