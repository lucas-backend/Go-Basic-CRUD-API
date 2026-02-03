# Go Basic CRUD API

Proyek ini adalah implementasi sederhana dari RESTful API untuk manajemen data pengguna (User) menggunakan bahasa pemrograman Go (Golang). Proyek ini mencakup operasi CRUD (Create, Read, Update, Delete) dasar dan terhubung ke database MySQL.

## Teknologi yang Digunakan

Proyek ini dibangun menggunakan teknologi dan library berikut:

*   **[Go](https://go.dev/)** (Versi 1.25.6) - Bahasa pemrograman utama.
*   **[Gorilla Mux](https://github.com/gorilla/mux)** - Router HTTP yang kuat dan fleksibel untuk Go.
*   **[Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)** - Driver database MySQL untuk Go `database/sql`.
*   **[Godotenv](https://github.com/joho/godotenv)** - Library untuk memuat variabel lingkungan dari file `.env`.

## Prasyarat

Sebelum menjalankan proyek ini, pastikan kalian telah menginstal:

1.  **Go** (Golang) terinstal di sistem kalian.
2.  **MySQL** atau **MariaDB** sebagai server database.
3.  **Git** (opsional, untuk kloning repositori).

## Panduan Instalasi dan Konfigurasi

Ikuti langkah-langkah berikut untuk menjalankan proyek di komputer lokal kalian:

### 1. Buka Proyek
Pastikan kalian berada di direktori root proyek ini.

```bash
cd path/to/my-api
```

### 2. Instal Dependensi
Unduh semua paket yang diperlukan menggunakan `go mod`.

```bash
go mod tidy
```

### 3. Konfigurasi Database
Buat database baru di MySQL dan buat tabel `users` (atau sesuaikan dengan kebutuhan kalian). Berikut adalah contoh query SQL sederhana untuk membuat tabel yang sesuai dengan model `User` di proyek ini:

```sql
CREATE DATABASE nama_database_kalian;

USE nama_database_kalian;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### 4. Konfigurasi Environment Variables (.env)
Buka file `.env` (atau buat baru jika belum ada) dan sesuaikan konfigurasinya dengan kredensial database kalian.

Contoh isi file `.env`:

```env
SERVER_PORT=8080

DB_USER=root
DB_PASSWORD=password_database_kalian
DB_HOST=localhost
DB_PORT=3306
DB_NAME=nama_database_kalian
```

### 5. Jalankan Server
Jalankan aplikasi menggunakan perintah berikut:

```bash
go run main.go
```

Jika berhasil, kalian akan melihat pesan seperti:
`Using API Server berjalan di port 8080` dan `Database berhasil terhubung!`

## Dokumentasi API Endpoint

Berikut adalah daftar endpoint yang tersedia:

| Method | Endpoint     | Deskripsi               | Body Request (JSON) |
| :---   | :---         | :---                    | :--- |
| `GET`  | `/user/`     | Mengambil semua data users | - |
| `GET`  | `/user/{id}` | Mengambil data user berdasarkan ID | - |
| `POST` | `/user/`     | Menambahkan user baru | `{"username": "...", "password": "..."}` |
| `PUT`  | `/user/{id}`| Mengupdate username user | `{"username": "..."}` |
| `DELETE`| `/user/{id}`| Menghapus user berdasarkan ID | - |

---
**Catatan:** Proyek ini adalah proyek latihan pribadi untuk mempelajari dasar-dasar pembuatan API dengan Go.
