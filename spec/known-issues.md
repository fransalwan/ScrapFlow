# Known Issues — hasil audit kode v1.0 (commit 22f50e0)

Ditemukan dengan membaca kode, **belum dicek terhadap data production**. Sebelum memperbaiki yang ditandai "butuh keputusan", putuskan dulu aturannya di `requirements.md`.

Prioritas: 🔴 bisa menyebabkan angka uang salah atau celah keamanan · 🟠 data tidak konsisten · 🟡 kualitas/operasional

---

### 1. 🔴 Berat alas tidak mengurangi nilai (butuh keputusan)
`controllers/summary_controller.go:54-55` (sekarang di `services/invoice_summary.go`)

Subtotal = `Weight × harga`. `AlasWeight` disimpan dan ditampilkan di UI sebagai pengurang (`ScaleDetail.vue:47`), tapi tidak dipakai di perhitungan. Kalau `weight` yang diinput adalah berat kotor, nilai invoice lebih besar dari seharusnya.

**Cek dulu:** staff menginput berat kotor atau berat bersih di kolom berat?

### 2. 🔴 Harga invoice lama berubah saat harga master diubah (butuh keputusan)
`controllers/summary_controller.go:39-43`

Summary membaca `item.PricePerKg` saat halaman dibuka. Tabel `summaries` punya kolom `sub_total_price` tapi tidak pernah diisi. Akibatnya menaikkan harga Besi Beton hari ini mengubah nilai semua invoice Besi Beton bulan lalu.

**Arah perbaikan:** simpan `price_per_kg` di setiap scale detail saat dibuat (snapshot), dan hitung dari snapshot itu.

### 3. 🔴 Nomor invoice bisa duplikat
`controllers/invoice_controller.go:100-105`

Nomor = jumlah invoice di tanggal itu + 1. Duplikat terjadi jika (a) invoice di tengah dihapus: ada 0001–0003, hapus 0002, invoice baru dapat 0003 lagi; atau (b) dua admin membuat invoice di saat bersamaan. Kolom `invoice_number` tidak punya unique index, jadi database tidak menolak.

**Arah perbaikan:** unique index + ambil nomor terakhir (bukan count) di dalam transaksi, dengan retry saat bentrok.

### 4. 🟠 Salinan FI tidak ikut berubah saat TL diedit atau dihapus (butuh keputusan)
`controllers/scale_detail_controller.go:80-92, 191, 260`

Membuat TL otomatis membuat FI. Mengedit atau menghapus TL tidak menyentuh FI-nya, dan tidak ada kolom yang menghubungkan keduanya. Selain itu salinan FI dibuat **sebelum** TL, tanpa transaksi: kalau simpan TL gagal, FI yatim tertinggal.

### 5. 🔴 JWT secret punya fallback default yang ada di repo publik
`utils/jwt.go:21-25`

Jika `JWT_SECRET` lupa diset di server, token ditandatangani dengan string yang bisa dibaca siapa pun di GitHub, sehingga siapa pun bisa membuat token admin.

**Perbaikan:** jika `APP_ENV=production` dan `JWT_SECRET` kosong, `log.Fatal`. Cek juga env production sekarang.

### 6. 🟠 Total invoice dan PDF menampilkan 0
`models/invoice.go`, `controllers/print_pdf_controller.go`, `pages/InvoiceList.vue:52`

`Invoice.TotalWeight` dan `TotalPrice` tidak pernah diisi, jadi daftar invoice menampilkan Rp 0. PDF membaca `invoice.Summaries` yang juga tidak pernah diisi, jadi tabel item di PDF kosong.

**Arah perbaikan:** PDF memakai `services.SummarizeScaleDetails`, sama dengan halaman summary.

### 7. 🟠 Response create invoice mengembalikan invoice yang salah
`controllers/invoice_controller.go:119`

`First(&fullInvoice, invoice.CustomerID)` mencari invoice dengan **ID = ID customer**. Seharusnya `invoice.ID`.

### 8. 🟡 Endpoint dashboard selalu 403
`routes/router.go:67`

`RoleMiddleware()` dipanggil tanpa daftar role, jadi tidak ada role yang lolos. Komentar menyebut staff & operator.

### 9. 🟠 Validasi timbangan belum ada di backend
`dto/scale_detailDTO.go`

`scale_type` menerima string apa pun, berat negatif lolos saat create, berat alas boleh negatif atau melebihi berat.

### 10. 🟡 Lain-lain
- Semua role bisa menghapus invoice; belum ada audit log perubahan data.
- `AutoMigrate` jalan otomatis di production setiap start.
- `GetInvoices` mengambil semua invoice tanpa paginasi.
- `CreateInvoice` mencatat seluruh raw body ke log.
- `created_by` tidak pernah diisi.
- `docker-compose.yml`: service `api` memakai `DB_PORT=5433`, padahal di dalam jaringan Docker Postgres mendengarkan di 5432, dan `JWT_SECRET` contoh tertulis di file.
- `tmp/main.exe` ter-commit; baris `.gitignore` rusak: `Thumbs.dbdocker-compose logs -f api`.
- Urutan item di summary acak — **sudah diperbaiki** lewat `services`.
