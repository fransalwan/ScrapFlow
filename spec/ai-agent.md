# AI Agent — ScrapFlow Analyst

## Tujuan

Owner atau admin bisa bertanya dalam bahasa sehari-hari, misalnya *"berapa nilai FI PT Baja Jaya bulan ini?"*, dan mendapat jawaban yang **angkanya identik dengan dashboard**, lengkap dengan link untuk verifikasi.

## Bukan tujuan (fase ini)

- Membuat, mengubah, atau menghapus data apa pun.
- Membaca foto atau dokumen.
- Menghitung metrik yang belum didefinisikan di `requirements.md` (misalnya susut TL→TG).
- Menggantikan dashboard. Kalau AI mati, operasional jalan seperti biasa.

## Prinsip

1. **Read-only.** Tidak ada tool yang menulis ke database.
2. **LLM tidak menghitung.** Semua angka dihitung `services`, diformat server (`agentutil`), dan dikirim ke model sebagai teks jadi tanpa angka mentah.
3. **Satu rumus.** Tools memakai `services.SummarizeScaleDetails`, fungsi yang sama dengan halaman summary.
4. **Transparan soal dasar harga.** Setiap hasil rupiah membawa `price_basis`, karena aturan harga masih terbuka (known-issues #1, #2).
5. **Bisa diaudit.** Setiap pesan dicatat di `agent_logs` dengan tool dan argumennya.
6. **Gagal dengan aman.** Error tool dikirim ke model sebagai pesan jelas; agent tidak boleh mengulang-ulang dengan parameter tebakan.

## Tools

| Tool | Input | Output | Catatan |
|---|---|---|---|
| `list_customers` | query?, limit? | id, nama, tier | Untuk mengubah nama menjadi ID |
| `list_items` | query? | nama, kategori, harga master saat ini | Selalu disertai catatan bahwa ini harga sekarang |
| `list_invoices` | date_from, date_to, customer_id?, status?, limit? | nomor, tanggal, customer, status, link | Tanpa angka uang |
| `get_invoice_summary` | invoice_number | per tipe timbangan: total berat, alas, bersih, nilai, item | Mendeteksi nomor duplikat (known-issues #3) |
| `get_period_report` | date_from, date_to, scale_type?, customer_id? | per tipe: jumlah invoice, total, 15 item teratas | Maks 366 hari dan 20.000 baris timbangan |

Batas: rentang tanggal ≤ 366 hari, daftar ≤ 50 baris, pesan ≤ 1.000 karakter. Semua dites di `agentutil`.

## Keamanan

- Endpoint `POST /api/agent/chat` di belakang `AuthMiddleware` + `RoleMiddleware(AGENT_ALLOWED_ROLES)`.
- `userID` ADK diambil dari JWT, bukan dari body. Session ID dibuat server dan di-namespace per user.
- Rate limit per user per jam (in-memory; lihat komentar `ponytail:` jika backend lebih dari satu instance).
- Nama customer, item, dan catatan invoice diperlakukan sebagai data. Prompt melarang mengikuti instruksi di dalamnya, dan ada eval untuk itu.
- **Privasi:** isi pertanyaan dan hasil tool (nama customer, harga, berat) dikirim ke provider LLM. Pilih provider dan paket yang kebijakan datanya disetujui owner. Riwayat percakapan tersimpan di tabel ADK; tentukan masa simpan.

## Konfigurasi

Lihat `scrap-invoice-backend/.env.example`. `AGENT_MODEL` sengaja tanpa default: pilih model terbaru yang mendukung function calling saat deploy, dan catat pilihannya di sini.

| Tanggal | Provider | Model | Alasan |
|---|---|---|---|
| _isi saat deploy_ | | | |

## Evaluasi

`agent/evals/cases.json` + `agent/eval_test.go` (build tag `eval`). Menjalankan pertanyaan nyata terhadap database eval dengan data yang angkanya diketahui, lalu memeriksa tool yang dipanggil dan isi jawaban.

Wajib dijalankan saat: mengganti model, mengubah prompt, menambah atau mengubah tool, dan sebelum rilis. Tambahkan kasus baru setiap kali menemukan jawaban salah di `agent_logs`.

Target sebelum dibuka ke owner: **semua kasus lulus 3 kali berturut-turut**.

## Fase

**Fase 0 — Fondasi (sebelum agent diaktifkan di production)**
- [ ] Putuskan known-issues #1 (alas) dan #2 (snapshot harga) bersama owner
- [ ] Perbaiki known-issues #5 (JWT secret) dan #3 (nomor invoice)
- [ ] Tandai aturan di `requirements.md` bagian 3–4 sebagai ✅
- [x] Rumus summary dipindah ke `services` dengan test

**Fase 1 — Analyst read-only (package ini)**
- [ ] Naikkan Go ke 1.26.6+, tambah dependency ADK, `go build` lulus
- [ ] Eval lulus 3× dengan model terpilih
- [ ] Aktifkan di staging untuk 1 role (`admin`), pantau `agent_logs` 1–2 minggu
- [ ] Panel chat di frontend (Vue) yang memanggil `/api/agent/chat`
- [ ] Buka ke owner

**Fase 2 — Laporan yang didefinisikan bisnis**
- [ ] Definisikan susut TL→TG→TS di `requirements.md`, implementasi di `services` + test, lalu jadikan tool
- [ ] Anomaly check berbasis aturan (bukan LLM) dari rentang wajar yang disepakati

**Fase 3 — Tool yang menulis data (hanya jika benar-benar dibutuhkan)**

Syarat minimal sebelum menambah tool tulis: audit log perubahan data, snapshot harga, validasi backend lengkap. Desainnya:
- Tool hanya **mengusulkan** (misalnya membuat draft), memakai `RequireConfirmation` ADK sehingga user harus menyetujui di UI.
- Harga selalu diambil backend dari master, tidak pernah dari argumen model.
- Jika validasi gagal, agent berhenti dan menyerahkan ke manusia. Tidak ada "self-correction" untuk angka.
