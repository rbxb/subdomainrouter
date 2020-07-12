package main

import (
	"flag"
	"net/http"
)

var path string
var port string

type router struct{}

func init() {
	flag.StringVar(&port, "port", ":80", "The address and port the fileserver listens at.")
	flag.StringVar(&path, "path", "./subdomains.txt", "The filepath of the configuration file.")
}

func main() {
	flag.Parse()
	http.ListenAndServe(port, &router{})
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world!"))
}

/*func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if server := getServer(req); server != "" {
		proxy(w, req, server)
	} else {
		http.Error(w, "Not found.", 404)
	}
}

func getServer(req *http.Request) string {
	host := strings.Split(req.URL.Host, ".")
	subdomain := ""
	if len(host) > 2 {
		subdomain = host[0]
	}
	log.Println(req.RemoteAddr, " ", subdomain)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	s := string(b)
	lines := strings.Split(s, "\n")
	server := ""
	for _, line := range lines {
		split := strings.Split(line, " ")
		if subdomain == "" && len(split) == 1 {
			server = split[0]
			break
		} else if split[0] == subdomain {
			server = split[1]
			break
		}
	}
	return strings.TrimSpace(server)
}

func proxy(w http.ResponseWriter, req *http.Request, server string) {
	u, err := url.Parse(server)
	if err != nil {
		panic(errors.New("couldn't parse URL '" + server + "'"))
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		http.Error(w, "Internal error.", 500)
	}
	proxy.ServeHTTP(w, req)
}
*/
