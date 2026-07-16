# ⚖️ ScrapFlow - Invoice & Scale Management System (v1.0.0)

Sistem manajemen invoice dan penimbangan (scale detail) yang dirancang khusus untuk operasional perdagangan bahan baku/scrap. Aplikasi ini menyediakan fitur lengkap mulai dari manajemen customer, pembuatan invoice otomatis, pencatatan detail timbangan per kategori, hingga rekapitulasi total harga dan berat secara real-time.

---

## 🚀 Fitur Utama (v1.0)

- **Manajemen Invoice**: Pembuatan, pembaruan, dan penghapusan invoice dengan penomoran otomatis berbasis tanggal (Format: `INV-YYYYMMDD-XXXX`).
- **Detail Timbangan (Scale Details)**: Pencatatan berat bersih dan berat alas per item, dikategorikan berdasarkan tipe timbangan (FI, TL, TG, TS).
- **Keamanan Data**: Validasi kepemilikan (ownership validation) pada endpoint update/delete untuk mencegah manipulasi data lintas invoice.
- **Rekapitulasi Otomatis (Summary)**: Perhitungan otomatis total berat (`total_weight`) dan subtotal harga (`sub_total_price`) per item, dilengkapi Grand Total per kategori timbangan.
- **Manajemen Master Data**: CRUD untuk Data Customer dan Data Item (Barang) lengkap dengan kategori dan harga per kg.
- **Autentikasi Aman**: Menggunakan JWT (JSON Web Token) dengan interceptor Axios di frontend dan middleware di backend.

---

## 🛠️ Tech Stack

### Backend
- **Bahasa**: Go (Golang) 1.21+
- **Framework**: Gin Web Framework
- **ORM**: GORM
- **Database**: PostgreSQL (via Docker)
- **Auth**: JWT (golang-jwt/jwt)

### Frontend
- **Framework**: Vue 3 (Composition API, `<script setup>`)
- **Bahasa**: TypeScript
- **State Management**: Pinia
- **Routing**: Vue Router
- **Styling**: Tailwind CSS
- **HTTP Client**: Axios (dengan Custom Interceptors)
- **UI/UX**: Vue Toastification, SweetAlert2, Lucide Icons

---

## 📋 Prasyarat (Prerequisites)

Pastikan mesin kamu telah terinstal:
- [Go](https://go.dev/doc/install) (v1.21 atau lebih baru)
- [Node.js](https://nodejs.org/) (v18 atau lebih baru) & npm/pnpm
- [Docker](https://www.docker.com/) & Docker Compose (untuk database lokal)
- [PostgreSQL](https://www.postgresql.org/) (opsional, jika tidak menggunakan Docker)

---
