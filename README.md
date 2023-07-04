# GOPasted

Backend untuk project code share (seperti Pastebin) dengan bahasa pemrograman GO dan MongoDB sebagai databasenya. Program ini saya buat untuk tujuan belajar bahasa pemrograman GO dan MongoDB.

- Frontend: (https://github.com/hsnfirdaus/react-pasted)[https://github.com/hsnfirdaus/react-pasted]

## Menjalankan Project

1. Clone Repositori ini:

```shell
git clone https://github.com/hsnfirdaus/gopasted.git
```

2. Install GO module dependensi:

```shell
go mod download
```

3. Buat file .env berisikan MongoDB connection string dan nama database:

```env
CONNECTION=mongodb://127.0.0.1:27017
DB_NAME=example
```

4. Jalankan development server:

```shell
go run .
```

# Hak Cipta

&copy; 2021 Muhammad Hasan Firdaus
