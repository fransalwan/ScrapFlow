# Architecture — ScrapFlow

## Gambaran

```text
Vue 3 + Pinia (Vercel)  ──HTTP/JWT──►  Go + Gin (satu binary)  ──GORM──►  PostgreSQL
                                          │
                                          └─ agent/ (ADK Go) ──► LLM API (Gemini / OpenAI-compatible)
```

Satu backend, satu database. AI agent **bukan** service terpisah: ia hidup di dalam binary Go supaya memakai autentikasi, koneksi database, dan logika perhitungan yang sama.

## Struktur backend

```text
scrap-invoice-backend/
├── main.go            # boot: DB, migrate, seed (dev), agent.Setup, router
├── config/            # koneksi database (global config.DB)
├── middleware/        # AuthMiddleware (JWT), RoleMiddleware
├── routes/            # daftar route
├── controllers/       # HTTP handler: bind input, panggil DB/services, tulis JSON
├── services/          # LOGIKA BISNIS MURNI + test. Satu-satunya tempat rumus uang.
├── dto/, models/      # bentuk input dan tabel
└── agent/             # AI analyst read-only
    ├── agent.go       # Setup: model, llmagent, session ADK di Postgres, runner
    ├── config.go      # env AGENT_*
    ├── tools.go       # tools read-only (GORM + services)
    ├── handler.go     # POST /api/agent/chat + tabel agent_logs
    ├── prompts/system.md
    ├── agentutil/     # helper murni + test: validasi, format angka, rate limit, laporan
    ├── evals/cases.json
    └── eval_test.go   # build tag `eval`, memanggil model sungguhan
```

## Aturan lapisan

1. **`services` tidak mengimpor gin, config, atau agent.** Menerima data, mengembalikan hasil. Mudah dites.
2. **Controller dan agent sama-sama memanggil `services`** untuk angka. Kalau halaman summary dan agent menampilkan angka berbeda, itu bug.
3. **`agentutil` tidak mengimpor ADK**, supaya logika guardrail dites tanpa model dan tanpa jaringan.
4. **Agent tidak memanggil HTTP API backend-nya sendiri.** Tools membaca lewat GORM langsung, dan identitas user diambil dari JWT di handler lalu diteruskan sebagai `userID` ADK.
5. Kegagalan agent (config salah, LLM down) **tidak boleh** menghentikan server atau fitur invoice.

## Data milik agent

| Tabel | Pemilik | Isi |
|---|---|---|
| `agent_logs` | ScrapFlow | pesan, jawaban, tool + argumen, error, latensi, token |
| `sessions`, `events`, `app_states`, `user_states` | ADK | riwayat percakapan, di-namespace per user |

Nama tabel ADK generik. Hindari membuat tabel ScrapFlow dengan nama yang sama.

## Dependency dan versi

- Go **1.26.6+** diwajibkan oleh `google.golang.org/adk/v2` v2.3.0 (sebelumnya repo memakai 1.24.4). Update `go.mod` dan image Docker.
- ADK Go v1 (`google.golang.org/adk`, Go 1.25) adalah alternatif jika belum bisa naik ke Go 1.26, dengan beberapa perbedaan API.

## Arah berikutnya (belum dikerjakan)

- Snapshot harga per timbangan (known-issues #2).
- Migrasi skema eksplisit menggantikan `AutoMigrate` di production.
- Audit log perubahan data bisnis (siapa mengubah apa, nilai lama dan baru).
