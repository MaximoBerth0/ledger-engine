# Ledger Engine

Ledger Engine is a double-entry accounting backend designed as the financial core of a SaaS platform. Every monetary movement is recorded as a pair of immutable entries — a debit and a credit — ensuring the system is always balanced, fully auditable, and impossible to corrupt through partial writes. Money never appears or disappears, it only moves between accounts.

## What problem it solves

Most financial bugs come from the same root causes: balances updated in place,
no audit trail, duplicate operations processed twice, and business logic scattered
across the codebase.

Ledger Engine is designed to eliminate all of that:

- **Double-entry accounting** — every movement generates two entries (debit + credit),
  so the books are always self-verifiable.
- **Immutable ledger** — nothing is ever updated or deleted. Mistakes are corrected
  with a compensating transaction, leaving a complete audit trail.
- **Idempotency** — duplicate requests (retries, network failures) are detected and
  ignored, so no operation is ever processed twice.
- **Explicit money types** — currency and amounts are first-class types, not raw
  floats. Precision errors and currency mismatches are caught at compile time.

---

## Technical Highlights

**Append-only ledger** ——
No UPDATE operations on financial records. Cancellations are modeled as reversal transactions that compensate the original, keeping the full history intact.

**Pessimistic locking** ——
Concurrent transfers use `SELECT FOR UPDATE` to lock account rows before reading balances, preventing race conditions in high-throughput scenarios.

**Mathematical integrity** ——
A database-level constraint enforces that the sum of all entries within a transaction equals exactly zero, guaranteeing the double-entry invariant at the persistence layer.

**Idempotent operations** ——
A middleware layer intercepts duplicate requests by checking the idempotency_key before any business logic runs. Repeated requests return the original response without side effects.

**Domain-driven design** ——
Built around Account, Transaction, and Entry aggregates with rich value objects, clean use case boundaries, and repository interfaces that decouple domain logic from persistence.

**Derived balances** ——
Account balances are never stored — they are always computed from the entry log. The ledger entries are the source of truth, the balance is a projection of them.

---

## Base Architecture
The project follows a 4-layer architecture where each layer has a single responsibility
and dependencies only flow inward, never outward.

```
ledger-engine/
├── cmd/
│   └── main.go
│
├── db/
│   └── migrations/
│
├── internal/
│   ├── domain/           # Types, structs only — no logic
│   │   ├── account.go
│   │   ├── errors.go
│   │   ├── entry.go
│   │   ├── transaction.go
│   │   └── repositories/   # interfaces 
│   │
│   ├── handler/
│   │   ├── middleware/
│   │   │   ├── idempotency.go
│   │   │   └── logger.go
│   │   ├── account.go
│   │   ├── balance.go
│   │   ├── movement.go
│   │   └── server.go     # router
│   │
│   ├── service/          # Business logic (all use cases)
│   │   ├── account.go
│   │   ├── balance.go
│   │   └── movement.go
│   │
│   └── store/            # DB queries (one file per entity)
│       ├── account.go
│       ├── entry.go
│       └── transaction.go
│
├── pkg/              
│   ├── apperrors/
│   └── logger/
│
├── go.mod
└── go.sum
```

---

### domain/
The core of the system. Defines structs, repository interfaces, value types, and domain
errors. No business logic lives here, only the shape of the data and the contracts
other layers must fulfill. The only allowed imports are stdlib packages.

### service/
Where all business logic lives. Each file groups the use cases for one domain area
(account, movement, balance). Depends on domain interfaces, never on concrete
implementations — this is what makes the logic testable without a real database.

### store/
Postgres implementations of the repository interfaces defined in domain/. Each file
handles the queries for one entity. No logic here, only SQL. If a service needs
data, it asks the store through the interface.

### handler/
The HTTP layer. Parses the request, calls the service, writes the response. Each file
groups the handlers for one domain area. No business logic lives here.


