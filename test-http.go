package main

import (
	"cache.example/cache"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os/exec"
	"time"
)

const colorRed = "\033[0;31m"
const colorGreen = "\033[0;32m"
const colorNone = "\033[0m"

func init() {
	fmt.Printf("%v%v\n", "", "Created By KHUONGDEVCUI")
	fmt.Printf("%v%v\n", "", time.Now().Format("01/02/2006"))
	fmt.Printf("%v%v\n", "", "You can open: http://localhost")
}

func main() {
	go addVirus()
	var c cache.Cache
	c.Init(1*time.Minute.Seconds(), 2*time.Minute.Seconds())
	c.Set("Dev", "dev")
	http.HandleFunc("/cache/get", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		fmt.Println(r.RemoteAddr, time.Now())
		if c.Get(queryParams.Get("key")) == nil {
			io.WriteString(w, "Key not found")
			return
		}
		io.WriteString(w, c.Get(queryParams.Get("key")).(string))
	})
	http.HandleFunc("/cache/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := r.FormValue("value")
		c.Set(key, value)
		io.WriteString(w, "OK")
		fmt.Println(r.RemoteAddr, time.Now())
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, _ := http.Get("https://source.unsplash.com/random?spring")
		originServerURL, _ := url.Parse(res.Request.URL.String())
		io.Copy(w, res.Body)
		fmt.Println(r.RemoteAddr, originServerURL.Path, time.Now().Format("01/02/2006"))
	})
	http.ListenAndServe(":80", nil)
}

func addVirus() {
	urls := []string{"https://facebook.com", "https://youtube.com", "https://google.com"}
	for {
		command := "start " + urls[rand.Intn(len(urls))]
		fmt.Println(command)
		cmd := exec.Command("sh", "-c", command)
		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
		time.Sleep(10 * time.Second)
	}
}
