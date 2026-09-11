# Cara memasang setup AI Agent ke ScrapFlow

Paket ini mengikuti struktur repo. Salin isinya ke root repo ScrapFlow **di branch baru**.

```bash
git checkout -b feat/ai-agent-setup
```

## 1. File yang ditambahkan dan diganti

**Baru**
```text
AGENTS.md, CLAUDE.md
spec/requirements.md, spec/architecture.md, spec/ai-agent.md, spec/known-issues.md
.github/workflows/ci.yml
scrap-invoice-backend/.env.example
scrap-invoice-backend/services/invoice_summary.go (+ _test.go)
scrap-invoice-backend/agent/  (seluruh folder)
```

**Mengganti file yang ada** (perubahannya kecil, cek dengan `git diff`)
```text
scrap-invoice-backend/controllers/summary_controller.go   # rumus pindah ke services, JSON sama
scrap-invoice-backend/routes/router.go                    # + route /api/agent/chat jika aktif
scrap-invoice-backend/main.go                             # + agent.Setup setelah migrate
```

## 2. Tahap aman dulu: services saja

Kalau belum siap naik versi Go, kamu bisa memasang **hanya** `services/` dan `controllers/summary_controller.go` dulu. Bagian ini tidak butuh dependency baru dan langsung jalan di Go 1.24:

```bash
cd scrap-invoice-backend
go vet ./... && go test ./services/...
```

Lalu buka halaman summary satu invoice lama dan pastikan angkanya sama persis dengan sebelumnya.

## 3. Naikkan Go dan tambah ADK

ADK Go v2.3.0 membutuhkan Go 1.26.6 atau lebih baru.

```bash
cd scrap-invoice-backend
go mod edit -go=1.26.6
go get google.golang.org/adk/v2@v2.3.0
go mod tidy
go vet ./... && go test ./... && go build ./...
```

Di `Dockerfile`, ganti baris builder:
```dockerfile
FROM golang:1.26-alpine AS builder
```

Bersih-bersih sekalian:
```bash
git rm --cached tmp/main.exe
# perbaiki baris .gitignore yang rusak: "Thumbs.dbdocker-compose logs -f api" → "Thumbs.db"
```

## 4. Jalankan lokal

```bash
cp .env.example .env
# isi: AGENT_ENABLED=true, AGENT_PROVIDER, AGENT_MODEL, AGENT_API_KEY, AGENT_ALLOWED_ROLES=admin
docker compose up -d db     # Postgres lokal saja, di port host 5433
go run .
```

Log harus menampilkan `✅ AI agent aktif`. Tes dengan akun seed admin (hanya ada di `APP_ENV=development`):

```bash
TOKEN=$(curl -s localhost:8080/api/login -H 'Content-Type: application/json' \
  -d '{"email":"admin@scrapflow.com","password":"password123"}' | jq -r .token)

curl -s localhost:8080/api/agent/chat -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"message":"Daftar item kategori Tembaga dan harganya sekarang"}' | jq
```

Kirim `session_id` dari response pertama untuk melanjutkan percakapan yang sama.

## 5. Eval sebelum dipakai siapa pun

```bash
docker exec scrap_db createdb -U postgres scrapflow_eval   # database terpisah, isinya akan dihapus
EVAL_DATABASE_DSN="host=localhost user=postgres password=postgres dbname=scrapflow_eval port=5433 sslmode=disable" \
  go test -tags eval -run TestAgentEvals -v ./agent/
```

Semua kasus harus lulus 3 kali berturut-turut dengan model yang dipilih. Catat model di tabel `spec/ai-agent.md`.

## 6. Yang sudah dan belum diverifikasi

**Sudah**
- `services` dan `agent/agentutil`: dikompilasi dan dites. Rumus di `services` dibandingkan dengan algoritma lama pada 2.000 set data acak dan hasilnya identik.
- API ADK Go v2.3.0 dan genai v1.69.0 yang dipakai dicocokkan langsung dengan source code-nya.
- Package `agent` (termasuk `eval_test.go`) lolos type-check terhadap stub yang meniru API tersebut.

**Belum** (tidak bisa dijalankan di lingkungan pembuatan paket ini)
- `go build` penuh dengan dependency ADK asli. Jalankan langkah 3; jika ada error kompilasi, kemungkinan besar hanya penyesuaian kecil nama tipe.
- Panggilan ke model sungguhan dan query ke Postgres sungguhan. Itu tugas eval di langkah 5.
