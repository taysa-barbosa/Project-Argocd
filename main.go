package main
import "net/http"

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
<<<<<<< HEAD
		w.Write ([]byte("<h1>Hello World 2 </h1>"))
=======
		w.Write ([]byte("<h1>Minha aplicação com GITOPS e ArgoCD</h1>"))
>>>>>>> 2a7e086 (ArgoCD em açãoC)
	})

	http.ListenAndServe(":8080", nil)
}
