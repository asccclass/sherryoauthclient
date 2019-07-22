package main

import(
   "os"
   "fmt"
   "bytes"
   "net/http"
   "encoding/json"
)

type ReplayPackage struct {
   Token	string	`json:"token"`
   Message	string	`json:"message"`
   Status	string	`json:"status"`
}

type Payload struct {
   Username string `json:"username"`
   Password string `json:"password"`
}

type Oauth struct {
   OauthServer string	// OAuth Server IP Address
}

// 取得Token
func (o *Oauth) Login(username, password string)(string, error) {
   return "", nil
}

func NewOauthClient(OauthServerUrl string)(*Oauth) {
   return &Oauth{
      OauthServer: OauthServerUrl,
   }
}


func main() {
   // or := NewOauthClient("https://wteamapi.its.sinica.edu.tw/coursehours/dorelogin")

   os.Setenv("HTTP_PROXY", "http://140.109.12.18:3128")

   data := Payload{"xeplusplatform", "12345"}
   payloadBytes, err := json.Marshal(data)
   if err != nil {
	panic("error payload")
   }
   body := bytes.NewReader(payloadBytes)

   req, err := http.NewRequest("POST", "http://localhost/coursehours/dorelogin", body)
   if err != nil {
      panic(err)
   }
   req.Header.Set("Content-Type", "application/json")

   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   x := ReplayPackage{}
   // b, _ := ioutil.ReadAll(resp.Body)
   err = json.NewDecoder(resp.Body).Decode(&x)
   if err != nil {
      panic(err)
   }
   y, _ := json.Marshal(x)
   fmt.Printf("%v", string(y))
}
