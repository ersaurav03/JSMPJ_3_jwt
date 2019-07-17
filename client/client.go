package main
import(
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"io/ioutil"
)
var mySignKey = []byte("mysupersecretkey")
func GenerateJWT()(string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorization"]=true
	claims["user"]="Saurav Kumar"
	claims["exp"]=time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySignKey)
	if err != nil {
		fmt.Errorf("Something went wrong : %s",err.Error())
		return "", err
	}
	return tokenString, nil
}
func HomePage(w http.ResponseWriter, r *http.Request){
	validToken, err := GenerateJWT()
	if err != nil {
		panic("Something went wrong")
	}
	Client := &http.Client{}
		req, _ := http.NewRequest("GET","http://localhost:8080/",nil)
		req.Header.Set("Token",validToken)
	res, err := Client.Do(req)
	if err != nil {
		panic("Some thing went wrong in neew client")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error in response")
	}


	fmt.Fprintf(w, string(body))
}
func HadnleRequest(){
newRouter:=mux.NewRouter()
newRouter.HandleFunc("/",HomePage).Methods("GET")
log.Fatal(http.ListenAndServe(":8080",newRouter))
}
func main(){
	fmt.Println("JSMPJ Client")
	HadnleRequest()
	// tokenString, err := GenerateJWT()
	// if err != nil{
	// 	panic("something went wrong")
	// }
	// fmt.Println("this is generated token",tokenString)
}