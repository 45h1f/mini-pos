# Desktop POS — Mini App Project Plan (v1)

## Goal
Build a cross-platform desktop mini POS (Windows, Linux, macOS) using Go + Wails v2 + Vue 3.
Fully offline. Sell products, print receipts, back up data. Sync and cloud features are v2.

## Tech Stack
- Go 1.24+
- Wails v2 (stable — do not use v3 alpha)
- Vue 3 + TypeScript (strict mode)
- Vite
- SQLite via **GORM + `glebarez/sqlite`** (pure Go driver — no CGO, clean cross-platform builds)
- Pinia
- Vue Router

> **Why `glebarez/sqlite`:** `mattn/go-sqlite3` requires CGO and a C toolchain per target OS. The pure-Go driver removes that entirely.

## Folder Structure
```text
pos-mini/
├── backend/
│   ├── app/            # Wails bindings (thin — no business logic)
│   ├── database/       # connection, AutoMigrate
│   ├── models/
│   ├── repository/
│   ├── services/
│   ├── printer/
│   └── utils/          # logger, errors, backup
├── frontend/
│   └── src/
│       ├── components/
│       ├── pages/
│       ├── layouts/
│       ├── router/
│       ├── stores/
│       ├── composables/
│       ├── types/
│       └── assets/
├── main.go
├── data/               # sqlite file, backups (user data dir at runtime)
├── logs/
└── build/
```

## Milestones (Mini App Scope)

### Phase 1 — Setup & Shell
- Install Go + Wails, scaffold project, init Git
- Verify `wails build` on at least Windows + Linux
- App shell: sidebar, header, router, dark/light theme toggle
- Native dialogs + toast notifications
- Structured logger writing to `logs/` (slog)

**Done when:** app builds and runs on 2+ platforms with working navigation.

### Phase 2 — Database & Core CRUD
Tables (GORM `AutoMigrate` — no separate migration tool for v1):
- `categories` — id, name
- `products` — id, category_id, name, sku, barcode, price, stock, is_active
- `customers` — id, name, phone
- `sales` — id, invoice_no, customer_id (nullable), subtotal, tax, discount, total, paid, change, payment_method, created_at
- `sale_items` — id, sale_id, product_id, name_snapshot, price_snapshot, qty, line_total
- `settings` — key/value

> All monetary values stored as **integer cents/paisa** — never float.
> `sale_items` snapshots name/price so old receipts survive product edits.

Repository + Service layers for each entity. Wails app layer only forwards calls.

**Done when:** full CRUD for products/categories/customers via Wails bindings, with unit tests on services.

### Phase 3 — POS Sale Screen (the core)
- Product grid with search (name/SKU/barcode)
- Cart: add/remove, qty edit, line discount
- Barcode via keyboard-wedge scanner (scanner = keyboard input, catch Enter-terminated bursts on a hidden focused input)
- Checkout: payment method (cash/other), amount tendered, change calculation
- Save sale + items in one transaction, decrement stock
- Sales list page + sale detail view

**Done when:** a complete sale can be scanned, checked out, and viewed in history.

### Phase 4 — Receipt Printing
- HTML receipt template → print via system print dialog (simplest reliable v1)
- Optional: PDF export of receipt
- Printer/paper-size setting stored in `settings`

> ESC/POS raw thermal printing is v2 — system print dialog covers v1.

### Phase 5 — Settings & Data Safety
- Settings page: store info, currency symbol, tax %, theme, receipt footer
- Backup: copy SQLite file to user-chosen location (native save dialog)
- Restore: import backup with confirmation
- CSV export/import for products

### Phase 6 — Polish & Release
- Keyboard shortcuts (F2 new sale, F4 checkout, / focus search)
- Pagination on lists, form validation everywhere
- Dashboard: today's sales count/total, low-stock list
- Simple reports: daily/date-range sales summary
- Final builds for Windows/Linux/macOS + README with build steps

## Definition of Done (v1)
- Builds on Windows/Linux/macOS
- Complete sale flow: search/scan → cart → checkout → receipt → history
- Stock decrements on sale
- Backup/restore works
- 100% offline, no network calls
- Services covered by unit tests

---

## v2 Backlog (deferred — do NOT build in v1)
Kept from the original plan, in build order:

1. **Laravel REST API** — login (Sanctum), products, customers, sales endpoints
2. **Offline sync engine** — sync queue table; each record: uuid, updated_at, deleted_at, device_id, version, sync_status. Conflict rule: last-write-wins by `version`, log conflicts.
3. **Background worker** — connectivity check + queue upload every **30–60s with exponential backoff** (not 5s)
4. **Auto-update** — requires code signing on macOS/Windows; plan certificates first
5. ESC/POS raw thermal printing, cash drawer kick
6. Inventory/Purchase/Supplier, multi-store, promotions, loyalty, SaaS mode

---

## AI Agent Instructions

### Coding Rules
- SOLID, Clean Architecture, Repository + Service pattern, DI via constructors
- No business logic in Wails app layer or Vue components
- Small functions, strong typing (Go + TS strict), meaningful names
- Every error handled and logged — no silent failures
- Money as integers; UUIDs for all primary keys (future sync-proofing)

### Git Workflow
- One feature per branch, one commit per logical change
- Conventional commits (`feat:`, `fix:`, `refactor:`)

### Every task must include
- Goal, Acceptance Criteria, Files to create, Files to modify, Tests, Docs

### Code Standards
- `gofmt` + `go vet`, ESLint + Prettier
- TypeScript strict mode, no `any`
- No duplicated code
