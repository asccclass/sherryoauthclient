package main  //SherryOauthClient

import(
   "os"
   "fmt"
   "time"
   "bytes"
   "net/http"
   "net/url"
   "encoding/json"
)

// 回傳訊息內容
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
   OauthServer	string	// OAuth Server IP Address
   ProxyUrl	string
}

// 取得Token
func (o *Oauth) Login(username, password string)(string, error) {
   if username == "" || password == "" {
      return "", fmt.Errorf("Account or password is empty.")
   }
   data := Payload{username, password}
   payloadBytes, err := json.Marshal(data)
   if err != nil { return "", err }
   var netClient = &http.Client{
      Timeout: time.Second * 10,
   }
   if o.ProxyUrl != "" {
      proxyURL, err := url.Parse(o.ProxyUrl)
      if err != nil {
         return "", err
      }
fmt.Println(proxyURL)
      netClient.Transport = &http.Transport{
         Proxy: http.ProxyURL(proxyURL),
      }
   }
   req, err := http.NewRequest("POST", o.OauthServer, bytes.NewReader(payloadBytes))
   if err != nil { return "", err }
   req.Header.Set("Content-Type", "application/json")
   resp, err := netClient.Do(req)
   if err != nil {
      return "", fmt.Errorf("http.Do error(%v)", err)
   }
   defer resp.Body.Close()
   x := ReplayPackage{}
   if err = json.NewDecoder(resp.Body).Decode(&x); err != nil {
      return "", err
   }
   y, _ := json.Marshal(x)
   return string(y), nil
}

func(o *Oauth)SetProxy(url string) {
   if url != "" {
      o.ProxyUrl = url
      os.Setenv("HTTP_PROXY", o.ProxyUrl)
   }
}

func NewOauthClient(OauthServerUrl string)(*Oauth, error) {
   if OauthServerUrl == "" {
      return nil, fmt.Errorf("Must have Oauth Server's Url.")
   }

   return &Oauth{
      OauthServer: OauthServerUrl,
      ProxyUrl: "",
   }, nil
}


func main() {
   // oauth, err := NewOauthClient("http://wteamapi.its.sinica.edu.tw/coursehours/dorelogin")
   oauth, err := NewOauthClient("http://devwteamapi.test5.sinica.edu.tw/coursehours/dorelogin")
   if err != nil {
      panic(err)
   }
   oauth.SetProxy("http://140.109.12.18:3128")
   token, err := oauth.Login("eplusplatform", "12345")
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Printf("%s\n", token)
}
