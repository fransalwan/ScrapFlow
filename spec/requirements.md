# Requirements — ScrapFlow

Dokumen ini adalah sumber kebenaran aturan bisnis. Isi awalnya disusun dengan **membaca kode v1.0**, jadi menggambarkan apa yang *dilakukan* sistem, belum tentu apa yang *seharusnya*. Setiap aturan diberi status:

- ✅ **Terverifikasi** — sudah dikonfirmasi pemilik bisnis
- 🔍 **Dari kode** — perilaku saat ini, belum dikonfirmasi
- ❓ **Terbuka** — perlu keputusan sebelum kode diubah

Ubah status menjadi ✅ hanya setelah dicek ke operasional nyata atau owner.

---

## 1. Peran user

| Role | Status | Perilaku saat ini |
|---|---|---|
| admin, staff, operator | 🔍 | Semua role bisa membaca, membuat, mengubah, dan menghapus invoice, timbangan, customer, kategori, dan item. |
| Dashboard count | 🔍 | Maksud komentar kode: hanya staff & operator. Kenyataannya tidak ada role yang bisa akses (lihat known-issues #8). |

❓ Siapa yang boleh menghapus invoice dan mengubah harga master item?
❓ Role mana yang boleh memakai AI agent? (dikontrol lewat `AGENT_ALLOWED_ROLES`)

## 2. Entitas

- **Customer** 🔍 — nama, telepon, email (unik), alamat, tier.
- **Kategori item** 🔍 — jenis bahan: Besi, Tembaga, Aluminium, Plastik, dan seterusnya.
- **Item** 🔍 — nama, kategori, harga per kg (harga master, bisa diubah kapan saja).
- **Invoice** 🔍 — nomor, customer, tanggal invoice, status (default `draft`), metode pembayaran, catatan.
- **Timbangan (scale detail)** 🔍 — satu kali timbang: invoice, item, berat, berat alas, foto, tipe timbangan.

## 3. Tipe timbangan

| Kode | Nama | Status |
|---|---|---|
| TL | Timbang Lapangan | ✅ |
| TG | Timbang Gudang | ✅ |
| TS | Timbang Sortir | ✅ |
| FI | Full Invoice | ✅ |

Aturan terkait:

- 🔍 Saat timbangan **TL dibuat**, sistem otomatis membuat salinan identik bertipe **FI**.
- 🔍 Timbangan FI juga bisa diinput langsung dari halaman FI.
- 🔍 Halaman summary invoice dibuka dari tab FI, jadi FI tampaknya dasar penagihan.
- ❓ Apakah FI memang dasar nilai yang ditagihkan ke customer?
- ❓ Kalau TL diedit atau dihapus, apakah salinan FI harus ikut berubah? (saat ini tidak, known-issues #4)
- ❓ Apa hubungan TL, TG, dan TS? Misalnya apakah selisih berat TL ke TG dianggap susut dan perlu dilaporkan?
- 🔍 Backend belum menolak tipe selain FI/TL/TG/TS.

**Aturan pelaporan:** angka antar tipe timbangan tidak boleh dijumlahkan, karena FI berisi salinan TL.

## 4. Perhitungan berat dan nilai

Diimplementasikan di `scrap-invoice-backend/services/invoice_summary.go`.

- 🔍 Timbangan dikelompokkan per tipe, lalu per item.
- 🔍 `subtotal item = total berat timbang × harga master item saat ini`
- 🔍 `grand total per tipe = jumlah subtotal item pada tipe itu`
- ❓ **Berat alas tidak mengurangi berat yang dihargai**, padahal UI menampilkan alas sebagai pengurang ("- 5 kg (alas)"). Mana yang benar: harga dari berat timbang atau dari berat bersih (timbang − alas)? (known-issues #1)
- ❓ **Harga tidak disimpan per transaksi.** Mengubah harga master mengubah nilai semua invoice lama. Apakah harga harus dikunci saat invoice dibuat atau difinalkan? (known-issues #2)
- ❓ Pembulatan rupiah: ke rupiah terdekat, ke ratusan, atau tanpa pembulatan?

## 5. Penomoran invoice

- 🔍 Format `INV-YYYYMMDD-NNNN`, tanggal dari `invoice_date`, `NNNN` = jumlah invoice pada tanggal itu + 1.
- ❓ Nomor harus unik selamanya, termasuk setelah invoice dihapus? (saat ini bisa duplikat, known-issues #3)

## 6. Status invoice

- 🔍 Default `draft`. Nilai lain dikirim bebas dari frontend, tidak divalidasi.
- ❓ Daftar status yang sah dan alurnya (misalnya draft → final → dibayar)? Apakah invoice final boleh diedit?

## 7. Validasi input

- 🔍 Berat wajib diisi dan > 0 di frontend. Backend: update menolak berat negatif, create belum.
- ❓ Rentang berat alas yang wajar per item/kategori (untuk anomaly check di masa depan)?
- ❓ Apakah berat alas boleh lebih besar dari berat timbang?
