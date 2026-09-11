# AGENTS.md — ScrapFlow

Instruksi untuk coding agent (Claude Code, Codex, dan sejenisnya) yang bekerja di repo ini.

ScrapFlow adalah **sistem produksi** yang dipakai operasional perdagangan scrap. Angka berat dan rupiah di sini adalah uang nyata. Kesalahan kecil di perhitungan lebih mahal daripada fitur yang telat.

## Sumber kebenaran

Baca sebelum mengubah perilaku apa pun:

- `spec/requirements.md` — aturan bisnis (tipe timbangan, rumus, penomoran, peran user)
- `spec/architecture.md` — struktur kode dan batas antar lapisan
- `spec/ai-agent.md` — ruang lingkup dan guardrail AI agent di dalam produk
- `spec/known-issues.md` — bug yang sudah diketahui beserta status keputusannya

Aturan: jangan mengubah perilaku di luar spec. Kalau implementasi menuntut perubahan aturan, perbarui spec dulu dan tandai sebagai **usulan** sampai pemilik repo menyetujui.

## Cara kerja

1. Pahami dulu. Baca kode yang disentuh dan telusuri alurnya dari route sampai database sebelum menulis apa pun.
2. Cek apakah memang perlu dibangun, apakah sudah ada di codebase, dan apakah standard library sudah cukup.
3. Tulis diff terkecil yang benar. Tanpa abstraksi, dependency, atau file baru yang tidak diminta.
4. Perbaiki akar masalah. Kalau mengubah fungsi bersama, cek semua pemanggilnya.
5. Penyederhanaan yang punya batas nyata (in-memory, satu instance, O(n²)) diberi komentar `ponytail:` yang menyebut batas dan jalur upgrade-nya.

Tidak boleh dihemat: validasi input di batas kepercayaan, penanganan error yang mencegah kehilangan data, keamanan, dan test untuk logika uang.

## Aturan keras ScrapFlow

- **Perhitungan berat dan harga hanya di `scrap-invoice-backend/services`.** Controller, PDF, dan AI agent memanggil fungsi di sana, tidak menghitung sendiri.
- **Setiap perubahan di `services` wajib punya table-driven test** (package `testing` bawaan). Test yang mendokumentasikan perilaku saat ini hanya boleh diubah bersamaan dengan spec.
- **Jangan menjumlahkan angka antar tipe timbangan.** FI berisi salinan TL.
- **Perubahan skema database** dicatat di spec dan PR. Jangan menghapus atau mengganti nama kolom tanpa rencana migrasi data.
- **Jangan pernah memakai kredensial production.** Kerja hanya dengan Postgres lokal (`docker-compose.yml`) dan data seed.
- **AI agent di `scrap-invoice-backend/agent` read-only.** Jangan menambah tool yang menulis data tanpa perubahan `spec/ai-agent.md` yang disetujui.
- Jangan commit `.env`, binary (`tmp/`, `*.exe`), atau log.

## Perintah

```bash
# Backend
cd scrap-invoice-backend
go vet ./...
go test ./...
go test -tags eval -run TestAgentEvals -v ./agent/   # butuh API key & DB eval, lihat spec/ai-agent.md

# Frontend
cd scrap-invoice-frontend
npm ci
npm run build   # vue-tsc + vite build, wajib lulus
```

Sebelum menyatakan tugas selesai: `go vet`, `go test`, dan `npm run build` harus lulus, dan jelaskan di ringkasan apa yang TIDAK kamu verifikasi.
