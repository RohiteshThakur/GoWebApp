package main

import (
	//"fmt"
	"net/http"
	//"html/template"
	//"strings"
	//"log"
	//"io/ioutil"
)


func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){		// "/" means root filesystem
		http.ServeFile(w, r, r.URL.Path)
		//http.ServeFile(w, r, "/home/ubuntu/GoWebApp/templates/html/index.html")
	})
	http.ListenAndServe(":9999" , nil)
	// URL: http://127.0.0.1:5050/home/ubuntu/GoWebApp/templates/html/index.html
}



/*
func main (){
	templates := RenderTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil{
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":5000" , nil)
}



func RenderTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result

}
*/

/*
type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(path, ".json"){
			contentType = "application/json"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 My dear - " + http.StatusText(404)))
	}
}

func main() {
    http.Handle("/", new(MyHandler))
    http.ListenAndServe(":5000", nil)
}
*/

/*
func main(){
	//fs := http.FileServer(http.Dir("style"))
	//http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.Handle("/css/", http.FileServer(http.Dir("statics")))
	http.Handle("/js/", http.FileServer(http.Dir("statics")))

	/*
	http.HandleFunc("/templates/statics/js/particles.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/templates/statics/js/particles.json")
	 })

	http.HandleFunc("/templates/statics/css/particle.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/templates/statics/css/particles.css")
	})
	*/


/*
	
	templates := template.Must(template.ParseFiles("../../templates/index.html"))
	http.HandleFunc("/templates/", func(w http.ResponseWriter, r *http.Request){
		if err:= templates.ExecuteTemplate(w, "index.html", nil) ; err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//fmt.Fprintf(w, "Welcome to Gopherland")
	})

	http.ListenAndServe(":5000", nil)
}

*/