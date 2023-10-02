# Google Oauth
Sample dilakukan dengan bahasa pemrograman Go.
Sebelum masuk masuk ke koding, kita perlu membuat credentials Oauth di console Google Cloud.

## Membuat Credentials Oauth
<p>Pertama-tama, kunjungi <a href="https://console.cloud.google.com/apis/credentials">Google Cloud Console</a>.
Setelah itu klik <b>CREATE CREDENTIALS</b> pada navbar atas dan pilih <b>OAuth client ID</b>.</p>
<img  src="img/console-1.png">
<p>Setelah itu pada pilihan <b>Application type</b> pilih <b>Web Application</b>.
Pada field <b>Name</b> silahkan diisi dengan nama credentials. 
Nama ini sebagai identitas yang akan memudahkan kita dalam mengelola credentials Oauth.</p>
<img src="img/console-2.png">
<p>Pada section <b>Authorized Javascript origins</b>. Tambahkan URL server dengan menekan tombol <b>Add URIs</b>.
Di sini akan kita isi dengan <b>http://localhost:8080</b>. Lanjut di bagian Authorized <b>Redirect URIs</b> silahkan diisi dengan URL yang akan digunakan untuk redirect ke halaman Google di sini contohnya <b>http://localhost:8080/auth/google/callback</b>. Jika sudah, kita bisa simpan dengan menekan tombol <b>Create</b>
</p>
<p>Setelah itu, akan ada popup yang menampilkan Client ID beserta secret milik kita. Silahkan copy masing-masing Client ID dan secret</p>
<img src="img/console-3.png">
<p>Sampai di tahap ini, pembuatan credentials Oauth sudah selesai.</p>

## Implementasi Go 
<p>Siapkan folder project.</p>


```sh 
mkdir belajar-oauth && cd belajar-oauth && go mod init belajar-oauth
```

<p>Import library berikut untuk berinteraksi dengan Oauth Google.</p>


```go 
import (
    "golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)
```

Setelah itu install libary dengan perintah berikut:

```sh
go mod tidy
```

Selanjutnya, kita initial config OAuth Credentials milik kita.


```go 
var oauth2.Config{
    RedirectURL: "http://localhost:8080/auth/google/callback", // Samakan dengan yang ada pada Redirect URIs
    ClientID: "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_SECRET",
    Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"}, // Scope profile yang akan diambil
    Endpoint: google.Endpoint,
}
```

Jangan lupa deklarasikan juga OAuth URL.

```go 
const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/v2/userinfo?access_token="
```

Buat handler untuk redirect ke OAuth Google.

```go 
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    url := oauthConfig.AuthCodeURL("state")
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
```

Terapkan handler tadi route.

```go 
http.HandleFunc("/auth/google/login", handleGoogleLogin)
```

Buat callback handler Google OAuth.

```go 
func handleCallbackGoogle(w http.ResponseWriter, r *http.Request) {
    // State harus unique
    if r.FormValue("state") != "state" {
        fmt.Println("Invalid oauth google state")
        return
    }

    data, err := getUserDataFromGoogle(r.FormValue("code"))

    if err != nil {
        fmt.Println("Failed to retrieve user information:", err)
        return
    }

    // Me-return response 
    fmt.Fprintf(w, "UserInfo: %s \n", data)
    
}
```
Buat helper function untuk mendapatkan Informasi User.

```go 
func getUserDataFromGoogle(code string) ([]byte, error) {
    // Gunakan parameter "code" untuk mendapat token dan informasi user 
    token ,err := oauthConfig.Exchange(context.Background(), code)
    
    if err != nil {
        return nil, fmt.Errorf("code exchange wrong")
    }

    resp, err := http.Get(oauthGoogleURLAPI + token.AccessToken)

    if err != nil {
        return nil, fmt.Errorf("failed getting user info: %s", err.Error())
    }

    defer resp.Body.Close()
    
    contents, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return nil, fmt.Errorf("failed read response: %s", err.Error())
    }

    return contents, nil
}


```

Selesai.
