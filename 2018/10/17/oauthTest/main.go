package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

var store = sessions.NewCookieStore([]byte("secret"))

func main() {
	http.HandleFunc("/", RenderMainView)
	http.HandleFunc("/auth", RenderAuthView)
	http.HandleFunc("/auth/callback", Authenticate)

	log.Fatal(http.ListenAndServe(":1333", nil))
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, _ := template.ParseFiles(name)
	tmpl.Execute(w, data)
}

func RenderMainView(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "main.html", nil)
}

func RenderAuthView(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options = &sessions.Options{
		Path:   "/auth",
		MaxAge: 300,
	}
	state := RandToken()
	session.Values["state"] = state
	session.Save(r, w)
	RenderTemplate(w, "auth.html", GetLoginURL(state))
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	state := session.Values["state"]

	delete(session.Values, "state")
	session.Save(r, w)

	if state != r.FormValue("state") {
		http.Error(w, "Invalid session state", http.StatusUnauthorized)
		return
	}

	token, err := OAuthConf.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := OAuthConf.Client(oauth2.NoContext, token)
	userInfoResp, err := client.Get(UserInfoAPIEndpoint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var authUser User
	json.Unmarshal(userInfo, &authUser)

	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400,
	}
	session.Values["user"] = authUser.Email
	session.Values["username"] = authUser.Name
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}

//출처 : https://mingrammer.com/getting-started-with-oauth2-in-go/
