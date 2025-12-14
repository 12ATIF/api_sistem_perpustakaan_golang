# Dokumentasi API Sistem Perpustakaan

Ini adalah dokumentasi dasar untuk RESTful API Sistem Perpustakaan.

## Autentikasi

### 1. Register User

- **Endpoint:** `POST /api/auth/register`
- **Deskripsi:** Mendaftarkan user baru (admin).
- **Request Body:**
  ```json
  {
    "name": "Admin User",
    "email": "admin@example.com",
    "password": "password123",
    "role": "admin"
  }
  ```
- **Response Sukses (200):**
  ```json
  {
    "status": true,
    "message": "User created successfully",
    "data": {
      "id": 1,
      "name": "Admin User",
      "email": "admin@example.com",
      "role": "admin",
      "created_at": "2025-12-14T14:00:00Z",
      "updated_at": "2025-12-14T14:00:00Z"
    }
  }
  ```

### 2. Login User

- **Endpoint:** `POST /api/auth/login`
- **Deskripsi:** Login untuk mendapatkan token JWT.
- **Request Body:**
  ```json
  {
    "email": "admin@example.com",
    "password": "password123"
  }
  ```
- **Response Sukses (200):**
  ```json
  {
    "status": true,
    "message": "Login successful",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
  }
  ```

---

## CRUD (Protected)

Semua endpoint di bawah ini memerlukan header `Authorization: Bearer <token>`.

### Anggota

- **Endpoint:** `GET /api/anggota`
- **Deskripsi:** Mendapatkan daftar anggota dengan pagination dan search.
- **Query Params:** `page` (number), `pageSize` (number), `search` (string, by nama)
- **Contoh:** `GET /api/anggota?page=1&pageSize=5&search=john`
- **Response Sukses (200):**
  ```json
  {
    "status": true,
    "message": "All anggota",
    "data": {
      "data": [
        {
          "id": 1,
          "nama": "John Doe",
          "alamat": "Jl. Pahlawan No. 1",
          "no_telp": "08123456789",
          "created_at": "2025-12-14T14:05:00Z",
          "updated_at": "2025-12-14T14:05:00Z"
        }
      ],
      "total": 1,
      "page": 1,
      "pageSize": 5,
      "totalPages": 1
    }
  }
  ```

- **Endpoint:** `POST /api/anggota`
- **Deskripsi:** Membuat anggota baru.
- **Request Body:**
  ```json
  {
    "nama": "Jane Doe",
    "alamat": "Jl. Merdeka No. 10",
    "no_telp": "08987654321"
  }
  ```

---

### Buku

- **Endpoint:** `GET /api/buku`
- **Deskripsi:** Mendapatkan daftar buku dengan pagination dan search.
- **Query Params:** `page` (number), `pageSize` (number), `search` (string, by judul)
- **Contoh:** `GET /api/buku?page=1&pageSize=10&search=golang`
- **Response Sukses (200):**
  ```json
  {
    "status": true,
    "message": "All buku",
    "data": {
      "data": [
        {
            "id": 1,
            "judul": "Belajar Golang",
            "penulis": "John Doe",
            "tahun_terbit": 2023,
            "stok": 10,
            "kategori_id": 1,
            "penerbit_id": 1,
            "created_at": "2025-12-14T14:10:00Z",
            "updated_at": "2025-12-14T14:10:00Z"
        }
      ],
      "total": 1,
      "page": 1,
      "pageSize": 10,
      "totalPages": 1
    }
  }
  ```
---

### Transaksi

#### 1. Peminjaman Buku

- **Endpoint:** `POST /api/peminjaman`
- **Deskripsi:** Membuat transaksi peminjaman baru. Stok buku akan otomatis berkurang.
- **Request Body:**
  ```json
  {
    "anggota_id": 1,
    "tanggal_pinjam": "2025-12-14",
    "tanggal_kembali": "2025-12-21",
    "detail": [
      {
        "buku_id": 1,
        "qty": 1
      }
    ]
  }
  ```

#### 2. Pengembalian Buku

- **Endpoint:** `POST /api/pengembalian`
- **Deskripsi:** Membuat transaksi pengembalian. Stok buku akan otomatis bertambah dan denda akan dihitung jika terlambat.
- **Request Body:**
  ```json
  {
    "peminjaman_id": 1,
    "tanggal_pengembalian": "2025-12-22"
  }
  ```
- **Response Sukses (200):**
  ```json
  {
    "status": true,
    "message": "Pengembalian created successfully",
    "data": {
        "id": 1,
        "peminjaman_id": 1,
        "tanggal_pengembalian": "2025-12-22T00:00:00Z",
        "denda": 1000,
        "created_at": "2025-12-14T14:15:00Z",
        "updated_at": "2025-12-14T14:15:00Z"
    }
  }
  ```
