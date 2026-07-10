# Online Shop Backend

## Target Project

Membuat backend untuk aplikasi **Online Shop** menggunakan **Golang** dan **PostgreSQL**.

---

# Menjalankan Database

```bash
docker run \
  --name postgresql \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=database \
  -d \
  -p 5432:5432 \
  postgres:16
```

---

# Modul yang Digunakan

| Modul | Kegunaan |
|--------|----------|
| `go get gin-gonic/gin` | Web framework (routing, middleware, binding request, validation) |
| `database/sql` | Menjalankan query database |
| `go get  github.com/jackc/pgx/v5/stdlib` | Driver PostgreSQL |
| `go get  github.com/google/uuid` | Membuat UUID |
| `go get  golang.org/x/crypto` | Hashing password menggunakan bcrypt |

---

# Fitur

## Fitur yang Akan Dibuat

1. Melihat produk
2. Checkout
3. Konfirmasi pembayaran
4. Melihat detail pesanan
5. Menambah produk (Admin)
6. Memperbarui produk (Admin)
7. Menghapus produk (Admin)

## Di Luar Cakupan Project

- Pengelolaan data admin
- Pengelolaan data user
- Pengelolaan data pesanan
- Register
- Login
- Upload file
- Notifikasi email

---

# Data Model

## Product

| Field | Tipe | Keterangan |
|--------|------|------------|
| id | UUID (string) | Identifier produk |
| name | string | Nama produk |
| price | integer | Harga produk |
| is_deleted | boolean | Soft delete |

> Product tidak benar-benar dihapus dari database agar histori transaksi tetap terhubung.

---

## Order

| Field | Tipe | Keterangan |
|--------|------|------------|
| id | UUID (string) | Identifier pesanan |
| email | string | Email pemesan |
| address | string | Alamat pengiriman |
| grand_total | integer | Total harga pesanan |
| passcode | string (hash) | Password akses pesanan |
| paid_at | timestamp | Tanggal pembayaran |
| paid_bank | string | Nama bank pembayaran |
| paid_account_number | string | Nomor rekening pembayaran |

> Karena project ini tidak memiliki sistem login, maka setiap pesanan memiliki **passcode** yang disimpan dalam bentuk **bcrypt hash**.

---

## Order Detail

Merupakan tabel **junction** antara **Order** dan **Product** sekaligus menyimpan histori harga produk saat checkout.

| Field | Tipe | Keterangan |
|--------|------|------------|
| id | UUID | Identifier detail pesanan |
| order_id | UUID | ID pesanan |
| product_id | UUID | ID produk |
| quantity | integer | Jumlah produk |
| price | integer | Harga produk saat checkout |
| total | integer | Total harga |

---

# Endpoint

## Public Endpoint

| Method | Endpoint | Keterangan |
|---------|----------|------------|
| GET | `/api/v1/products` | Melihat seluruh produk |
| GET | `/api/v1/products/{id}` | Melihat detail produk |
| POST | `/api/v1/checkout` | Checkout dan membuat pesanan |

Semua endpoint menggunakan **API Versioning** (`v1`).

---

## Endpoint Menggunakan Passcode

| Method | Endpoint | Keterangan |
|---------|----------|------------|
| GET | `/api/v1/orders/{id}` | Melihat detail pesanan |
| POST | `/api/v1/orders/{id}/confirm` | Konfirmasi pembayaran |

Secara implementasi endpoint ini dapat dibatasi agar hanya dapat diakses dari jaringan internal.

---

## Endpoint Admin

| Method | Endpoint | Keterangan |
|---------|----------|------------|
| POST | `/admin/products` | Tambah produk |
| PUT | `/admin/products/{id}` | Update produk |
| DELETE | `/admin/products/{id}` | Soft delete produk |

Semua endpoint admin menggunakan **Authorization Header** melalui middleware.

---

# Alur Fitur

## 1. Lihat Produk

Alur:

1. Request diterima.
2. Ambil seluruh produk yang belum dihapus dari database.
3. Kembalikan response.

---

## 2. Lihat Detail Produk

Alur:

1. Request diterima.
2. Ambil `id` dari URL.
3. Cari produk berdasarkan `id`.
4. Kembalikan response.

---

## 3. Checkout / Membuat Pesanan

Alur:

1. Request diterima.
2. Baca email.
3. Baca alamat.
4. Baca daftar produk beserta jumlahnya.
5. Ambil data produk dari database.
6. Hitung total harga.
7. Generate passcode.
8. Hash passcode menggunakan bcrypt.
9. Simpan data order.
10. Simpan seluruh order detail.
11. Kembalikan response beserta passcode.

---

## Passcode

Karena project ini tidak memiliki sistem login, maka akses pesanan menggunakan **passcode**.

Passcode digunakan untuk:

- Melihat detail pesanan.
- Melakukan konfirmasi pembayaran.

Untuk endpoint yang tidak memiliki request body (misalnya `GET`), passcode dikirim melalui **query parameter**.

Passcode hanya diberikan satu kali saat checkout dan disimpan dalam bentuk hash.

---

## 4. Konfirmasi Pembayaran

Alur:

1. Request diterima.
2. Baca passcode.
3. Baca data pembayaran.
4. Ambil data pesanan.
5. Validasi passcode.
6. Pastikan pesanan belum pernah dibayar.
7. Pastikan nominal pembayaran sesuai.
8. Simpan informasi pembayaran.
9. Kembalikan response.

---

## 5. Lihat Detail Pesanan

Alur:

1. Request diterima.
2. Baca ID pesanan.
3. Ambil data pesanan.
4. Validasi passcode.
5. Ambil seluruh order detail.
6. Kembalikan response.

---

# Otorisasi Admin

Seluruh endpoint admin menggunakan middleware untuk memvalidasi **Authorization Header** sebelum request diproses.

---

## Admin - Tambah Produk

Alur:

1. Request diterima.
2. Baca nama produk.
3. Baca harga produk.
4. Simpan ke database.
5. Kembalikan response.

---

## Admin - Perbarui Produk

Alur:

1. Request diterima.
2. Baca ID produk.
3. Baca nama produk.
4. Baca harga produk.
5. Cari produk berdasarkan ID.
6. Perbarui data produk.
7. Kembalikan response.

---

## Admin - Hapus Produk

Alur:

1. Request diterima.
2. Baca ID produk.
3. Ubah `is_deleted = TRUE`.
4. Kembalikan response.

---

# Database

## Products

```sql
CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price BIGINT NOT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);
```

---

## Orders

```sql
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    passcode VARCHAR(255) NOT NULL,
    grand_total BIGINT NOT NULL,
    paid_at TIMESTAMP NULL,
    paid_bank VARCHAR(255),
    paid_account_number VARCHAR(255)
);
```

---

## Order Details

```sql
CREATE TABLE IF NOT EXISTS order_details (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity INT NOT NULL,
    price BIGINT NOT NULL,
    total BIGINT NOT NULL,
    FOREIGN KEY (order_id)
        REFERENCES orders(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);
```

---

# SQL

## Ambil Semua Produk

```sql
SELECT
    id,
    name,
    price
FROM products
WHERE is_deleted = FALSE;
```

---

## Ambil Produk Berdasarkan ID

```sql
SELECT
    id,
    name,
    price
FROM products
WHERE is_deleted = FALSE
AND id = $1;
```

---

## Ambil Banyak Produk

```sql
SELECT
    id,
    name,
    price
FROM products
WHERE is_deleted = FALSE
AND id IN ($1);
```

---

## Tambah Produk

```sql
INSERT INTO products (
    id,
    name,
    price
)
VALUES (
    $1,
    $2,
    $3
);
```

---

## Perbarui Produk

```sql
UPDATE products
SET
    name = $1,
    price = $2
WHERE id = $3;
```

---

## Hapus Produk (Soft Delete)

```sql
UPDATE products
SET is_deleted = TRUE
WHERE id = $1;
```

---

## Membuat Pesanan

```sql
BEGIN;

INSERT INTO orders (
    id,
    email,
    address,
    passcode,
    grand_total
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
);

INSERT INTO order_details (
    id,
    order_id,
    product_id,
    quantity,
    price,
    total
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);

COMMIT;
```

---

## Update Pembayaran

```sql
UPDATE orders
SET
    paid_at = $1,
    paid_bank = $2,
    paid_account_number = $3
WHERE id = $4;
```

---

## Ambil Pesanan

```sql
SELECT
    id,
    email,
    address,
    passcode,
    grand_total,
    paid_at,
    paid_bank,
    paid_account_number
FROM orders
WHERE id = $1;
```

---

## Ambil Detail Pesanan

```sql
SELECT
    id,
    order_id,
    product_id,
    quantity,
    price,
    total
FROM order_details
WHERE order_id = $1;
```