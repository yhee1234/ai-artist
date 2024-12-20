package main

import (
	"ai-artist/gateway/handler"
	"ai-artist/gateway/setting"
	usermanager "ai-artist/gateway/userManager"
	"ai-artist/gateway/utils/corsController"
	"ai-artist/gateway/utils/logging"
	"net/http"

	"github.com/urfave/negroni"
)

const key = "key.pem"
const cert = "cert.pem"

func initiallization() {
	setting.Init()
	usermanager.Init()
	logging.Init()
}

func startServerHTTPS() {
	mux := handler.CreateHandler()
	handler := negroni.Classic()
	defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE, OPTIONS", "*", true))
	handler.UseHandler(mux)

	logging.Logger.Info("HTTPS server start.")
	http.ListenAndServeTLS(":"+setting.Setting.ServerPort, cert, key, handler)
}

func startServerHTTP() {
	mux := handler.CreateHandler()
	handler := negroni.Classic()
	defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE, OPTIONS", "*", true))
	handler.UseHandler(mux)

	logging.Logger.Info("HTTP server start.")
	http.ListenAndServe(":"+setting.Setting.ServerPort, handler)
}

func main() {
	initiallization()

	startServerHTTPS()
	//startServerHTTP()
}
