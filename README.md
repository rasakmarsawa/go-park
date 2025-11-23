# 🚗 Parking Lot System — Golang CLI

Sistem parkir berbasis Command Line Interface (CLI) yang membaca perintah dari file dan mengeksekusi operasi seperti membuat parking lot, parkir mobil, menghapus mobil, serta menampilkan status.
Dibuat menggunakan **OOP-style Golang**.

---

## 📂 Project Structure

```
.
├── commands.txt
├── filehandler.go
├── go.mod
├── go-park
├── main.go
├── parkinglot.go
├── result.txt
└── structure.txt

1 directory, 8 files
```

---

# 🧠 Features

✔ Membuat parking lot dengan kapasitas tertentu
✔ Parkir mobil (slot paling dekat yang kosong)
✔ Leave/unpark mobil dan menghitung biaya parkir
✔ Menampilkan status slot parkir
✔ Membaca perintah dari file

---

# 📝 Commands Supported

| Command                      | Description                              |
| ---------------------------- | ---------------------------------------- |
| `create_parking_lot {n}`     | Membuat parking lot dengan n slot        |
| `park {car_number}`          | Memarkir mobil pada slot kosong terdekat |
| `leave {car_number} {hours}` | Mengeluarkan mobil + menghitung biaya    |
| `status`                     | Menampilkan status slot parkir           |

---

# ▶️ How to Run

1. Clone repo:

```sh
git clone https://github.com/yourname/parking-app
cd parking-app
```

2. Init module (kalau belum):

```sh
go mod tidy
```

3. Jalankan aplikasi:

```sh
go run . commands.txt
```

---

# 🧮 Parking Charge Policy

* **$10 untuk 2 jam pertama**
* **+$10 untuk setiap jam berikutnya**
  
---

# 🛠️ Tech Stack

* **Golang**
* **Go Modules**
* **Standard Library Only (I/O, bufio, os, strings, fmt)**

---

# 📄 License

MIT License.
