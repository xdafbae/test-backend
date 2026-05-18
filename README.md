# SharingVision Backend

REST API untuk manajemen artikel menggunakan **Go + Gin + GORM + MySQL**.

## Requirements

- Go 1.21+
- MySQL 8.0+

## Tech Stack

| Layer | Library |
|---|---|
| Language | Go 1.21+ |
| Framework | Gin |
| ORM | GORM |
| Database | MySQL |
| Validation | go-playground/validator (via Gin binding) |
| Env | godotenv |

## Setup

1. **Clone repo**
   ```bash
   git clone <repo-url>
   cd backend
   ```

2. **Buat database MySQL**
   ```sql
   CREATE DATABASE article;
   ```
   Atau jalankan script lengkap:
   ```bash
   mysql -u root -p < database.sql
   ```

3. **Konfigurasi environment**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` sesuai kredensial MySQL Anda:
   ```
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=your_password
   DB_NAME=article
   ```

4. **Install dependencies**
   ```bash
   go mod tidy
   ```

5. **Jalankan server**
   ```bash
   go run main.go
   ```

Server berjalan di `http://localhost:8080`

## Struktur Project

```
backend/
├── main.go                        ← Entry point
├── go.mod
├── go.sum
├── .env                           ← Kredensial DB (tidak di-commit)
├── .env.example                   ← Template env
├── database.sql                   ← SQL manual untuk buat tabel
├── SharingVision.postman_collection.json
├── config/
│   └── database.go                ← Koneksi MySQL + AutoMigrate
├── models/
│   └── post.go                    ← Struct Post + TableName()
├── handlers/
│   └── post_handler.go            ← Semua handler CRUD
└── routes/
    └── routes.go                  ← Setup Gin router + CORS
```

## Endpoints

| Method | URL | Deskripsi |
|---|---|---|
| POST | `/article` | Buat artikel baru |
| GET | `/article/:limit/:offset` | Get semua artikel dengan paging |
| GET | `/article/:id` | Get artikel by ID |
| PUT | `/article/:id` | Update artikel by ID |
| DELETE | `/article/:id` | Hapus artikel by ID |

## Validasi

| Field | Rule |
|---|---|
| title | required, min 20 karakter |
| content | required, min 200 karakter |
| category | required, min 3 karakter |
| status | required, salah satu: `publish` / `draft` / `thrash` |

## Contoh Request

### POST /article
```json
{
  "title": "Judul artikel minimal dua puluh karakter",
  "content": "Konten artikel minimal dua ratus karakter. Lorem ipsum dolor sit amet, consectetur adipiscing elit...",
  "category": "Tech",
  "status": "publish"
}
```

### Response 201
```json
{
  "id": 1,
  "title": "Judul artikel minimal dua puluh karakter",
  "content": "...",
  "category": "Tech",
  "created_date": "2024-01-01T00:00:00Z",
  "updated_date": "2024-01-01T00:00:00Z",
  "status": "publish"
}
```

## Postman Collection

Import file `SharingVision.postman_collection.json` ke Postman.  
Set variable `base_url = http://localhost:8080`.
