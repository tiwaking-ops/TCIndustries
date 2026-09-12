# TCIndustries Initial GitHub Repository and Documentation Architecture Setup Instructions

**Document Type:** Project Setup / Documentation Governance Instruction  
**Project:** Tiwakings Craftworld Industries (TCIndustries)  
**Document Status:** NON-CANONICAL — Initial Repository Setup Guidance  
**Authority Status:** Advisory — Requires Human Confirmation Before Authority-Changing Adoption  
**Author:** GPT-5.6 Luna  
**Date:** 2026-09-12  
**Intended Executor:** OpenCode  
**Human Authority:** Project Designer / Owner  

---

## 1. Purpose

This document defines the minimum recommended procedure for establishing a GitHub repository for the **Tiwakings Craftworld Industries (TCIndustries)** project.

The repository is intended to become the controlled documentation environment for TCIndustries development.

The documentation workflow should be broadly analogous to the existing **Tiwas TTRPG** documentation architecture while remaining intentionally smaller at initial setup.

The repository architecture MUST be allowed to expand as the project documentation corpus, governance requirements, and development processes become better understood.

This document therefore establishes an **initial minimum architecture**, not a complete long-term documentation architecture.

---

## 2. Design Intent

The TCIndustries repository should support the following workflow:

```text
Human Designer
      |
      v
Project Decisions / Direction
      |
      v
Git Repository
      |
      +--> Canonical Project Documentation
      |
      +--> Governance Documentation
      |
      +--> Proposals
      |
      +--> Investigations / Audits
      |
      +--> Sources / References
      |
      +--> Archive
      |
      v
OpenCode
  as Documentarian
```

OpenCode should primarily function as a **project documentarian, repository steward, auditor, and documentation assistant**.

OpenCode is NOT automatically a game-design authority.

---

# 3. Minimum Repository Architecture

The initial repository SHOULD contain only the following structure:

```text
TCIndustries/
├── README.md
├── AGENTS.md
│
├── canonical/
│   ├── README.md
│   └── gdd/
│       └── TCIndustries_Master_GDD_v1.1.1_Status_Patch.md
│
├── governance/
│   ├── README.md
│   └── authority.md
│
├── proposals/
│   └── README.md
│
├── investigations/
│   └── README.md
│
├── sources/
│   └── README.md
│
└── archive/
    └── README.md
```

This structure is deliberately minimal.

OpenCode SHOULD NOT create numerous additional directories merely because they might eventually become useful.

Additional architecture SHOULD be proposed when examination of the repository demonstrates a genuine documentation or governance need.

---

# 4. Repository Creation

## 4.1 GitHub Repository

Create a new GitHub repository for TCIndustries.

Recommended initial repository name:

```text
TCIndustries
```

Alternative, if a more descriptive repository name is preferred:

```text
Tiwakings-Craftworld-Industries
```

The repository MAY be private during development.

The repository SHOULD initially contain:

- `README.md`
- Git version control
- the default `main` branch

No elaborate GitHub configuration is required at this stage.

---

# 5. Initial Canonical Material

The currently identified working TCIndustries documentation should be placed into the initial repository without silently changing its content or authority.

The current working GDD reference is:

```text
TCIndustries_Master_GDD_v1.1.1_Status_Patch.md
```

It SHOULD initially be located at:

```text
canonical/gdd/TCIndustries_Master_GDD_v1.1.1_Status_Patch.md
```

The existing authority/provenance material SHOULD initially be represented under:

```text
governance/authority.md
```

Where the original source document contains more information than is appropriate for the initial governance file, its content SHOULD be preserved rather than silently rewritten.

Any transformation, consolidation, reclassification, or authority change MUST be explicitly identified.

---

# 6. Treatment of Existing Project Documents

The following previously existing documents should NOT automatically be treated as independent canonical authorities merely because they exist:

```text
TCIndustries_Master_GDD_v1.0_Consolidated.md
TCIndustries_Master_GDD_v1.0_Canonical_Audit.md
TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md
TCIndustries_Authority_Provenance_Reconciliation_Matrix.md
TCIndustries_Master_GDD_v1.1.1_Status_Patch.md
```

Their authority MUST be determined according to their documented status and provenance rather than by:

- filename alone;
- chronology alone;
- version number alone;
- location alone;
- document length;
- apparent detail;
- Git commit recency;
- duplication;
- OpenCode's interpretation.

Historical, superseded, audit, provenance, and working documents SHOULD be preserved where useful.

They MAY eventually be placed under `archive/`, `investigations/`, `governance/`, or another appropriate location after their roles have been formally assessed.

OpenCode MUST NOT silently destroy or discard historical project information during repository setup.

---

# 7. Initial AGENTS.md Requirements

The root `AGENTS.md` SHOULD establish the repository's documentation-control rules.

At minimum, it SHOULD state that:

1. This repository is the controlled documentation record for TCIndustries.
2. `canonical/` contains project material currently designated as canonical.
3. `governance/` contains documentation-governance and authority rules.
4. Canonical status MUST NOT be inferred solely from presence in the repository.
5. OpenCode acts as a documentarian and repository steward.
6. OpenCode may:
   - inspect project documentation;
   - organize documentation;
   - identify contradictions;
   - identify missing documentation;
   - identify provenance problems;
   - perform audits;
   - propose documentation architecture;
   - propose governance improvements.
7. OpenCode MUST NOT silently make game-design decisions.
8. OpenCode MUST NOT silently promote a proposal, inference, historical statement, or prototype into canonical design.
9. Human approval is required for authority-changing decisions.
10. Provenance and document status MUST be preserved.
11. Evidence, inference, recommendation, and human ruling SHOULD remain distinguishable.
12. Repository architecture MAY evolve as project understanding improves.

---

# 8. Documentation Status Discipline

The TCIndustries repository SHOULD distinguish at least the following concepts:

| Status | Meaning |
|---|---|
| `LOCKED` | Explicitly established project authority |
| `PROPOSED` | Candidate design not yet authoritative |
| `TBD` | Decision intentionally unresolved |
| `DEFERRED` | Decision intentionally postponed |
| `PROTOTYPE` | Experimental implementation or evidence |
| `HISTORICAL` | Preserved for project history |
| `DERIVED` | Consequence inferred from established authority |
| `ASSUMPTION` | Temporary working assumption |

These statuses SHOULD NOT be treated as interchangeable.

In particular:

> **Presence in `canonical/` MUST NOT by itself be interpreted as proof of human approval.**

Where an existing TCIndustries document already contains its own authoritative status system, OpenCode SHOULD preserve that system unless the human designer explicitly authorizes a change.

---

# 9. OpenCode's Initial Role

After the repository has been created and the initial documents have been committed, OpenCode SHOULD receive a first-pass repository audit task.

The first task SHOULD NOT be to redesign the game.

It SHOULD be to understand the documentation corpus.

OpenCode SHOULD:

1. Inspect all existing project documentation.
2. Identify each document's apparent purpose.
3. Identify stated authority and status.
4. Identify provenance where available.
5. Identify relationships between documents.
6. Identify superseded or duplicated documents.
7. Identify contradictions.
8. Identify unresolved authority questions.
9. Identify missing documentation infrastructure.
10. Evaluate whether the initial repository structure is sufficient.
11. Recommend additional folders or documents only where justified.
12. Produce a report describing its findings.

OpenCode SHOULD NOT:

- silently rewrite game mechanics;
- silently resolve design contradictions;
- promote inferred principles to locked design;
- delete historical material;
- replace human rulings with LLM judgments;
- assume that the newest document is automatically authoritative;
- assume that the most detailed document is automatically authoritative.

---

# 10. Initial Git Commit

Once the initial repository structure has been established and checked, create the first commit.

Recommended commit message:

```text
Establish initial TCIndustries documentation architecture
```

The initial commit SHOULD establish the repository architecture and preserve the initial project documentation without unnecessary design changes.

---

# 11. Initial OpenCode Audit

After the initial commit, OpenCode SHOULD perform a repository-wide documentation audit.

Suggested task:

```text
Audit the TCIndustries documentation repository as a documentation
governance and architecture problem.

Do not redesign the game.

Determine:

1. What documentation currently exists.
2. What each document appears to represent.
3. Which documents claim canonical, proposed, historical, audit,
   prototype, or other status.
4. What authority and provenance information exists.
5. Where documents overlap or conflict.
6. Which documents appear superseded.
7. Which documentation relationships are unclear.
8. What governance infrastructure is currently missing.
9. Whether the initial repository architecture is adequate.
10. What additional repository structure should be considered.

Treat the human designer as the final authority.

Do not silently promote, demote, merge, delete, or rewrite project
authority.

Distinguish:
- documented fact;
- source evidence;
- inference;
- audit finding;
- recommendation;
- unresolved question.

Produce recommendations for improving the repository architecture
without making those recommendations canonical automatically.
```

This audit should become the basis for the next stage of repository development.

---

# 12. Expansion Strategy

The repository SHOULD evolve incrementally.

The intended process is:

```text
Initial Minimal Architecture
        |
        v
Repository Audit
        |
        v
Identify Genuine Documentation Needs
        |
        v
OpenCode Recommendations
        |
        v
Human Review
        |
        v
Repository Architecture Update
        |
        v
Repeat as Project Matures
```

This prevents premature creation of a large documentation bureaucracy before the project's actual needs are understood.

Possible future additions MAY include specialized documentation areas for:

- systems;
- economy;
- world design;
- professions;
- crafting;
- social systems;
- player organizations;
- technical architecture;
- implementation planning;
- playtesting;
- balancing;
- decisions;
- change logs;
- research;
- external references.

However, none of these SHOULD be created merely because they are theoretically useful.

They should be introduced when the project corpus demonstrates a need for them.

---

# 13. Authority Boundary

The repository architecture is a documentation-control mechanism.

It is NOT itself a mechanism for deciding game design.

The following distinction MUST be maintained:

```text
Repository Structure
        ≠
Game Design Authority
```

Likewise:

```text
OpenCode Recommendation
        ≠
Human Project Decision
```

And:

```text
Canonical File Location
        ≠
Automatic Human Approval
```

Any authority-changing action MUST be explicitly identifiable as such and SHOULD require human confirmation.

---

# 14. Recommended Initial Workflow

The minimum practical workflow is:

### Step 1 — Create GitHub Repository

Create the repository named `TCIndustries` or an equivalent descriptive name.

### Step 2 — Clone Repository

Clone the repository to the local development machine.

### Step 3 — Create Initial Directory Structure

Create the minimum directory structure defined in Section 3.

### Step 4 — Add Existing Documentation

Place the current working GDD and relevant governance/provenance material into the repository while preserving provenance.

### Step 5 — Create `README.md`

Describe the project and explain that the repository is the controlled project documentation environment.

### Step 6 — Create `AGENTS.md`

Establish OpenCode's role and the documentation-authority rules.

### Step 7 — Create Governance README

Explain the purpose of the `governance/` directory.

### Step 8 — Create Placeholder READMEs

Create minimal README files for:

```text
canonical/
proposals/
investigations/
sources/
archive/
```

### Step 9 — Review Before Commit

Confirm that no game-design authority has been changed merely as a result of repository setup.

### Step 10 — Commit

Use:

```text
Establish initial TCIndustries documentation architecture
```

### Step 11 — Push

Push the initial repository to GitHub.

### Step 12 — Give OpenCode the Initial Audit Task

Allow OpenCode to inspect the corpus and recommend the next documentation-architecture changes.

---

# 15. Acceptance Criteria

The initial repository setup is considered successful when:

- [ ] A GitHub repository exists.
- [ ] The repository can be cloned locally.
- [ ] The minimum directory architecture exists.
- [ ] The current working GDD is preserved.
- [ ] Existing authority/provenance information is preserved.
- [ ] `AGENTS.md` defines OpenCode's documentarian role.
- [ ] Canonical status is not inferred solely from file location.
- [ ] Human authority remains explicit.
- [ ] Historical and superseded material is not silently discarded.
- [ ] The initial commit contains no unapproved game-design changes.
- [ ] OpenCode can inspect the repository and perform a documentation audit.
- [ ] Future repository expansion can occur incrementally.

---

# 16. Final Designer-Controlled Principle

The initial repository should be **small, explicit, and expandable**.

The purpose of this setup is not to predict the final shape of TCIndustries documentation.

The purpose is to create enough structure that OpenCode can safely examine the existing project, document what actually exists, identify what is missing, and recommend progressively better architecture.

The intended long-term principle is:

> **Document the project that exists; do not invent the project that ought to exist.**

Any future expansion of the documentation system should follow the same authority, provenance, status, and human-approval discipline.

---

## 17. Status

**Status:** NON-CANONICAL ADVISORY  
**Decision Required:** Human approval of the initial repository architecture and OpenCode governance approach  
**Authority Effect:** None until explicitly adopted by the project designer  
**Next Recommended Action:** Establish the minimum GitHub repository, then perform an OpenCode repository/documentation audit before expanding the architecture.