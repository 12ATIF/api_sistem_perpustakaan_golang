# Dokumentasi API

Ini adalah dokumentasi untuk REST API Sistem Manajemen Perpustakaan.

## Base URL

Semua URL yang tercantum di sini menggunakan base URL: `/api`

## Otentikasi

Endpoint untuk otentikasi tidak memerlukan token JWT.

| Method | Endpoint         | Deskripsi                  |
| ------ | ---------------- | -------------------------- |
| `POST` | `/auth/login`    | Login untuk mendapatkan token. |
| `POST` | `/auth/register` | Registrasi pengguna baru.  |

---

## Rute Terproteksi

Semua endpoint di bawah ini memerlukan header `Authorization: Bearer <token>` dengan token JWT yang valid dan role **admin**.

### 1. Anggota

| Method   | Endpoint        | Deskripsi                       |
| -------- | --------------- | ------------------------------- |
| `GET`    | `/anggota/`     | Mendapatkan semua data anggota. |
| `GET`    | `/anggota/:id`  | Mendapatkan anggota by ID.      |
| `POST`   | `/anggota/`     | Membuat anggota baru.           |
| `PUT`    | `/anggota/:id`  | Memperbarui data anggota.       |
| `DELETE` | `/anggota/:id`  | Menghapus data anggota.         |

### 2. Kategori

| Method   | Endpoint         | Deskripsi                        |
| -------- | ---------------- | -------------------------------- |
| `GET`    | `/kategori/`     | Mendapatkan semua data kategori. |
| `GET`    | `/kategori/:id`  | Mendapatkan kategori by ID.      |
| `POST`   | `/kategori/`     | Membuat kategori baru.           |
| `PUT`    | `/kategori/:id`  | Memperbarui data kategori.       |
| `DELETE` | `/kategori/:id`  | Menghapus data kategori.         |

### 3. Penerbit

| Method   | Endpoint         | Deskripsi                       |
| -------- | ---------------- | ------------------------------- |
| `GET`    | `/penerbit/`     | Mendapatkan semua data penerbit.|
| `GET`    | `/penerbit/:id`  | Mendapatkan penerbit by ID.     |
| `POST`   | `/penerbit/`     | Membuat penerbit baru.          |
| `PUT`    | `/penerbit/:id`  | Memperbarui data penerbit.      |
| `DELETE` | `/penerbit/:id`  | Menghapus data penerbit.        |

### 4. Buku

| Method   | Endpoint      | Deskripsi                    |
| -------- | ------------- | ---------------------------- |
| `GET`    | `/buku/`      | Mendapatkan semua data buku. |
| `GET`    | `/buku/:id`   | Mendapatkan buku by ID.      |
| `POST`   | `/buku/`      | Membuat buku baru.           |
| `PUT`    | `/buku/:id`   | Memperbarui data buku.       |
| `DELETE` | `/buku/:id`   | Menghapus data buku.         |

### 5. Peminjaman

| Method   | Endpoint         | Deskripsi                |
| -------- | ---------------- | ------------------------ |
| `POST`   | `/peminjaman/`   | Membuat data peminjaman. |

### 6. Pengembalian

| Method   | Endpoint           | Deskripsi                  |
| -------- | ------------------ | -------------------------- |
| `POST`   | `/pengembalian/`   | Membuat data pengembalian. |