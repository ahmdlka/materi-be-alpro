# Materi Backend Camin ALPRO

---

## Daftar Isi

- [Backend 101: Konsep Fundamental](#backend-101-konsep-fundamental)
  - [Apa Itu Backend?](#apa-itu-backend)
  - [Jenis-Jenis Database](#jenis-jenis-database)
    - [Relational Database (SQL)](#relational-database-sql)
    - [NoSQL Database](#nosql-database)
  - [Authentication: Session vs JWT](#authentication-session-vs-jwt)
    - [Session-Based Authentication (Stateful)](#session-based-authentication-stateful)
    - [JWT -- JSON Web Token (Stateless)](#jwt----json-web-token-stateless)
  - [FAQ](#faq)
- [Environment Setup](#environment-setup)
  - [Instalasi Go](#instalasi-go)
  - [Setup Project Go](#setup-project-go)
- [Go Crash Course](#go-crash-course)
  - [Variabel & Tipe Data](#variabel--tipe-data)
  - [Array & Slices](#array--slices)
  - [Fungsi](#fungsi)
  - [Struct](#struct)
  - [Error Handling](#error-handling)
- [Backend Di Go](#backend-di-go)
  - [Gin Web Framework](#gin-web-framework)
  - [Gorm Library](#gorm-library)
  - [Contoh Boilerplate](#contoh-boilerplate)
    - [Alur Request -- Dari HTTP ke Database](#alur-request----dari-http-ke-database)
    - [Pola Per-Module](#pola-per-module)
    - [Cara Jalankan Project](#cara-jalankan-project)
- [Latihan Implementasi](#latihan-implementasi)
  - [Demo: `POST /users` (Create User)](#demo-post-users-create-user)
  - [Challenge A -- `GET /users/:id`](#challenge-a----get-usersid)
  - [Challenge B -- `GET /users`](#challenge-b----get-users)

---

## Backend 101: Konsep Fundamental

Sebelum menulis kode, penting untuk memahami konsep-konsep dasar yang menjadi fondasi setiap sistem backend, terlepas dari bahasa pemrograman yang digunakan.

### Apa Itu Backend?

Ketika kamu membuka sebuah aplikasi (misalnya Instagram), yang kamu lihat di layar adalah **frontend** -- tampilan visual yang kamu interaksikan. Tapi di balik layar, ada **backend** -- server yang memproses data, menyimpan informasi ke database, mengecek apakah kamu sudah login, dan mengirimkan data yang diminta oleh frontend.

Secara singkat:
- **Frontend** = apa yang dilihat pengguna (UI/tampilan)
- **Backend** = apa yang terjadi di belakang layar (server, database, logic)
- **API (Application Programming Interface)** = "kontrak" komunikasi antara frontend dan backend. Frontend mengirim **request**, backend membalas dengan **response**.

> [!NOTE]
> Dalam tutorial ini, kita akan membangun sebuah **REST API** -- jenis API yang paling umum digunakan. REST API menggunakan metode HTTP seperti `GET` (ambil data), `POST` (kirim data baru), `PUT` (update data), dan `DELETE` (hapus data).

### Jenis-Jenis Database

Tidak ada satu jenis database yang cocok untuk semua kebutuhan. Pilihan database harus disesuaikan dengan karakteristik data dan kebutuhan sistem.

#### Relational Database (SQL)

**Contoh:** PostgreSQL, MySQL, SQLite

Menyimpan data dalam bentuk tabel dengan baris dan kolom. Hubungan antar tabel didefinisikan secara eksplisit.

| Kapan digunakan | Contoh kasus |
|---|---|
| Data terstruktur dengan skema yang jelas | Data transaksi keuangan |
| Butuh ACID (Atomicity, Consistency, Isolation, Durability) | Sistem pemesanan tiket |
| Ada relasi kompleks antar entitas | Sistem manajemen pengguna |

```sql
-- Contoh query SQL
SELECT users.name, orders.total
FROM users
JOIN orders ON users.id = orders.user_id
WHERE users.id = 1;
```

> [!NOTE]
> SQL (Structured Query Language) adalah bahasa yang digunakan untuk berkomunikasi dengan relational database. `SELECT` untuk mengambil data, `FROM` untuk menentukan tabel, `JOIN` untuk menggabungkan data dari tabel berbeda, dan `WHERE` untuk memfilter.

#### NoSQL Database

**Contoh:** MongoDB (document), Redis (key-value), Cassandra (wide-column)

Menyimpan data dalam format yang lebih fleksibel -- bisa dokumen JSON, pasangan key-value, graph, dll.

| Kapan digunakan | Contoh kasus |
|---|---|
| Skema data sering berubah | Katalog produk e-commerce |
| Volume data sangat besar dan perlu scale horizontal | Social media feed |
| Kebutuhan read/write sangat cepat (caching) | Session storage, rate limiting |

> [!TIP]
> **Aturan praktis:** Mulai dengan SQL. Pindah ke NoSQL hanya jika ada alasan teknis yang jelas, bukan karena terasa "lebih modern".

---

### Authentication: Session vs JWT

Authentication adalah proses verifikasi identitas: *"Apakah kamu benar-benar siapa yang kamu klaim?"*

> [!NOTE]
> Bayangkan authentication seperti penjaga pintu masuk gedung. Kamu harus menunjukkan kartu identitas (username + password) untuk masuk. Setelah diverifikasi, kamu diberi tanda pengenal (token/session) agar tidak perlu menunjukkan kartu identitas setiap kali berpindah ruangan.

#### Session-Based Authentication (Stateful)

Server menyimpan informasi sesi di memorinya (atau di database/Redis). Setiap request dari client membawa session ID, dan server mencarinya di penyimpanan.

```
[Client]  ->  Login dengan username/password
[Server]  ->  Buat session, simpan di memory/Redis
[Server]  ->  Kirim session_id ke client (biasanya via cookie)
[Client]  ->  Setiap request berikutnya bawa cookie session_id
[Server]  ->  Lookup session_id, verifikasi, proses request
```

> [!NOTE]
> "Stateful" artinya server harus **mengingat** siapa yang sudah login. Ibarat hotel yang mencatat semua tamu di buku besar -- setiap kali tamu datang, petugas harus cek buku dulu.

| Kelebihan | Kekurangan |
|---|---|
| Mudah di-revoke (hapus session dari server) | Stateful: server harus simpan data session |
| Cocok untuk aplikasi monolitik | Sulit di-scale horizontal (session harus di-share antar server) |

#### JWT -- JSON Web Token (Stateless)

Server tidak menyimpan apapun. Semua informasi yang dibutuhkan ada di dalam token yang dikirim oleh client.

```
[Client]  ->  Login dengan username/password
[Server]  ->  Generate JWT, tanda-tangani dengan secret key
[Server]  ->  Kirim JWT ke client
[Client]  ->  Simpan JWT (localStorage atau httpOnly cookie)
[Client]  ->  Setiap request kirim JWT di header: Authorization: Bearer <token>
[Server]  ->  Verifikasi signature JWT dengan secret key -- tidak perlu database lookup
```

> [!NOTE]
> "Stateless" artinya server **tidak perlu mengingat** apapun. Semua informasi sudah ada di dalam token itu sendiri. Ibarat surat izin resmi yang bertanda tangan -- siapapun bisa memverifikasi bahwa surat itu asli hanya dari tanda tangannya, tanpa perlu menelepon kantor yang mengeluarkan surat.

Struktur JWT terdiri dari tiga bagian yang dipisahkan titik:
```
header.payload.signature

eyJhbGciOiJIUzI1NiJ9.eyJ1c2VySWQiOjF9.SflKxwRJSMeKKF2QT4fwpMeJf36P
```

| Kelebihan | Kekurangan |
|---|---|
| Stateless: mudah di-scale | Sulit di-revoke sebelum expired |
| Tidak butuh database lookup untuk verifikasi | Payload ter-encode (bukan terenkripsi) -- jangan simpan data sensitif |
| Bisa digunakan lintas domain/service | Token besar -> overhead di setiap request |

> [!TIP]
> **Rekomendasi:** JWT cocok untuk microservices dan API publik. Session cocok untuk web app tradisional di mana kontrol penuh atas session lebih penting.

---

### FAQ

Pertanyaan yang sering muncul:

**Q: Kapan saya harus pakai Gin vs framework lain seperti Echo atau Fiber?**
A: Gin adalah pilihan yang aman untuk pemula -- dokumentasinya banyak dan komunitasnya besar. Echo dan Fiber juga baik; perbedaannya minor untuk project skala kecil-menengah.

**Q: Apakah Go wajib untuk backend?**
A: Tidak. Go unggul di performa, concurrency, dan binary deployment. Tapi Python, Node.js, atau Java juga valid tergantung kebutuhan tim dan project.

**Q: Apa langkah selanjutnya setelah tutorial ini?**
A: Untuk menambah pengetahuan mengenai Go, dapat mengikuti [tour ini](https://go.dev/tour/list).

---

## Environment Setup
### Instalasi Go

Kamu dapat mendownload Go di website resmi mereka [disini](https://go.dev/doc/install).

<img width="1862" height="926" alt="image" src="https://github.com/user-attachments/assets/2aa526f3-1bf1-4770-82c5-a07a520b0c1a" />

Setelah penginstallan selesai, kamu dapat mengecek jika berhasil di terminal, dan menjalankan command berikut:

```bash
go version
# Output yang diharapkan: go version go1.21.x linux/amd64 (atau sejenisnya)
```

### Setup Project Go

Setelah Go berhasil diinstall, kamu dapat membuat project dengan menjalankan command di bawah:

```bash
# Buat folder project
mkdir go-workshop && cd go-workshop

# Inisialisasi Go module
go mod init {username kamu}/go-workshop
```

> [!NOTE]
> `mkdir` artinya "make directory" (buat folder). Tanda `&&` berarti "setelah command pertama selesai, jalankan command kedua." `cd` artinya "change directory" (pindah ke folder tersebut).

`go mod init` bertanggung jawab untuk membuat file `go.mod`, yang dapat kamu anggap seperti peletak project kamu dengan dependency di dalam ataupun di luar project.  

Isi di dalam `go.mod` kamu harusnya ini.
```
module github.com/username/go-workshop

go 1.26.2
```

> [!NOTE]
> `go.mod` mirip dengan `package.json` di Node.js atau `requirements.txt` di Python -- file ini mencatat nama project dan semua library eksternal yang dipakai.

Untuk mulai menulis kode kamu, kamu dapat membuat file baru `main.go` sebagai file mulainya program kamu.  
Sebagai pemanasan, kita coba print Hello world sebagai tradisi mempelajari bahasa baru!

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```

> [!NOTE]
> **Catatan sintaks Go:**
> - `package main` -- Menandakan file ini adalah program yang bisa dijalankan langsung. Setiap program Go yang bisa dieksekusi **harus** memiliki `package main`.
> - `import "fmt"` -- Mengimpor package `fmt` (singkatan dari "format"), package bawaan Go untuk mencetak teks ke terminal.
> - `func main()` -- Fungsi entry point, titik mulai program. Ketika program dijalankan, Go akan mencari dan menjalankan fungsi `main()` ini pertama kali. Mirip `public static void main` di Java.
> - `fmt.Println(...)` -- Mencetak teks ke terminal dan otomatis menambah baris baru di akhir.

```bash
go run main.go
# Output: Hello world!
```

> [!NOTE]
> `go run` mengompilasi dan langsung menjalankan file Go. Kamu tidak perlu compile manual seperti di C/C++.

## Go Crash Course

### Variabel & Tipe Data

Go memiliki dua cara mendeklarasikan variabel:

```go
// Cara 1: deklarasi eksplisit dengan 'var'
var nama string = "Budi"
var umur int = 25

// Cara 2: short declaration dengan ':='
nama := "Budi"
umur := 25
isActive := true
```

> [!NOTE]
> - `var nama string = "Budi"` -- Deklarasi eksplisit: kamu menulis tipe datanya (`string`) secara manual.
> - `nama := "Budi"` -- Short declaration: Go otomatis menebak tipe datanya berdasarkan nilai yang diberikan. Cara ini **hanya bisa dipakai di dalam fungsi**, bukan di level global.
> - Tipe data dasar di Go: `string` (teks), `int` (bilangan bulat), `float64` (bilangan desimal), `bool` (true/false).

Kamu juga dapat mendeklarasikan 2 variabel seperti di bahasa Python

```go
result, err := getName()
```

> [!NOTE]
> Go memperbolehkan fungsi mengembalikan **lebih dari satu nilai**. Ini sangat umum di Go -- biasanya nilai kedua adalah `error` (lihat bagian [Error Handling](#error-handling)).

### Array & Slices

Kamu dapat mendeklarasikan array seperti berikut:

```go
angka = [7]int{1, 2, 3, 4, 5, 6, 7}

// Kamu dapat mendapatkan nilai-nilai dalam array seperti berikut
fmt.Println(angka[3])

// Mendapatkan dari index 1 ke 3.
fmt.Println(angka[1:4])

// Mendapatkan dari index 0 ke 1.
fmt.Println(angka[:2])
```

> [!NOTE]
> `[7]int` artinya array dengan kapasitas tepat 7 elemen bertipe `int`. Angka di dalam `[]` menentukan ukuran array. Indeks dimulai dari `0`, jadi `angka[3]` adalah elemen ke-4 (nilainya `4`).

Array dalam Go bersifat statis dan besarnya tidak bisa diganti.

Kamu juga dapat menyimpan cuplikan dari array dalam bentuk slice. Slice berupa reference ke array sebenarnya, dan seluruh perubahan pada slice akan berlaku di arraynya, dan vice versa.

```go
a = angka[0:2]
b = angka[1:4]
fmt.Println(a, b)

a[1] = 10
fmt.Println(a, b)
```

> [!NOTE]
> Slice adalah "jendela" ke sebuah array. Karena slice adalah **reference** (referensi), mengubah isi slice juga mengubah array aslinya. Di contoh di atas, `a[1] = 10` juga akan mengubah `angka[1]` dan mempengaruhi `b`.

### Fungsi

```go
// Fungsi tanpa return value
func sapa(nama string) {
    fmt.Println("Halo,", nama)
}

// Fungsi dengan return value
func tambah(a int, b int) int {
    return a + b
}

// Fungsi dengan multiple return (idiom khas Go)
func bagi(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("tidak bisa dibagi nol")
    }
    return a / b, nil
}
```

> [!NOTE]
> **Catatan sintaks fungsi:**
> - `func sapa(nama string)` -- `func` adalah keyword untuk mendefinisikan fungsi. `nama string` berarti parameter `nama` bertipe `string`. Tidak ada return value.
> - `func tambah(a int, b int) int` -- Bagian `int` terakhir setelah kurung tutup adalah **tipe return value**.
> - `func bagi(a, b float64) (float64, error)` -- Fungsi ini mengembalikan **dua nilai**: hasil pembagian (`float64`) dan error. `(float64, error)` di dalam kurung menandakan multiple return.
> - `nil` artinya "tidak ada" atau "kosong" -- mirip `null` di Java/JavaScript. Di sini, `nil` berarti tidak ada error.
> - `fmt.Errorf(...)` membuat objek error baru dengan pesan yang kamu tentukan.

### Struct

Go bukan bahasa yang mendukung object oriented programming. Sebagai gantinya, gunakan `struct` untuk mengelompokkan data.

```go
// Mendefinisikan struct
type User struct {
    ID    uint
    Name  string
    Email string
}

// Membuat instance dari struct
user := User{
    ID:    1,
    Name:  "Budi",
    Email: "budi@email.com",
}

fmt.Println(user.Name) // Output: Budi
```

> [!NOTE]
> - `type User struct { ... }` -- Mendefinisikan tipe data baru bernama `User`. Mirip `class` di Java/Python, tapi **tanpa method bawaan dan tanpa inheritance**.
> - `uint` artinya "unsigned integer" -- bilangan bulat yang tidak bisa negatif (0, 1, 2, ...).
> - `user.Name` -- Mengakses field `Name` dari struct `user`, mirip mengakses property objek di bahasa lain.
> - Di Go, **huruf besar di awal nama** (`Name`, `Email`) berarti field tersebut bersifat **publik** (bisa diakses dari package lain). Huruf kecil (`name`) berarti **privat**.

### Error Handling

Go tidak menggunakan `try-catch`. Error dikembalikan sebagai nilai return biasa.  

Go memiliki aturan yang sangat ketat yang ditetapkan, salah satunya adalah **tidak bolehnya ada variabel yang tidak dipakai**. Sebagai efeknya, error yang dikembalikan wajib diurus oleh programmer, jika tidak, program Go tidak akan jalan!

```go
result, err := bagi(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Hasil:", result)
```

> [!IMPORTANT]
> Pola `if err != nil` adalah **pola paling dasar dan paling penting** di Go. Hampir setiap fungsi yang bisa gagal mengembalikan `error` sebagai return value terakhir. Kamu **wajib** mengecek apakah error-nya `nil` (kosong) atau ada isinya sebelum melanjutkan. Ini adalah cara Go memastikan programmer tidak mengabaikan error.

Untuk menambah pengetahuan mengenai Go, dapat mengikuti [tour ini](https://go.dev/tour/list).


## Backend Di Go

Go adalah salah satu bahasa yang paling populer dalam backend development, dengan alasan-alasan berikut:
- Go cukup gampang untuk dimengerti, sehingga kode-kode yang dibuat lebih readable dan gampang dimaintain
- Go adalah bahasa yang sangat _opinionated_ sampai aturan formattingnya ditetapkan secara universal, sehingga bahkan seorang BE developer yang tidak pernah menyentuh codebase, dapat mengerti control flow nya dengan cepat.
- Fitur-fitur dasar Go sangat cocok untuk backend development, berupa `net/http` dan `goroutines`, dan sudah cukup robust untuk production, dan tidak terlalu perlu memanggil dependency dari library eksternal.
- Go sangat gampang untuk dideploy, karena semua hampir dependency dapat diinstall dengan Go sebagai "package manager" nya.

Walau `net/http` cukup untuk membangun projek Backend yang solid, kita akan memakai framework web `gin` untuk menghindari _reinventing the wheel_.

### Gin Web Framework

Pada singkatnya, Gin adalah framework web yang membuat routing, parsing request, dan penulisan response JSON menjadi jauh lebih ringkas.

> [!NOTE]
> Framework adalah kumpulan library/alat yang menyediakan struktur siap pakai, sehingga kamu tidak perlu membangun semuanya dari nol. Gin mengurus hal-hal teknis seperti parsing HTTP request, routing URL, dan mengelola middleware -- kamu tinggal fokus menulis logic bisnis.

Untuk menginstall gin, jalankan command ini di direktori projek Go kalian.
```bash
# Install Gin
go get github.com/gin-gonic/gin
```

> [!NOTE]
> `go get` adalah command untuk mendownload dan menginstall library/package eksternal ke project Go kamu. Library yang didownload akan otomatis tercatat di file `go.mod`.

Kamu dapat merubah `main.go` kamu menjadi dibawah untuk mencoba apakah Gin berhasil diinstall.
```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  // Create a Gin router with default middleware (logger and recovery)
  r := gin.Default()

  // Define a simple GET endpoint
  r.GET("/ping", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  r.Run()
}
```

> [!NOTE]
> **Catatan sintaks Gin:**
> - `import ( ... )` -- Kurung buka-tutup digunakan untuk mengimpor **beberapa package sekaligus**.
> - `"net/http"` -- Package bawaan Go yang berisi konstanta HTTP seperti `http.StatusOK` (kode 200), `http.StatusBadRequest` (kode 400), dll.
> - `r := gin.Default()` -- Membuat **router** Gin dengan middleware bawaan (logger untuk mencatat request, dan recovery untuk menangkap panic agar server tidak crash).
> - `r.GET("/ping", ...)` -- Mendaftarkan **endpoint** yang akan dipanggil saat ada HTTP GET request ke URL `/ping`.
> - `func(c *gin.Context)` -- Fungsi anonim (tanpa nama) yang menerima parameter `c` bertipe pointer ke `gin.Context`. Context ini berisi semua informasi tentang request yang masuk dan menyediakan method untuk mengirim response.
> - `*gin.Context` -- Tanda `*` artinya **pointer** (referensi ke data, bukan data-nya sendiri). Kamu akan sering melihat ini di Go; untuk sekarang, cukup tahu bahwa `c` adalah objek yang berisi data request dan response.
> - `c.JSON(http.StatusOK, ...)` -- Mengirim response dalam format JSON dengan status code 200.
> - `gin.H{ ... }` -- Shortcut untuk membuat map/dictionary sederhana yang akan dikonversi jadi JSON. `gin.H{"message": "pong"}` menjadi `{"message": "pong"}`.

Jalankan dengan `go run main.go`, lalu buka browser ke `GET http://localhost:8080/ping`. Hasilnya:

```json
{
  "message": "pong"
}
```

---

### Gorm Library

ORM (Object-Relational Mapper) adalah alat yang menjembatani kode Go dengan database. Dengan GORM, kamu tidak perlu menulis SQL secara manual untuk operasi CRUD dasar.

> [!NOTE]
> CRUD adalah singkatan dari **Create, Read, Update, Delete** -- empat operasi dasar yang dilakukan terhadap data di database. ORM memungkinkan kamu menulis query database menggunakan kode Go (struct dan method), bukan menulis SQL langsung.

Untuk menginstall gorm, jalankan command ini di direktori projek Go kalian. Workshop ini akan memakai `postgres` sebagai sistem databasenya.
```bash
# Install GORM dan driver PostgreSQL
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

GORM memetakan `struct` Go menjadi tabel di database secara otomatis.:

```go
type User struct {
    ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    Name     string    `gorm:"not null" json:"name"`
    Email    string    `gorm:"unique;not null" json:"email"`
    Password string    `gorm:"not null" json:"-"` // tidak muncul di JSON response
}
```

> [!NOTE]
> **Catatan tentang struct tag:**
> Teks di dalam backtick (`` ` ``) setelah tipe data disebut **struct tag**. Tag ini memberi instruksi khusus ke library tertentu:
> - `gorm:"not null"` -- Memberi tahu GORM bahwa kolom ini **tidak boleh kosong** di database.
> - `gorm:"unique;not null"` -- Kolom harus unik (tidak boleh ada duplikat) dan tidak boleh kosong.
> - `gorm:"primary_key"` -- Menandai kolom ini sebagai primary key tabel.
> - `json:"name"` -- Memberi tahu encoder JSON bahwa field ini harus tampil sebagai `"name"` di output JSON.
> - `json:"-"` -- Tanda minus berarti field ini **disembunyikan** dari JSON response. Berguna untuk data sensitif seperti password.
> - `uuid.UUID` -- Tipe data UUID (Universally Unique Identifier), yaitu ID unik yang panjang seperti `550e8400-e29b-41d4-a716-446655440000`.

Konfigurasi database dipisahkan ke `config/database.go`:
Dalam struct GORM, kamu dapat memberikan atribut ke baris di entity dengan menambahkan field tags.  
Untuk field tags yang lengkap, dapat dilihat di [dokumentasi GORM](https://gorm.io/docs/models.html#Fields-Tags)

```go
package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"-"` // tidak muncul di JSON response
}

Database dijalankan via Docker agar tidak perlu install PostgreSQL secara lokal:
func main() {
	// Harusnya menggunakan .env but for demonstration only
	dsn := "host=localhost user=username password=password dbname=go_workshop port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Karena user menggunakan uuid
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	db.AutoMigrate(&User{})

	newUser := User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "hashed_password_here",
	}
	db.Create(&newUser)
	fmt.Println("Created User ID:", newUser.ID)

	var user User
	db.First(&user, "id = ?", newUser.ID)
	fmt.Println("User:", user)

	db.Delete(&user, "id = ?", user.ID)
	fmt.Println("Deleted user!")
}
```
> [!NOTE]
> - `dsn` (Data Source Name) -- String yang berisi semua informasi untuk koneksi ke database (alamat server, username, password, nama database, port).
> - `os.Getenv("DB_HOST")` -- Mengambil nilai dari **environment variable**. Environment variable adalah pengaturan yang disimpan di luar kode (biasanya di file `.env`) agar data sensitif seperti password tidak di-hardcode.
> - `fmt.Sprintf(...)` -- Seperti `fmt.Println` tapi tidak mencetak ke terminal, melainkan **mengembalikan string** dengan format yang ditentukan. `%s` adalah placeholder yang akan diganti nilainya.
> - `&gorm.Config{}` -- Tanda `&` mengambil **alamat memori** (pointer) dari objek. Ini kebalikan dari `*` yang mengambil nilai dari pointer. Untuk sekarang, cukup ingat: `&` = "berikan referensi ke objek ini."
> - `panic(...)` -- Menghentikan program secara paksa dengan pesan error. Digunakan untuk error yang **tidak bisa dipulihkan** seperti gagal koneksi database.

Perlu diingat bahwa kamu perlu mensetup Postgres kamu terlebih dahulu:
- Buat user di postgres `(CREATE USER your_username WITH PASSWORD 'your_password';)`
- Buat database di postgres `(CREATE DATABASE your_dbname;)`
- Berikan user akses ke database tersebut `(ALTER DATABASE your_dbname OWNER TO your_username;)`

---

Jika kita menggabungkan kedua library ini, bentuk `main.go` akan seperti ini:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"-"` // Does not appear in JSON response
}

func main() {
	// Harusnya menggunakan .env but for demonstration only
	dsn := "host=localhost user=postgres password=root dbname=go_workshop port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Karena user menggunakan uuid
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.AutoMigrate(&User{})
	fmt.Println("Database connected and migrated successfully!")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/users", func(c *gin.Context) {
		var input struct {
			Name     string `json:"name" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser := User{
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password, // Di dalam praktek, jangan lupa untuk menghash password terlebih dahulu!
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": newUser})
	})

	r.GET("/users", func(c *gin.Context) {
		var users []User

		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": users})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		if err := db.Delete(&User{}, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})

	fmt.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
```

Melihat kode `main.go`, hasilnya cukup panjang dan ribet bukan? Dan ini baru satu API CRUD entity.  
Oleh karena itu, walaupun tidak diwajibkan oleh bahasa, kita perlu struktur project yang jelas dan terstruktur.

### Contoh Boilerplate

Boilerplate singkat kata adalah template yang sering dipakai berulang-ulang dan biasanya sedikit / tidak ada variasi, yang bertujuan untuk memberikan struktur ke project kamu.  

Terdapat beberapa boilerplate backend Go yang tersedia, tetapi di workshop ini kita akan menggunakan [boilerplate ini](https://github.com/Caknoooo/go-gin-clean-starter/tree/main).

Berikut adalah struktur folder boilerplate yang kita gunakan, beserta penjelasan tanggung jawab masing-masing:

```
.
|-- cmd/
|   +-- main.go                        # Entry point -- hanya panggil providers & jalankan server
|
|-- config/
|   |-- database.go                    # Setup koneksi GORM ke PostgreSQL
|   |-- email.go                       # Konfigurasi SMTP untuk kirim email
|   |-- logger.go                      # Setup query logger ke file
|   +-- logs/query_log                 # Output log query SQL
|
|-- database/
|   |-- entities/                      # Definisi struct yang di-mapping ke tabel DB
|   |   |-- common.go                  # Base struct (ID, CreatedAt, UpdatedAt, dll)
|   |   |-- user_entity.go             # Tabel `users`
|   |   +-- refresh_token_entity.go    # Tabel `refresh_tokens`
|   |-- migrations/                    # Versi skema database (seperti git untuk DB)
|   |   |-- 20240101000000_create_users_table.go
|   |   +-- 20240101000001_create_refresh_tokens_table.go
|   |-- seeders/                       # Data awal untuk development/testing
|   |   |-- json/users.json            # Data seed dalam format JSON
|   |   +-- seeds/user_seed.go         # Logic untuk insert seed data
|   |-- manager.go                     # Orchestrator: jalankan migrate & seed
|   |-- migration.go                   # Runner untuk file-file migration
|   +-- seeder.go                      # Runner untuk file-file seeder
|
|-- middlewares/
|   |-- authentication.go              # Validasi JWT token di setiap protected route
|   +-- cors.go                        # Izinkan/blokir request lintas domain
|
|-- modules/                           # <-- FOKUS UTAMA workshop
|   |-- auth/                          # Semua yang berkaitan dengan login/logout/token
|   |   |-- controller/auth_controller.go
|   |   |-- dto/auth_dto.go
|   |   |-- repository/refresh_token_repository.go
|   |   |-- service/auth_service.go
|   |   |-- service/jwt_service.go
|   |   |-- validation/auth_validation.go
|   |   |-- tests/auth_validation_test.go
|   |   +-- routes.go
|   +-- user/                          # Semua yang berkaitan dengan data user
|       |-- controller/user_controller.go
|       |-- dto/user_dto.go
|       |-- query/user_query.go
|       |-- repository/user_repository.go
|       |-- service/user_service.go
|       |-- validation/user_validation.go
|       |-- tests/user_validation_test.go
|       +-- routes.go
|
|-- pkg/
|   |-- constants/common.go            # Konstanta global (pesan error, status, dll)
|   |-- helpers/password.go            # Helper bcrypt: hash & compare password
|   +-- utils/
|       |-- aes.go                     # Enkripsi/dekripsi data sensitif
|       |-- email.go                   # Fungsi kirim email via SMTP
|       |-- file.go                    # Helper upload & manajemen file
|       +-- response.go                # Standarisasi format JSON response
|
|-- providers/
|   +-- core.go                        # Dependency injection: wiring semua layer
|
|-- script/
|   |-- command.go                     # Definisi command CLI (migrate, seed, dll)
|   +-- script.go                      # Runner untuk perintah dari terminal
|
|-- docker/
|   |-- Dockerfile                     # Build image untuk production
|   |-- nginx/default.conf             # Konfigurasi reverse proxy
|   +-- postgresql/                    # Konfigurasi PostgreSQL container
|
|-- docker-compose.yml                 # Jalankan seluruh stack (App + DB + Nginx)
|-- Makefile                           # Shortcut command (make migrate, make run, dll)
+-- create_module.sh                   # Script otomatis buat module baru
```

> [!NOTE]
> **Istilah-istilah penting:**
> - **Migration** -- Cara mengubah struktur database secara terkontrol dan terlacak. Mirip "version control untuk database." Setiap perubahan skema (tambah tabel, tambah kolom) dicatat dalam file migration.
> - **Seeder** -- Script untuk mengisi database dengan data awal/contoh, berguna saat development agar tidak perlu input data manual setiap kali.
> - **Middleware** -- Kode yang dijalankan **sebelum** request sampai ke handler utama. Contoh: mengecek apakah user sudah login (JWT valid) sebelum memproses request.
> - **CORS (Cross-Origin Resource Sharing)** -- Mekanisme keamanan browser yang mengontrol apakah website dari domain A boleh mengakses API di domain B.
> - **Dependency Injection** -- Pola di mana sebuah komponen tidak membuat sendiri dependensinya, melainkan "menerima" dependensi dari luar. Contoh: controller menerima service, service menerima repository.
> - **DTO (Data Transfer Object)** -- Struct khusus yang mendefinisikan bentuk data yang diterima dari request atau dikirim sebagai response. Memisahkan format data API dari format data database.

#### Alur Request -- Dari HTTP ke Database

Setiap request yang masuk melewati lapisan-lapisan ini secara berurutan:

```
HTTP Request
    |
    v
[Middleware]           -> Auth check (JWT valid?), CORS header
    |
    v
[Controller]           -> Terima request Gin, panggil Validation, kirim response
    |
    v
[Validation]           -> Validasi input (field required, format email, dll)
    |
    v
[Service]              -> Business logic (hash password, generate token, dll)
    |
    v
[Repository]           -> Query ke database via GORM
    |
    v
[Database / Entity]    -> PostgreSQL
```

#### Pola Per-Module

Setiap fitur baru dibuat dalam satu folder `modules/<nama_fitur>/` yang memiliki struktur seragam:

| File | Tanggung Jawab |
|---|---|
| `controller/` | Terima `*gin.Context`, parsing input, kirim JSON response |
| `dto/` | Struct untuk request body & response (Data Transfer Object) |
| `validation/` | Aturan validasi input sebelum masuk ke service |
| `service/` | Business logic -- tidak boleh tahu soal HTTP atau database secara langsung |
| `repository/` | Semua query GORM -- satu-satunya layer yang boleh sentuh DB |
| `query/` | Raw query atau filter kompleks yang dipakai oleh repository |
| `routes.go` | Daftarkan semua endpoint milik module ini |
| `tests/` | Unit test untuk validation & service |

> [!TIP]
> **Pola ini membuat kode mudah ditemukan.** Kalau ada bug di response format, cari di `controller`. Kalau ada bug di kalkulasi bisnis, cari di `service`. Kalau ada bug di query lambat, cari di `repository`.

#### Cara Jalankan Project

```bash
# Clone dan install dependency
go mod tidy
# `go mod tidy` mendownload semua dependency yang tercatat di go.mod
# dan menghapus dependency yang tidak dipakai.

# Jalankan seluruh stack dengan Docker
docker-compose up -d

# Jalankan migrasi database
make migrate
# atau: go run script/script.go migrate

# Jalankan seeder (isi data awal)
make seed
# atau: go run script/script.go seed

# Jalankan development server
go run cmd/main.go
```

---

## Latihan Implementasi

Teori sudah cukup. Sekarang waktunya tangan kotor dengan kode sungguhan di dalam boilerplate. Semua file yang dibuat mengikuti pola yang sudah ada di module `auth` dan `user`.

### Demo: `POST /users` (Create User)

Bagian ini mendemonstrasikan alur pembuatan satu endpoint utuh, dari mendaftarkan route sampai data tersimpan ke database.

**`database/entities/user_entity.go`** -- Skema tabel
```go
package entities

type User struct {
    Common                    // embed ID, CreatedAt, UpdatedAt, DeletedAt
    Name     string `gorm:"not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Role     string `gorm:"default:'user'"`
}
```

> [!NOTE]
> - `Common` -- Ini disebut **struct embedding**. Struct `User` "mewarisi" semua field dari struct `Common` (biasanya berisi `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`). Mirip konsep inheritance tapi tanpa polymorphism.
> - `gorm:"default:'user'"` -- Jika field `Role` tidak diisi saat membuat user baru, database akan otomatis mengisinya dengan `'user'`.

**`modules/user/dto/user_dto.go`** -- Shape data masuk & keluar
```go
package dto

type CreateUserRequest struct {
    Name     string `json:"name"     binding:"required"`
    Email    string `json:"email"    binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Role  string `json:"role"`
}
```

> [!NOTE]
> - `binding:"required"` -- Tag dari Gin yang berarti field ini **wajib diisi**. Jika client tidak mengirimnya, Gin akan otomatis mengembalikan error.
> - `binding:"required,email"` -- Selain wajib, nilainya juga **harus berformat email** yang valid.
> - `binding:"required,min=8"` -- Selain wajib, panjang string **minimal 8 karakter**.
> - `CreateUserRequest` digunakan untuk **menerima** data dari client, sedangkan `UserResponse` digunakan untuk **mengirim** data kembali ke client (tanpa password!).

**`modules/user/validation/user_validation.go`** -- Validasi input
```go
package validation

import (
    "github.com/gin-gonic/gin"
    "github.com/username/boilerplate/modules/user/dto"
)

func ValidateCreateUser(c *gin.Context) (*dto.CreateUserRequest, error) {
    var req dto.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        return nil, err
    }
    return &req, nil
}
```

> [!NOTE]
> - `c.ShouldBindJSON(&req)` -- Membaca body request (JSON), lalu **mengisinya ke struct** `req`. Jika JSON tidak valid atau field yang `required` kosong, fungsi ini mengembalikan error.
> - `var req dto.CreateUserRequest` -- Membuat variabel `req` dengan tipe `CreateUserRequest` (bernilai kosong/default).
> - `return &req, nil` -- Mengembalikan **pointer** ke `req` dan `nil` (tidak ada error). Pointer digunakan agar data tidak perlu di-copy (lebih efisien untuk struct besar).

**`modules/user/repository/user_repository.go`** -- Query database
```go
package repository

import (
    "github.com/username/boilerplate/database/entities"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *entities.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*entities.User, error) {
    var user entities.User
    err := r.db.First(&user, id).Error
    return &user, err
}

func (r *UserRepository) FindAll() ([]entities.User, error) {
    var users []entities.User
    err := r.db.Find(&users).Error
    return users, err
}
```

> [!NOTE]
> - `func (r *UserRepository) Create(...)` -- Ini disebut **method**. Bedanya dengan fungsi biasa: method terikat ke sebuah struct. `(r *UserRepository)` disebut **receiver** -- artinya method `Create` milik struct `UserRepository`, dan bisa diakses via `r`.
> - `r.db.Create(user)` -- Method GORM untuk insert data ke database. GORM otomatis membuat query `INSERT INTO users ...` dari struct yang diberikan.
> - `r.db.First(&user, id)` -- Mengambil satu record pertama berdasarkan primary key (`id`). Mirip `SELECT * FROM users WHERE id = ? LIMIT 1`.
> - `r.db.Find(&users)` -- Mengambil semua record. Mirip `SELECT * FROM users`.
> - `.Error` -- Setiap operasi GORM mengembalikan result yang memiliki field `.Error`. Jika operasi berhasil, nilainya `nil`.
> - `[]entities.User` -- Tanda `[]` di depan tipe artinya **slice** (array dinamis). Jadi `[]entities.User` berarti "kumpulan User."

**`modules/user/service/user_service.go`** -- Business logic
```go
package service

import (
    "github.com/username/boilerplate/database/entities"
    "github.com/username/boilerplate/modules/user/dto"
    "github.com/username/boilerplate/modules/user/repository"
    "github.com/username/boilerplate/pkg/helpers"
)

type UserService struct {
    repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*entities.User, error) {
    // Business logic: hash password sebelum disimpan
    hashedPassword, err := helpers.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    user := &entities.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashedPassword,
    }

    err = s.repo.Create(user)
    return user, err
}
```

> [!NOTE]
> - `helpers.HashPassword(req.Password)` -- Mengubah password plain text menjadi **hash** (string acak yang tidak bisa di-decode balik). Ini adalah praktik keamanan standar -- password **tidak boleh** disimpan dalam bentuk plain text di database.
> - Service layer **tidak boleh tahu** tentang HTTP atau database secara langsung. Ia hanya menerima data (DTO), memproses business logic, dan memanggil repository untuk operasi database.
> - Pola `New...()` (seperti `NewUserService`, `NewUserRepository`) adalah **constructor pattern** di Go -- fungsi yang membuat dan mengembalikan instance baru dari sebuah struct.

**`modules/user/controller/user_controller.go`** -- Handle HTTP
```go
package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/username/boilerplate/modules/user/service"
    "github.com/username/boilerplate/modules/user/validation"
    "github.com/username/boilerplate/pkg/utils"
)

type UserController struct {
    service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
    return &UserController{service: service}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    req, err := validation.ValidateCreateUser(c)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    user, err := ctrl.service.CreateUser(req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Gagal membuat user")
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "User berhasil dibuat", user)
}
```

> [!NOTE]
> - `http.StatusBadRequest` = kode 400 (request dari client salah/tidak valid).
> - `http.StatusInternalServerError` = kode 500 (error di sisi server).
> - `http.StatusCreated` = kode 201 (data berhasil dibuat).
> - `err.Error()` -- Mengonversi objek error menjadi string pesan error yang bisa dibaca.
> - `utils.ErrorResponse` dan `utils.SuccessResponse` -- Fungsi helper untuk memastikan semua response API memiliki **format yang seragam** (konsisten).

**`modules/user/routes.go`** -- Daftarkan endpoint
```go
package user

import (
    "github.com/gin-gonic/gin"
    "github.com/username/boilerplate/middlewares"
    "github.com/username/boilerplate/modules/user/controller"
)

func RegisterUserRoutes(r *gin.RouterGroup, ctrl *controller.UserController) {
    users := r.Group("/users")
    {
        users.POST("", ctrl.CreateUser)                                          // POST /api/users
        users.GET("", middlewares.Authentication(), ctrl.GetAllUsers)            // GET  /api/users (protected)
        users.GET("/:id", middlewares.Authentication(), ctrl.GetUserByID)        // GET  /api/users/:id (protected)
    }
}
```

> [!NOTE]
> - `r.Group("/users")` -- Membuat **route group**. Semua endpoint di dalam group ini otomatis diawali dengan `/users`.
> - `users.POST("", ...)` -- Karena sudah di group `/users`, endpoint ini menjadi `POST /users`.
> - `users.GET("/:id", ...)` -- `:id` disebut **URL parameter**. Nilainya dinamis -- misalnya `/users/5` berarti `id = 5`.
> - `middlewares.Authentication()` -- Middleware yang mengecek JWT token. Endpoint yang dibungkus middleware ini hanya bisa diakses oleh user yang sudah login. `POST /users` (registrasi) tidak memerlukan login, jadi tanpa middleware.

> [!IMPORTANT]
> Perhatikan `middlewares.Authentication()` -- endpoint yang membutuhkan login dibungkus middleware ini. Middleware akan memvalidasi JWT token sebelum request diteruskan ke controller.

---

### Challenge A -- `GET /users/:id`

> Ambil satu user berdasarkan ID. Kembalikan `404` jika tidak ditemukan.

### Challenge B -- `GET /users`

> Ambil semua user. Kembalikan array JSON.

> [!WARNING]
> **Tips debugging:** Error paling umum di Go adalah lupa menangani `if err != nil`. Kalau ada `panic`, cari baris yang tidak menangani error return-nya. Gunakan `fmt.Println(err)` untuk print error ke terminal.

---

<div align="center">

*Tutorial ini adalah awal, bukan akhir. Yang paling penting adalah terus membangun sesuatu.*

**Selamat ngoding!**

</div>
# be-alpro
# materi-be-alpro
# materi-be-alpro
