---
doc_id: ECO-OBSERVABILITY-REQS-T07
title: TCIndustries — Economic Observability Requirements Brief (Task T-07)
version: "0.1"
date: 2026-10-23
status: PROPOSED — Analysis only. Creates no design authority.
authority_class: null   # proposal/requirements record; rules nothing, decides nothing
human_approved: false
controlling_documents: [MASTER-GDD-v1.1.1, HUMAN-RULINGS-REGISTER, EIC-PROGRAM-STATE]
reading_priority: advisory
archived_author_note: "Author: OpenCode documentarian agent under owner approval of Gap Analysis Task T-07 (2026-10-23). Assessor: none (single-session investigation)."
change_note: "v0.1 — Initial formal economic observability requirements brief for OQ-009 / SAFE-003 / RET-002, executed under the newly established HD-EIC-10/11 structural boundaries."
---

# TCIndustries — Economic Observability Requirements Brief (Task T-07)

**Document Type:** Requirements Specification / Proposal (Non-Canonical)  
**Status:** PROPOSED — Analysis only. Creates no design authority, changes no GDD status, selects no currency model, and sets zero numeric values.

---

## 1. Objective & Scope

Task T-07 converts the proven observability pattern demonstrated in testbed evidence (faucet/sink-tagged ledgers, append-only sale histories, snapshot jobs, volatility/concentration queries, and anomaly review lists per E-08/E-09) into formal **TCIndustries requirements** (SAFE-003). 

These requirements specify *what* must be observable, at *what* event boundaries, and under *what* privacy/fairness constraints, so that whatever currency, faucet/sink, or economic stabilization design the owner later selects is measurable and auditable from day one.

**Contraining Context:**
- **Currency neutrality:** This proposal selects **no currency model**. It does not assume single unified currency (HD-ECO-01) or multi-currency; it specifies the observability schema that supports economic governance regardless of currency structure.
- **HD-EIC-10/11 boundaries:** Observability must support verification of the adopted durable commitments (C1 non-internalisable scarcity, C2 persistent info/discovery asymmetry, C4 non-economic institutional/social constraint) and provide input evidence for future EIC validation without opening new gates.

---

## 2. Core Observability Requirements (SAFE-003)

### Req-Obs-01 — Append-Only Transaction & Faucet/Sink Ledger (E-09 Substrate)
- **Requirement:** Every economic transfer, creation (faucet), and destruction (sink) of value must be immutably recorded in an append-only ledger categorized by faucet/sink class and actor tier.
- **Event Boundary:** Recorded synchronously with state-changing transactions.
- **Purpose:** Enables real-time verification of faucet-exceeds-sink mandates (HD-ECO-01) and money supply velocity without invasive polling.

### Req-Obs-02 — Anomaly Review Lists (E-08 Pattern)
- **Requirement:** The system must automatically generate query-driven anomaly review lists covering:
  1. *Wash-trade pairs:* Repeated cyclic transfers between identical or linked nominal identities with near-zero net value change.
  2. *Near-zero-net churn:* High-volume inventory movements producing negligible economic surplus.
  3. *Placement bursts:* Abnormal concentration of structure or resource claims in time and space.
  4. *Price-spread anomalies:* Trades executed at extreme variance from regional moving medians.
- **Event Boundary:** Asynchronous batch processing (daily snapshot or rolling window).
- **Purpose:** Operationalises HD-SAFE-01 (Option A detection-first, human-review-only policy) by surfacing reviewable signals for human authority.

### Req-Obs-03 — Market Concentration & Volatility Metrics (RET-002)
- **Requirement:** Regional trading hubs and player vendors (HD-RET-01) must expose aggregate telemetry on:
  1. *Herfindahl-Hirschman Index (HHI):* Market concentration per resource kind and crafted item class.
  2. *Price Volatility Index:* Standard deviation of rolling transaction prices over time.
  3. *Inventory Aging:* Duration of stock retention on player vendors.
- **Event Boundary:** Periodic aggregation.
- **Purpose:** Ensures the project can detect monopolistic bottlenecks or market failures without imposing artificial price controls or account-level restrictions.

### Req-Obs-04 — Privacy & Scope Limits
- **Requirement:** Observability telemetry is strictly restricted to **economic and transactional surfaces** (SAFE-003 scope). Zero telemetry may capture social chat, private communications, or non-economic movement patterns.
- **Event Boundary:** Hard architectural boundary.
- **Purpose:** Prevents observability drift into invasive surveillance, protecting player privacy while preserving economic auditability.

---

## 3. Conclusion & Integration

These requirements provide the foundational audit machinery required by SAFE-003 and HD-SAFE-01. By defining observability independently of currency choice or numeric tuning, the project ensures that any future economic design is measurable, auditable, and aligned with project pillars.

**Status:** PROPOSED requirements specification. Ready for owner review and potential inclusion in future system specifications.
