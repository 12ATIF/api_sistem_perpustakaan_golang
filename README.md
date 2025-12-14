# Sistem Manajemen Perpustakaan

Ini adalah REST API untuk aplikasi Sistem Manajemen Perpustakaan yang dibangun menggunakan bahasa pemrograman Go.

## Fitur

*   Authentikasi & Autorisasi Pengguna (JWT)
*   Manajemen Buku
*   Manajemen Anggota
*   Manajemen Kategori Buku
*   Manajemen Penerbit Buku
*   Manajemen Peminjaman
*   Manajemen Pengembalian

## Struktur Proyek

Proyek ini mengikuti tata letak proyek Go standar:

```
Nama_file/
├── cmd/
│   └── api/
│       └── main.go         # Entry point aplikasi
├── config/
│   └── database.go         # Konfigurasi database
├── internal/
│   ├── entity/             # Struct/model data
│   ├── handler/            # Handler untuk request HTTP
│   ├── middleware/         # Middleware HTTP
│   ├── repository/         # Logika akses database
│   └── usecase/            # Logika bisnis
├── pkg/
│   └── jwt.go              # Utilitas untuk JWT
├── go.mod
├── go.sum
└── .env                    # File untuk environment variables
```

## Instalasi & Menjalankan

1.  **Clone repository:**
    ```bash
    git clone <url-repository-anda>
    cd coba_dulu
    ```

2.  **Konfigurasi Environment:**
    Buat file `.env` di root direktori dan sesuaikan variabel di dalamnya (misalnya, untuk koneksi database).

    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=password
    DB_NAME=perpustakaan
    JWT_SECRET=secret
    ```

3.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Jalankan aplikasi:**
    ```bash
    go run ./cmd/api/main.go
    ```

    Aplikasi akan berjalan di port yang ditentukan (default biasanya `:8080`).

## Dokumentasi API

Untuk detail endpoint API, silakan lihat file `docs/api.md`.
