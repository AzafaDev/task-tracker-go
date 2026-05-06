Berikut README.md yang lengkap dan siap pakai:

```markdown
# Task Tracker CLI

Aplikasi command-line sederhana untuk manajemen tugas (to-do list) yang ditulis dalam bahasa Go. Data disimpan secara lokal dalam file JSON.

Proyek ini dibuat sebagai bagian dari pembelajaran Go dari nol, mencakup konsep dasar seperti operasi file, parsing JSON, argumen command line, dan struktur proyek yang rapi.

---

## Daftar Isi

- [Fitur](#fitur)
- [Instalasi](#instalasi)
- [Cara Pakai](#cara-pakai)
- [Contoh Output](#contoh-output)
- [Struktur Proyek](#struktur-proyek)
- [Teknologi](#teknologi)
- [Pembelajaran](#pembelajaran)
- [Rencana Pengembangan](#rencana-pengembangan)
- [Lisensi](#lisensi)

---

## Fitur

- ✅ Tambah tugas baru dengan judul
- ✅ Validasi judul tidak boleh kosong atau hanya spasi
- 📋 Lihat daftar semua tugas dalam bentuk tabel
- 🔍 Filter tugas berdasarkan status (`todo`, `in-progress`, `done`)
- ✏️ Ubah judul tugas berdasarkan ID
- 🗑️ Hapus tugas berdasarkan ID
- 🔄 Ubah status tugas: `mark-in-progress` dan `mark-done`
- 💾 Data tersimpan otomatis di file `tasks.json`
- ⏰ Setiap tugas memiliki timestamp dibuat dan diubah

---

## Instalasi

### Prasyarat
- **Go** versi 1.26 atau lebih baru

### Langkah Instalasi

```bash
# Clone repositori
git clone https://github.com/AzafaDev/task-tracker-go.git

# Masuk ke direktori proyek
cd task-tracker-go

# Build binary
go build -o task-cli .

# (Opsional) Install agar bisa dipanggil dari mana saja
go install .
```

Binary `task-cli` sekarang siap digunakan.

---

## Cara Pakai

### Format Umum
```bash
./task-cli [command] [arguments...]
```

### Perintah yang Tersedia

| Perintah | Argumen | Deskripsi |
|----------|---------|-----------|
| `add` | `<title>` | Menambah tugas baru dengan status `todo` |
| `list` | `[status]` | Menampilkan semua tugas (opsional: filter status) |
| `update` | `<id> <new title>` | Mengubah judul tugas |
| `delete` | `<id>` | Menghapus tugas |
| `mark-in-progress` | `<id>` | Mengubah status jadi `in-progress` |
| `mark-done` | `<id>` | Mengubah status jadi `done` |

### Contoh Lengkap

```bash
# Menambah tugas
./task-cli add "Belajar Go"
./task-cli add "Membaca buku Clean Code"
./task-cli add "Olahraga pagi"

# Melihat semua tugas
./task-cli list

# Filter tugas yang masih todo
./task-cli list todo

# Filter tugas yang sedang dikerjakan
./task-cli list in-progress

# Filter tugas yang sudah selesai
./task-cli list done

# Ubah judul tugas dengan ID 1
./task-cli update 1 "Belajar Go dari dasar sampai mahir"

# Tandai tugas ID 1 sedang dikerjakan
./task-cli mark-in-progress 1

# Tandai tugas ID 2 selesai
./task-cli mark-done 2

# Hapus tugas ID 3
./task-cli delete 3
```

---

## Contoh Output

### Menambah Tugas
```bash
$ ./task-cli add "Belajar Go"
Added task: [1] Belajar Go

$ ./task-cli add "Membaca buku"
Added task: [2] Membaca buku

$ ./task-cli add "Olahraga"
Added task: [3] Olahraga
```

### Menampilkan Daftar Tugas
```bash
$ ./task-cli list
ID  | Status        | Judul
----|---------------|------
1   | todo          | Belajar Go
2   | todo          | Membaca buku
3   | todo          | Olahraga
```

### Filter Tugas
```bash
$ ./task-cli list todo
ID  | Status        | Judul
----|---------------|------
1   | todo          | Belajar Go
2   | todo          | Membaca buku
3   | todo          | Olahraga
```

### Ubah Status Tugas
```bash
$ ./task-cli mark-in-progress 1
Task [1] now is in-progress

$ ./task-cli list
ID  | Status        | Judul
----|---------------|------
1   | in-progress   | Belajar Go
2   | todo          | Membaca buku
3   | todo          | Olahraga
```

### Ubah Judul Tugas
```bash
$ ./task-cli update 1 "Belajar Go dengan proyek"
Task [1] updated successfully
```

### Hapus Tugas
```bash
$ ./task-cli delete 3
Task [3] deleted successfully
```

### Tandai Selesai
```bash
$ ./task-cli mark-done 2
Task [2] is now done
```

---

## Struktur Proyek

```
task-tracker-go/
├── cmd/                    # Logika setiap perintah CLI
│   ├── add.go              # Perintah "add"
│   ├── delete.go           # Perintah "delete"
│   ├── list.go             # Perintah "list"
│   ├── mark.go             # Perintah "mark-in-progress" & "mark-done"
│   └── update.go           # Perintah "update"
├── internal/
│   └── task/               # Package internal untuk model dan penyimpanan
│       ├── model.go        # Definisi struct Task
│       └── store.go        # Fungsi load dan save ke file JSON
├── main.go                 # Entry point aplikasi
├── go.mod                  # Go module definition
├── tasks.json              # File penyimpanan data (muncul setelah "add" pertama)
├── README.md               # Dokumentasi proyek
└── .gitignore              # File yang diabaikan Git
```

### Penjelasan Singkat

- **`cmd/`** — Setiap file menangani satu perintah. Mudah dibaca dan dirawat.
- **`internal/task/`** — Package `internal` tidak bisa diimpor dari luar proyek. Cocok untuk logika bisnis internal.
- **`main.go`** — Hanya parsing argumen command line dan memanggil fungsi di package `cmd`.
- **`tasks.json`** — File JSON yang menyimpan semua data tugas.

---

## Teknologi

Proyek ini hanya menggunakan **Go Standard Library**, tidak ada dependensi eksternal:

| Package | Kegunaan |
|---------|----------|
| `encoding/json` | Serialisasi dan deserialisasi data JSON |
| `fmt` | Output ke terminal |
| `os` | Operasi file dan membaca argumen command line |
| `strconv` | Konversi string ke integer |
| `strings` | Manipulasi string (trim, replace) |
| `time` | Timestamp untuk created_at dan updated_at |

---

## Pembelajaran

Proyek ini cocok untuk pemula yang ingin belajar Go, khususnya:

- Membuat aplikasi CLI (Command Line Interface)
- Membaca dan menulis file
- Serialisasi dan deserialisasi JSON dengan struct tags
- Parsing argumen dari terminal (`os.Args`)
- Error handling dasar di Go
- Manipulasi slice (append, delete, find)
- Struktur proyek Go yang modular (package, internal)
- Validasi input pengguna

---

## Rencana Pengembangan

Fitur yang mungkin ditambahkan di masa depan:

- [ ] Prioritas tugas (low, medium, high)
- [ ] Due date (tenggat waktu)
- [ ] Pencarian tugas berdasarkan kata kunci
- [ ] Export ke CSV
- [ ] Unit test
- [ ] Pewarnaan output terminal (menggunakan package `color`)
- [ ] Mode interaktif
- [ ] Menyimpan data ke SQLite (menggunakan `database/sql`)

---

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).

Bebas digunakan, dimodifikasi, dan didistribusikan untuk keperluan belajar maupun komersial.

---

## Author

**AzafaDev**

- GitHub: [@AzafaDev](https://github.com/AzafaDev)
- Proyek: [task-tracker-go](https://github.com/AzafaDev/task-tracker-go)

---

Dibuat dengan ❤️ dan Go
```

Tinggal salin semua isi di atas ke file `README.md`. Kalau ada bagian yang ingin disesuaikan (misal nama author, link, atau rencana fitur), bilang saja!
