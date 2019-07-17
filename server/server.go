package main
import(
	"fmt"
	"net/http"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
)
var mySigningKey = []byte("mysupersecretkey")

func HomePage(w http.ResponseWriter, r*http.Request){
fmt.Fprint(w,"JSMPJ Corporation success")
}
func isAuthorizes(endpoints func(http.ResponseWriter, *http.Request)) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    if r.Header["Token"] != nil {
       token,err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token)(interface{},error){
        if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there is an error")
		}
		return mySigningKey, nil
	})
	    if err != nil{
			fmt.Fprintf(w, err.Error())
		}
		if token.Valid{
			endpoints(w,r)
		}
	
	} else{
		fmt.Fprintf(w, "Not authorized")
	}
	})
}
func handleRequests(){

	http.Handle("/",isAuthorizes(HomePage))
	log.Fatal(http.ListenAndServe(":8081",nil))
}
func main(){
	fmt.Println("JSMPJ Server")
	handleRequests()
	
}