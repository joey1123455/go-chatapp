package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// handles the HTTP requests
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "The address of the application.")
	flag.Parse() //parse the flags
	//setup gomniauth
	gomniauth.SetSecurityKey("tyu567hreuwyshfbhgywuudhwuuy3yhy74t65dtysghghduyiye7t636teyghsfdyrt7ewy7ye87266rygg7t27teu")
	gomniauth.WithProviders(
		github.New("80eca123afcb2a6d87ee", "8f22d5801e542c43395a818125e95597cb323121", "http://127.0.0.1:8080/auth/callback/github"),
		google.New("842680277389-ln21cp0oga1g9t0o1s1g2tq0mivg7vi8.apps.googleusercontent.com", "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", "http://127.0.0.1:8080/auth/callback/google"),
	)
	// creates the room
	r := newRoom(UseGravater)

	//routes
	//templateHandler renders the front end template on the routes
	//MustAuth wrappes a authentication checker and redirect on the route
	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	// provides the websocket connection for chat rooms
	http.Handle("/room", r)
	//user login template link
	http.Handle("/login", &templateHandler{filename: "login.html"})
	//handles oauth client requests
	http.HandleFunc("/auth/", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	//user avater upload template link
	http.Handle("/upload", MustAuth(&templateHandler{filename: "upload.html"}))
	//handles avater upload
	http.HandleFunc("/uploader", uploadHandler)

	//initiate the room
	go r.run()
	//start the web server
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Listen and Serve:", err)
	}
}
