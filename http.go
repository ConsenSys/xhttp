package xhttp

import (
	"fmt"
	"github.com/ConsenSys/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger = log.NewLogrusLogger(log.New("http-server"))

func Start(setters ...Option){
	options := &DefaultOptions
	// apply options
	for _, setter := range setters {
		setter(options)
	}
	// build http.Handler
	var handler http.Handler
	mux := http.NewServeMux()
	handler = mux
	// add routes
	for _, route := range options.Routes {
		logger.Debugln("add route: ", route.Path)
		mux.Handle(route.Path, route.Handler)
	}
	// handle cross origin
	if options.EnableCrossOrigin{
		handler = AccessControl(mux)
	}
	// set http.Handler
	http.Handle("/", handler)
	// start http server
	go StartServer(options.Address, options.Errors)
	go HandleSigInt(options.Errors)

}

func StartServer(httpAddr string, errs chan error){
	logger.Log("transport", "http", "address", httpAddr, "msg", "listening")
	errs <- http.ListenAndServe(httpAddr, nil)
}


func HandleSigInt(errs chan error){
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	errs <- fmt.Errorf("%s", <-c)
}

func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
