# RestFul API GO Kontak Panggilan Darurat Rumah Sakit

## Daftar Isi

- [Pra-kata](#pra-kata)

- [Tata Cara *clone project*](#Clone-project-ini)

- [Dokumentasi](#Documentation)
    - [HTTP Requests](#HTTP-Requests)
    - [HTTP Responses](#HTTP-Responses)
    - [HTTP Responses Code](#HTTP-Responses-Code)
    - [GET](#GET)
    - [POST](#POST)
    - [PUT](#PUT)
    - [DELETE](#DELETE)

- [Authorization & Authentication](#Authorization-&-Authentication)

- [Contoh penggunaan melalui Postman](#Tutorial-Postman)


---
## Pra-kata
---
Kontak Panggilan Darurat Rumah Sakit adalah tugas project-2 pada Mata Kuliah `Pemrograman Integratif` yang dibuat oleh `Herwinda Marwaa` dengan `NRP 05311940000009`. 
Berikut merupakan beberapa informasi umum mengenai RestFul API ini:
- Restful API dibangun menggunakan bahasa pemrograman `Go` dan *framework* `echo labstack v4`.
- API ini memiliki 4 metode dan 1 endpoint.
- Pada *Authentication* dan *Authorization* penulis menerapkan `JWT Token`. 
- `Metode GET` dari `daftarInstansi` yang tersedia pada *database* dapat diakses oleh semua pengguna
- Pengguna yang telah terdaftar pada database akan mendapatkan `JWT Token`, sehingga pengguna tersebut dapat mengakses Restful API ini secara penuh dengan 
ke - 4 method yang telah di buat `(GET, POST, PUT, DELETE)`.

<br>


## Clone project ini
---
1. Pertama, install `Golang` beserta dengan framework `Echo labstack`. Jika belum silahkan klik tautan berikut : <br>
[Golang](https://golang.org/doc/install) dan [Echo Labstack Framework](https://echo.labstack.com/guide/installation/)
2. Setelah itu, *clone* project pada github ini dengan `GO-PATH` yang telah didefinisikan saat instalasi `Golang`.
3. Atur `config.JSON` maupun `config.go` pada folder *config* sesuai dengan device masing-masing.
4. Lakukan inisialisasi `go mod init`  dan dilanjutkan `go mod tidy` serta `go mod vendor` untuk otomatisasi *import* terhadap `library` pendukung yang dibutuhkan. 
5. Pada header terdapat `file import go` yang harus disesuaikan terhadap penyimpanan folder masing-masing *device*.
6. Atur `Routes` pada file `routes.go`.
7. Lalu lakukan `run and execution` program dengan perintah <br> `go run main.go`.
8. Jika terdapat **warning**, jalankan `go mod vendor` terlebih dahulu.

<br>

## Documentation
---
Docummentation secara lengkap dapat dilihat pada:
[Swagger](http://localhost:1234/documentation/index.html)

<br>

## Authorization-&-Authentication
---
- Untuk proses *Authorization* dan *Authentication*, penulis membutuhkan bantuan `JSON Web Token` yang diintegrasikan dengan [library](https://wwww.github.com/dgrijalva/jwt-go). 
`JWT` akan  diberikan kepada pengguna yang telah terdaftar di *database* agar dapat mengakses ke-4 metode dari `Restful API` yang telah dibuat. 
Oleh karena itu saya telah menyiapkan sebuah akun *dummy* dengan `username : Herwinda24` dan `password : Herwinda24`. 

- Sebelumnya, akan ada pencocokan antara `password plainteks` dengan `password cipherteks` yang dikonversikan ke bentuk [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt). 

- Berikut adalah *screenshot* dari `JWT Token` yang didapatkan *user* setelah melakukan validasi *login* melalui `POSTMAN` :
![Postman1]()<br>
Pada *screenshot* ditunjukan menggunakan `metode POST` dengan url `localhost:1234/login?username=Herwinda24&password=Herwinda24`
<br>

Adapun untuk *Authorization* dan *Authentication* untuk *Guest* atau pengguna tidak terdaftar tanpa melalui *login*  hanya dapat mengakses `metode GET` dari Daftar Instansi Rumah Sakit yang tersedia di *database* seperti gambar berikut :
![Postman2]()<br>

<br>


## Tutorial Postman
---
