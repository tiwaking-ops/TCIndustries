# Executive Verdict

**Author:** Perplexity (per filename, unverified)

The corpus has already moved beyond uncontrolled document creation: it contains a working canonical GDD, a governance/provenance matrix, audits, an EIC investigation package, simulation evidence, and human-ruling records. However, it is not yet reliably auditable by independent LLMs because the current state is distributed across several documents and the corpus lacks a single machine-readable state manifest that unambiguously binds authority, supersession, current decisions, unresolved questions, evidence, and change history.

The recommended solution is **moderate consolidation**: retain the current Master GDD as the high-level canonical design reference; add a compact project-state manifest, decision/change register, and system-package index; treat the EIC documents as a controlled package with explicit lifecycle roles; and preserve all superseded, historical, proposed, audit, and evidence documents in an immutable archive.

This is a documentation-architecture recommendation only. No design decision is made, and no files should be changed without human authorization.

# Current Documentation Assessment

## Corpus examined

The available project corpus contains 19 Markdown documents:

- Three Master GDD generations:
  - `TCIndustries_Master_GDD_v1.0_Consolidated.md`
  - `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md`
  - `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md`
- One canonical audit:
  - `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md`
- One authority/provenance reconciliation record:
  - `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md`
- Fourteen EIC investigation, proposal, simulation, evidence, decision, and ruling records.

The files are also available as a merged corpus, but the individual filenames remain essential for provenance and should not be replaced by the merged text as the authoritative project representation.

## Directly established findings

The corpus explicitly states that:

- `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` is the current working canonical design reference.
- The Master GDD must not be regenerated merely for stylistic improvement.
- Canonical presence does not itself prove human approval.
- Historical proposals do not automatically re-enter canon.
- Prototype behaviour is evidence rather than design authority.
- `PROPOSED`, `TBD`, `PROTOTYPE / EVIDENCE`, `DEFERRED`, `ASSUMPTION`, `HISTORICAL`, and `DERIVED CONSTRAINT` must not be silently promoted to `LOCKED`.

These rules are directly stated in the v1.1.1 Status Patch and the Authority & Provenance Reconciliation Matrix. [ppl-ai-file-upload.s3.amazonaws](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/20dd6917-89c0-47cf-94ae-7f9837d0a87e/TCIndustries-project-files-v1-merged.txt)

## Inference from the corpus

The documentation system is conceptually strong but operationally fragmented. A new LLM can reconstruct much of the current state, but it must compare at least:

1. the v1.1.1 Status Patch;
2. the Authority & Provenance Reconciliation Matrix;
3. the v1.0 Canonical Audit;
4. relevant human-ruling files;
5. relevant EIC evidence and proposal files;
6. prior GDD versions to understand supersession.

That is workable for a human familiar with the project, but not sufficiently deterministic for independent cross-LLM comparison.

# Current Documentation Inventory

| Document | Corpus role | Current assessment |
|---|---|---|
| `TCIndustries_Master_GDD_v1.0_Consolidated.md` | Consolidated design reference | Superseded by v1.1 and v1.1.1; preserve as provenance |
| `TCIndustries_Master_GDD_v1.0_Canonical_Audit.md` | Governance/design audit | Retain as audit record; not current design authority |
| `TCIndustries_Master_GDD_v1.1_Canonical_Baseline.md` | Corrected canonical baseline | Superseded by v1.1.1; preserve as provenance |
| `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | Current working canonical GDD | Active and authoritative for design-status interpretation, subject to its own stated limits |
| `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | Authority and provenance governance | Active governance reference; it creates no game-design decisions |
| `TCIndustries_Economic_Interdependence_Core_Design_Investigation.md` | Cross-system design investigation | Active working investigation, explicitly non-canonical |
| `TCIndustries_EIC_Candidate_v0.1.md` | Candidate architecture | Proposed design candidate; not authoritative |
| `TCIndustries_EIC_Candidate_v0.1_Falsification_Pass.md` | Adversarial/falsification analysis | Audit/evidence record for candidate evaluation |
| `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.md` | Simulation method | Evidence/prototype specification |
| `TCIndustries_EIC_Comparative_Simulation_Specification_v0.1.1.md` | Revised simulation method | Likely supersedes v0.1 if the file explicitly states that; this should be machine-recorded |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.md` | Simulation output | Evidence record |
| `TCIndustries_EIC_Comparative_Simulation_Results_v0.1.1.md` | Revised simulation output | Likely current EIC evidence result; not design authority |
| `TCIndustries_EIC_Minimum_Simulation_Results_v0.1.md` | Reduced-scope simulation evidence | Evidence record; scope must be distinguished from comparative results |
| `TCIndustries_EIC_Provisional_Functional_Shapes_v0.1.md` | Provisional system-shape synthesis | Proposed/non-canonical; likely a bridge between evidence and decision |
| `TCIndustries_EIC_Functional_Shapes_Refinement_Targeted_γ_Closure_v0.2.md` | Refined functional-shape analysis | Proposed or analytical refinement; not authoritative unless human rulings say otherwise |
| `TCIndustries_EIC_Targeted_γ_Prime_Adversarial_Retest_Results_v0.1.md` | Adversarial retest evidence | Evidence/audit record; not design authority |
| `TCIndustries_EIC_Human_Decision_Brief_v0.1.md` | Decision-support summary | Human-consumption proposal document; not itself a decision |
| `TCIndustries_EIC_Human_Rulings_2026-08-25.md` | Human ruling record | Potentially authoritative for the decisions explicitly recorded inside it |
| `TCIndustries_EIC_Human_Rulings_Architecture_C_Falsification_2026-08-25.md` | Human ruling record concerning Architecture C/falsification | Potentially authoritative only for explicit rulings recorded in that file |

The final two human-ruling files require particularly careful treatment. A filename containing “Human Rulings” is evidence of a governance role, but authority should be determined from the document’s internal ruling fields, not from its filename alone. This follows the project’s own rule that document naming or canonical presence must not substitute for explicit authority. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033620Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=f7517fa58615d6e7c5fd9936e45d7989db8d065d22e754d28dc3a3b5d4610ff0)

## Authority hierarchy currently visible

The corpus supports the following hierarchy:

1. Explicit human/project ruling.
2. Explicit project principle or pillar.
3. `HUMAN-LOCKED` / `LOCKED` material whose authority is traceable to one of the above.
4. `DERIVED CONSTRAINT`, provided the derivation is genuinely logical.
5. `PROPOSED`.
6. `TBD`.
7. `PROTOTYPE / EVIDENCE`.
8. `ASSUMPTION`.
9. `HISTORICAL` or external reference.

The Authority Matrix formalises this more precisely as authority classes A–G. In particular, classes C–E are not substitutes for human approval, while prototype evidence and historical material do not create design authority. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033620Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=f7517fa58615d6e7c5fd9936e45d7989db8d065d22e754d28dc3a3b5d4610ff0)

# Redundancy and Sprawl Analysis

## Master GDD duplication

The three GDD generations repeat most of the same design content:

- vision;
- pillars;
- loops;
- professions;
- resources;
- crafting;
- manufacturing;
- economy;
- services;
- combat;
- buildings;
- organisations;
- progression;
- emergence;
- unresolved questions;
- provenance and change notes.

This repetition is not inherently wrong because each version records historical state. The problem is that the active-state distinction is not externalised into a concise manifest. An LLM must infer the version chain from document headers and lineage sections.

The v1.0 audit explicitly identifies overlap and restatement inflation, including:

- `ECO-001` overlapping with `VIS-003`;
- `PROG-001` restating `LOOP-003`;
- several anti-goal and derived-constraint statements repeating higher-level principles.

The audit considered this acceptable for readability but identified it as material that could eventually be tightened. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Economic_Interdependence_Core_Design_Investigation.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033646Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=37bb20062c0bea3aa4ac959305b922eb8a22875574b1a6cfd22005ed45d2c820)

## EIC package duplication

The EIC documents repeat a common conceptual structure:

- specialisation;
- dynamic resources;
- crafting differentiation;
- manufacturing constraints;
- NPC substitution;
- demand;
- multi-accounting;
- comparative architecture assessment;
- Principle → Constraint → Invariant → Failure Condition → Test.

The repetition has legitimate functions when the documents are sequential records:

- the investigation defines the problem;
- candidate documents describe alternatives;
- simulations provide evidence;
- falsification and retest documents challenge the candidates;
- provisional-shape documents synthesise findings;
- human-ruling documents record decisions.

The danger is not duplication itself. The danger is that an LLM may mistake repeated proposal language for increased authority. The investigation explicitly warns that its recommendations are `PROPOSED` and that it creates no `LOCKED` decisions.[4]

## Redundancy that should not be removed

The following duplication is governance-useful and should be preserved:

- a decision appearing in a human-ruling record and being reflected in the current canonical document;
- a proposal being restated in a simulation specification;
- an evidence result referring back to the tested proposal;
- an audit record documenting why a status changed;
- a superseded document retaining its original wording.

The architecture should therefore reduce **uncontrolled discovery duplication**, not eliminate historical records or all repeated context.

# Authority and Provenance Risks

## 1. Canonical ambiguity

The v1.1.1 Status Patch states that it is current, but the corpus also contains the v1.1 Baseline and v1.0 Consolidated GDD. A new LLM must read each document’s header to establish the chain.

This is vulnerable to:

- reading files in alphabetical order;
- treating the first “canonical” document encountered as current;
- treating all documents marked canonical as simultaneously authoritative;
- failing to recognise that a later status patch changes classification without adding major design content.

## 2. Human approval versus canonical status

The Authority Matrix identifies several items that were marked `LOCKED` but whose independent human authority was not established:

- `WRLD-001`;
- `PLR-001`;
- `PLR-003`;
- `CRFT-001`;
- `SAFE-001`.

The Matrix records these as requiring authority review rather than silently demoting them. This is an important example of preserved uncertainty. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033620Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=f7517fa58615d6e7c5fd9936e45d7989db8d065d22e754d28dc3a3b5d4610ff0)

The v1.1.1 Status Patch subsequently records human confirmation of these principles. An independent LLM needs a clear binding relationship between the ruling record and the exact canonical status entries. Without that binding, it may report the earlier authority concern as still open or incorrectly treat the GDD status as sufficient evidence.

## 3. Status transition ambiguity

The corpus contains important transitions:

- `CRFT-006`: from `LOCKED` or derived-looking material to `PROPOSED`;
- `PROG-002`: from `LOCKED`/derived material to `DERIVED CONSTRAINT` in v1.1, with the Matrix later recommending `PROPOSED`;
- `MFG-004`: from `LOCKED` to `DERIVED CONSTRAINT`;
- `ECO-003`: from `LOCKED` to `DERIVED CONSTRAINT`;
- `CMBT-003`: from `LOCKED` to `DERIVED CONSTRAINT`.

The Matrix is explicit about these dispositions, but the status history needs to be represented as structured transitions rather than reconstructed through prose comparison. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033620Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=f7517fa58615d6e7c5fd9936e45d7989db8d065d22e754d28dc3a3b5d4610ff0)

## 4. Prototype authority contamination

The project explicitly states that prototype behaviour is evidence only and that the GDD governs unless a later human ruling changes that relationship. This is a strong rule, but it needs to be visible in the active audit entry point, not only inside the GDD and governance documents.

The v1.0 Audit identifies prototype inspection as a high-severity process gap: no prototype had been inspected, so claims about implementation reality remained unsupported. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Economic_Interdependence_Core_Design_Investigation.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033646Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=37bb20062c0bea3aa4ac959305b922eb8a22875574b1a6cfd22005ed45d2c820)

## 5. Historical proposal re-entry

The Matrix provides an explicit non-canonical boundary for earlier proposals, including:

- specific skill-box architecture;
- two-to-three concurrent profession targets;
- particular resource lifecycle models;
- specific experimentation-point models;
- permanent factory exclusion from Exceptional output;
- a single unified currency;
- Commerce Directory;
- specific attribute pools;
- use-based skill acquisition.

This boundary is highly valuable, but it is currently embedded in a governance document rather than exposed through a project-wide proposal register. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Authority_Provenance_Reconciliation_Matrix.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033620Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=f7517fa58615d6e7c5fd9936e45d7989db8d065d22e754d28dc3a3b5d4610ff0)

## 6. “Human Rulings” records require explicit ruling granularity

A human-ruling document should not merely contain prose such as “Architecture C is preferred.” It must distinguish:

- decision identifier;
- exact ruling;
- scope;
- status;
- approval authority;
- date;
- affected rule IDs;
- what remains unresolved;
- whether the ruling changes canon;
- source evidence considered.

Without that granularity, an LLM may promote an entire candidate architecture when the human only approved a constraint, a test direction, or a rejection of one failure mode.

# Cross-LLM Auditability Assessment

| Criterion | Assessment | Evidence / reasoning |
|---|---|---|
| A. State reconstruction | Partial | The v1.1.1 GDD is identified as current, but state is distributed across GDD, Matrix, rulings, and EIC files |
| B. Authority reconstruction | Strong principles, weak indexing | Governance vocabulary is explicit, but binding source-to-rule links are not centralised |
| C. Temporal reconstruction | Moderate | Version lineage exists, but a machine-readable supersession graph is absent |
| D. Provenance | Moderate to strong | The Matrix and GDD contain provenance, but important statements require cross-document reconciliation |
| E. Cross-LLM comparability | Partial | Independent LLMs may choose different reading orders and different interpretations of repeated proposal language |
| F. Change detection | Weak to moderate | Version notes exist, but no compact authoritative change ledger covers the whole project |
| G. Scope control | Strong in principle, moderate in operation | The status vocabulary and non-canonical boundary are strong; discoverability remains the weakness |
| H. Scalability | Moderate | The GDD can remain high-level, but future systems need package indexing and lifecycle control |

## Overall result

The project passes the governance-discipline test but only partially passes the reproducibility test.

Two independent LLMs are likely to agree on:

- the existence of the current working GDD;
- the status vocabulary;
- the non-canonical nature of older proposals;
- the distinction between evidence and authority;
- the unresolved nature of detailed mechanics.

They may disagree on:

- whether particular EIC documents remain active;
- which simulation result is current;
- whether a human ruling approved an architecture, a constraint, or merely a direction;
- whether an item remains in the authority-review queue;
- which document controls a particular EIC rule;
- whether a later provisional-shape document supersedes an earlier candidate or merely analyses it.

# Recommended Documentation Architecture

The recommended architecture is a **small active control layer plus domain packages plus immutable archive**.

The control layer should answer “What is current?” without requiring an LLM to inspect every file. Domain packages should contain substantive material. The archive should preserve historical documents without allowing them to participate in current-state reconstruction unless explicitly requested.

## Proposed active document set

| Document | Purpose | Authority | Active/Archive | Primary Consumer |
|---|---|---|---|---|
| `TCIndustries_Project_State_Manifest.md` | Current project state, active documents, authority rules, status vocabulary, unresolved gates, latest changes | Governance index; does not create game-design decisions unless explicitly stated | Active | Every independent LLM |
| `TCIndustries_Master_GDD_v1.1.1_Status_Patch.md` | High-level canonical game design and current rule statuses | Current working canonical design reference | Active | Designers, auditors, LLMs |
| `TCIndustries_Authority_Provenance_Reconciliation_Matrix.md` | Authority classes, human-approval boundaries, non-canonical boundary, status reconciliation | Governance authority; no new game design | Active | Auditors and LLMs |
| `TCIndustries_Decision_and_Change_Register.md` | Decision IDs, status transitions, approvals, supersession, open/closed decisions | Authoritative governance record for change history | Active | Auditors, maintainers, LLMs |
| `TCIndustries_Open_Questions_Register.md` | Current unresolved questions, dependencies, ownership, blocking status, next required evidence | Authoritative for unresolved-state tracking; does not resolve questions | Active | Planning and audit |
| `TCIndustries_System_Package_Index.md` | Index of system packages, current package state, controlling documents, evidence, proposals, and audit records | Governance index | Active | LLMs navigating system work |
| `TCIndustries_EIC_Current_Package.md` | Compact current EIC state: approved constraints, proposals under consideration, evidence summary, unresolved questions, links to records | Only explicitly approved portions are authoritative; all other material remains labelled | Active | EIC auditors and designers |
| Existing detailed EIC records | Investigation, candidate, simulation, falsification, retest, and ruling history | Each retains its own declared status | Archive/package records | Targeted audit |
| Prior GDD versions | Historical canonical/provisional states | Historical provenance only | Archive | Historical reconstruction |
| Historical proposals and source material | Origin and rejected/non-canonical design material | Historical/non-canonical | Archive | Provenance and comparative analysis |

## `Project_State_Manifest` contents

This should be short enough for every audit and include:

- project identifier;
- manifest version;
- state date;
- current canonical GDD;
- current governance documents;
- current decision register;
- current open-question register;
- active system packages;
- archive boundary;
- authority hierarchy;
- status vocabulary;
- prototype/evidence rule;
- historical proposal rule;
- latest change IDs;
- known audit limitations;
- required reading order.

It should explicitly state:

> Canonical presence is not proof of human approval.

It should also state:

> No document outside the active set is part of current state unless referenced by an active document or specifically requested for historical/provenance review.

## `Decision_and_Change_Register` contents

Each entry should include:

```yaml
change_id: CHG-2026-08-25-001
subject_id: CRFT-006
change_type: status_change
previous_status: DERIVED CONSTRAINT
new_status: PROPOSED
authority_class: A
decision_source: TCIndustries_EIC_Human_Rulings_2026-08-25.md
effective_date: 2026-08-25
controlling_document: TCIndustries_Master_GDD_v1.1.1_Status_Patch.md
supersedes:
  - source-reference-or-prior-entry
human_approval: explicit
scope: game-design-status
notes: "Mechanism for achieving specialisation remains unresolved."
```

The exact field names may be adjusted, but the concepts should remain stable.

## `System_Package_Index` contents

Each system package should have:

- package ID;
- package name;
- scope;
- controlling canonical rules;
- current package summary;
- current status;
- active investigation;
- active proposal;
- latest evidence;
- latest audit/falsification;
- latest human ruling;
- unresolved questions;
- excluded historical proposals;
- package owner or approval authority, if established;
- last reviewed date.

This avoids requiring an LLM to discover all EIC-related files through filenames.

# Active Documentation Set

## 1. Project State Manifest

**Purpose:** Provide the deterministic entry point for any new LLM.

**Authority:** Governance index. It should not silently create design authority. It may authoritatively identify which existing documents control which questions.

**Audience:** Every LLM, auditor, contributor, and human reviewer.

**Belongs inside:**

- current state summary;
- active-document list;
- authority hierarchy;
- status vocabulary;
- supersession rules;
- required reading order;
- latest change IDs;
- explicit uncertainty and limitations.

**Must not contain:**

- new game-design proposals;
- unapproved architecture recommendations;
- duplicated system specifications;
- long historical explanations.

**Relationship:** Points to the GDD, governance matrix, registers, and system-package index.

**Authoritative?** Yes for navigation and document state; no for unapproved design content.

**Update frequency:** Every accepted structural or status change.

**Normally read during audit?** Always.

## 2. Master GDD

**Purpose:** Maintain the high-level canonical design reference.

**Authority:** Current working canonical design reference, subject to explicit status labels.

**Belongs inside:**

- vision;
- pillars;
- approved rules;
- derived constraints;
- proposals and TBDs where needed for design context;
- high-level unresolved questions;
- links to detailed packages.

**Must not contain:**

- full investigation histories;
- repeated simulation results;
- every rejected proposal;
- detailed technical implementation;
- unapproved EIC architecture presented as current design.

**Relationship:** Controlled by the authority matrix and decision register; detailed packages must not override it without explicit change.

**Authoritative?** Yes for the design statements it contains, according to their status.

**Update frequency:** Only after approved design/status changes; not after every analysis.

**Normally read during audit?** Always, after the manifest.

The project’s own v1.1.1 document says that no further LLM consolidation of the Master GDD should occur. The recommended architecture preserves that principle. [ppl-ai-file-upload.s3.amazonaws](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/20dd6917-89c0-47cf-94ae-7f9837d0a87e/TCIndustries-project-files-v1-merged.txt)

## 3. Authority/Provenance Matrix

**Purpose:** Explain how authority is assigned.

**Authority:** Governance authority, not game-design authority.

**Belongs inside:**

- A–G authority classes;
- human-approval rules;
- historical boundary;
- evidence boundary;
- derived-constraint test;
- known authority-review queues.

**Must not contain:**

- new mechanics;
- informal architecture decisions;
- untracked status changes.

**Normally read during audit?** Always.

## 4. Decision and Change Register

**Purpose:** Make temporal reconstruction deterministic.

**Authority:** Authoritative record of accepted decisions and status transitions.

**Belongs inside:**

- decision IDs;
- status changes;
- approval source;
- affected rule IDs;
- supersession;
- effective date;
- change rationale;
- unresolved remainder.

**Must not contain:**

- unapproved recommendations;
- full debates;
- simulation detail.

**Normally read during audit?** Yes, after the GDD and authority matrix.

## 5. Open Questions Register

**Purpose:** Establish what remains unresolved.

**Authority:** Authoritative inventory of unresolved questions, not answers.

**Belongs inside:**

- question ID;
- exact unresolved question;
- current status;
- dependency;
- blocking classification;
- evidence available;
- decision required;
- last review;
- related proposals and tests.

**Must not contain:**

- guessed answers;
- proposal language presented as resolution.

The v1.0 Audit already identifies EIC dependencies such as specialisation budget, resource attributes, experimentation, manufacturing, provenance, currency, city/governance, and prototype evidence. These should be indexed centrally rather than rediscovered inside multiple documents. [pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws](https://pplxfilegitgateway-projectfiles-prod-use1.s3.us-east-1.amazonaws.com/projectfiles-prod/gateway-object-store/repos/files-86a01c8bc03d47db8256f30f54010452/downloads/6b4c61eb64e7fb5abfaa31067bf7c9a776bc6806/TCIndustries_Economic_Interdependence_Core_Design_Investigation.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Checksum-Mode=ENABLED&X-Amz-Credential=ASIA2F3EMEYE4BREUBFJ%2F20260826%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20260826T033646Z&X-Amz-Expires=600&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEEcaCXVzLWVhc3QtMSJHMEUCIQCrRv%2BZhAbdJyzHC2idER32drIT7DB0h3eKx7Tcq2BirwIgaIcLajPlXnItaJDnMfZqzes91ORoPerjWbI%2FZdhbqMUqpwUIEBABGgw2OTk3NTMzMDk3MDUiDHAV971C%2F2xWWptJqCqEBcBKlYxt%2BZG6has7OY2TEeeAWa6WvJTAfQ6hysINc8H%2FAT8xA7sWSOicsI12DCVeECO2oxuFosix%2FirBSlR8nbe9IrII8ct%2FU9EKfpnns2FaH%2BqrTmxZ%2BPbQGkClTe32KX%2BmHk%2B5HmOG%2Bc6vw5nvePFthToNMr%2B14FqnQIRSWKu86uNwJv5eG2YV4rCDJiOQt1dr2wHzvBK7yZJJyZooolpMJ7ePkZDyvqzdy8CgDsriSva%2Fpw%2F%2FhrV91MVEtNygQwerlUzmNO8cTCYERFmXzqcE9vgSWbX8MViQ0l2ABoyzlrM9WhZdpZlFCETVl%2FuvsXBkuwuayzq6iYg%2BJY3Q4PffCzsVubq5ygewK62u1Gpkmk3Z2VqhfLw%2BtDd9BYtsXDBG6fYKtX89i8AupShINyh3NvjysZWlaVbbYFpR0DcTEarBBu4vEvzqqSdp3707vmbZFJd0YorMSsqeo1Nmi0ECP2bU%2FjbelJ9E5DOBj%2F5bKrDA9I2kS85awgsr8mJWwBRb7vdtlxWcAxKIRnZyGiS3Oj%2Bilt9ZhVXI5sKJ0bAYNXMNmpI17q%2FrurLaJS04IvQ5DZ1ErjBciKhA6kkZfemZPjADoATGIzCv9TWij5fL2ll4M%2FAbt%2BIsks5S8NzBA5W6zGVZt4GGNivHt2RzKjaVFF0EoLOp2dJ3EjcfrSmcbS6VhX4oYJv%2BWh%2Fw%2BHrwEFYmWh25puMHslZ%2BO7IbDrImwEHJoAFQj%2FATDPNWS1HTZIJqlrh1Er9go0Wvfxm6Pfd%2FwiKYxlG%2FOFHW0eL5TZ8VRianlpy3GxdO5Jz9nQEfvFpxW1k44DaYUpr2VJ3ZR8XwkmdwvIdDpjK0WuT21VGyOhLjMN%2B4uNQGOpABoz3JLafTsm23SN2KKNOOBNmDRtvw26HEsqeA%2B%2Bxhre%2FhGsNSnhBQNQi09NlTp892TEAKZ6AGk4Nzh%2FoEJ5QdwV7uJvjWqlsVo0Y5%2Bs7SKx6iqdxUheofdDJb3nk0mv4%2B%2FI%2F%2FtV1wTBNyrAg%2FDft%2BKdsAKUqAU6TVIUy8Rk5BZudwcrbTi8p1Q7e40Rdc1BZM&X-Amz-SignedHeaders=host&x-id=GetObject&X-Amz-Signature=37bb20062c0bea3aa4ac959305b922eb8a22875574b1a6cfd22005ed45d2c820)

## 6. System Package Index

**Purpose:** Prevent system-specific document sprawl.

**Authority:** Navigation and lifecycle index.

**Belongs inside:**

- package-level state;
- controlling documents;
- evidence;
- proposals;
- decisions;
- audits;
- archive links.

**Must not contain:**

- full specifications;
- duplicate design prose.

**Normally read during audit?** Read when auditing a specific system; otherwise skim.

## 7. EIC Current Package

**Purpose:** Provide a compact state view without deleting EIC history.

**Authority:** Only explicitly approved elements are authoritative. The package must label proposals, evidence, inference, and human decisions separately.

**Belongs inside:**

- canonical inputs;
- explicit human rulings;
- current candidate status;
- evidence summary;
- rejected or falsified candidates;
- open questions;
- links to full records;
- “not established” statements.

**Must not contain:**

- unresolved proposals written as decisions;
- unqualified conclusions from simulations;
- historical SWG material treated as TCIndustries authority;
- full duplicated simulation methodology.

**Normally read during audit?** Yes for an EIC audit; not necessarily for an unrelated audit.

The EIC investigation itself establishes that its recommendations are non-canonical and that it creates no new locked design decisions. That distinction should remain visible in the current package.[4]

# Archive / Provenance Structure

The archive should be immutable by convention and separated by record type:

```text
archive/
  gdd/
    superseded/
  governance/
    completed-audits/
  decisions/
    historical-rulings/
  systems/
    eic/
      investigations/
      candidates/
      specifications/
      results/
      falsification/
      retests/
      decision-briefs/
  historical-proposals/
  external-reference/
```

The exact directory names are implementation details and require approval. The important governance properties are:

- archived files remain readable;
- archived files are never silently treated as current;
- every archived record has a stable document ID;
- every archived record has status and supersession metadata;
- the active manifest identifies the archive boundary;
- historical proposals can be queried without entering the current state.

The existing Master GDD versions should not be deleted. v1.0 is explicitly superseded by v1.1, and v1.1 is superseded by v1.1.1; each remains valuable for temporal reconstruction. [ppl-ai-file-upload.s3.amazonaws](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/20dd6917-89c0-47cf-94ae-7f9837d0a87e/TCIndustries-project-files-v1-merged.txt)

# LLM Audit Reading Protocol

## Deterministic reading order

### Step 1: Read the state manifest

Determine:

- current state date;
- current canonical documents;
- active package documents;
- authority hierarchy;
- latest change ID;
- archive boundary;
- known limitations.

### Step 2: Read the current Master GDD

Extract only statements marked:

- `LOCKED` / `HUMAN-LOCKED`;
- `DERIVED CONSTRAINT`;
- `PROPOSED`;
- `TBD`;
- `PROTOTYPE / EVIDENCE`;
- `DEFERRED`;
- `ASSUMPTION`;
- `HISTORICAL`.

Do not treat explanatory prose as overriding a status tag.

### Step 3: Read the Authority/Provenance Matrix

For every important statement, determine:

- authority class;
- whether human approval is explicit;
- whether the statement is derived;
- whether it is only inferred;
- whether it is historical or evidentiary.

### Step 4: Read the Decision and Change Register

Check:

- whether the statement changed;
- previous and current statuses;
- effective date;
- controlling document;
- superseded documents;
- unresolved remainder.

### Step 5: Read the Open Questions Register

Record all relevant unresolved questions. Do not convert a proposed answer into a resolved decision.

### Step 6: Read the relevant system-package index entry

For a system such as EIC, identify:

- current package;
- active investigation;
- latest evidence;
- latest audit;
- latest ruling;
- historical records.

### Step 7: Read the current package summary

Separate:

- approved rules;
- derived constraints;
- proposals;
- evidence;
- inferences;
- rejected alternatives;
- unresolved questions.

### Step 8: Consult historical records selectively

Read archives only when:

- provenance is requested;
- a status transition requires verification;
- a conflict remains unresolved;
- an audit asks whether a proposal was previously considered;
- the current package explicitly links to a historical record.

## Standard audit output

Every audit should report:

1. **Scope and files read**
2. **Current state**
3. **Human-approved material**
4. **Derived constraints**
5. **Proposals**
6. **Unresolved questions**
7. **Evidence and prototype status**
8. **Superseded or historical material**
9. **Conflicts**
10. **Missing information**
11. **Independent verification required**
12. **Confidence and uncertainty**

## Evidence citation format

Use stable document IDs and rule IDs:

```text
[DOC:TCI-GDD-1.1.1 §12 CRFT-006]
[DOC:TCI-AUTH-1.0 §4]
[DOC:TCI-EIC-INV-1.0 §6]
[DEC:2026-08-25-003]
[EV:EIC-SIM-0.1.1]
```

Line numbers may be added where available, but section and rule IDs should remain primary because line numbers can change.

## Uncertainty reporting

Use explicit categories:

- `ESTABLISHED`: directly supported by an authoritative current record.
- `DERIVED`: logically follows from established rules.
- `PROPOSED`: recommended but not approved.
- `EVIDENCE-ONLY`: observed in a prototype or simulation.
- `HISTORICAL`: prior or external material.
- `CONFLICTED`: sources disagree and no authority rule resolves the conflict.
- `NOT ESTABLISHED`: the corpus does not establish the claim.

# Recommended Standard Metadata

A lightweight YAML front matter should be added to future documents:

```yaml
document_id: TCI-EIC-INV-001
title: Economic Interdependence Core Design Investigation
document_type: investigation
version: 1.0
status: non-canonical
authority: none
authority_class: E
project: TCIndustries
created: 2026-08-25
updated: 2026-08-25
controlling_document: TCI-GDD-1.1.1
supersedes: []
superseded_by: []
related_decisions: []
related_rules:
  - VIS-004
  - PIL-003
  - CRFT-001
  - MFG-001
related_questions:
  - OQ-001
  - OQ-003
  - OQ-004
  - OQ-005
evidence_role: analysis
human_approval: none
creates_design_decisions: false
archive_policy: retain
```

## Minimum metadata fields

- `document_id`
- `title`
- `document_type`
- `version`
- `status`
- `authority`
- `authority_class`
- `created`
- `updated`
- `controlling_document`
- `supersedes`
- `superseded_by`
- `related_decisions`
- `related_rules`
- `related_questions`
- `evidence_role`
- `human_approval`
- `creates_design_decisions`
- `archive_policy`

## Controlled values

`document_type` should use a small controlled vocabulary:

- `canonical-design`
- `governance`
- `audit`
- `decision`
- `investigation`
- `proposal`
- `evidence-specification`
- `evidence-result`
- `falsification`
- `prototype-record`
- `open-questions`
- `historical`
- `external-reference`
- `index`

`status` should preserve project vocabulary but may include document-level lifecycle values such as:

- `active`
- `superseded`
- `archived`
- `draft`
- `non-canonical`

These document lifecycle labels must not replace design statuses such as `LOCKED`, `DERIVED CONSTRAINT`, `PROPOSED`, or `TBD`.

# Migration / Consolidation Plan

This is a proposal-only staged plan. No files should be moved, renamed, merged, or modified without authorization.

## Stage 0 — Freeze interpretation

Create no new design content.

Agree that:

- v1.1.1 is the current working canonical GDD;
- the Authority Matrix controls authority interpretation;
- older GDDs are historical;
- EIC proposals and evidence do not become canon automatically;
- the archive is preserved.

## Stage 1 — Create identity and lifecycle metadata

Assign stable document IDs to the existing corpus.

Record:

- document type;
- version;
- status;
- authority;
- supersession;
- controlling document;
- evidence role;
- related rule IDs;
- related decision IDs;
- related question IDs.

Do not rewrite substantive content at this stage.

## Stage 2 — Create the Project State Manifest

The manifest should point to the existing documents without replacing them.

It should identify:

- current GDD;
- current governance matrix;
- active EIC package documents;
- current human-ruling records;
- current open questions;
- latest changes;
- archive boundary.

## Stage 3 — Create the Decision and Change Register

Backfill only material changes already documented, including:

- GDD v1.0 → v1.1;
- v1.1 → v1.1.1;
- status introduction of `DERIVED CONSTRAINT`;
- reclassification of `CRFT-006`;
- reclassification of `PROG-002`;
- reclassification of `MFG-004`, `ECO-003`, and `CMBT-003`;
- human confirmation of the authority-review queue;
- EIC rulings, but only to the extent explicitly recorded.

## Stage 4 — Create the Open Questions Register

Extract existing questions rather than inventing new ones.

At minimum, index the EIC questions identified by the audit and investigation:

- specialisation budget;
- resource lifecycle;
- resource attributes;
- experimentation;
- manufacturing;
- currency;
- durability and repair;
- respecialisation;
- multi-accounting;
- prototype reconciliation;
- vendor/retail;
- provenance.

The v1.1.1 GDD and EIC investigation already preserve many of these as unresolved. [ppl-ai-file-upload.s3.amazonaws](https://ppl-ai-file-upload.s3.amazonaws.com/web/direct-files/attachments/7508106/20dd6917-89c0-47cf-94ae-7f9837d0a87e/TCIndustries-project-files-v1-merged.txt)

## Stage 5 — Create the System Package Index

Start with EIC because it has the largest active document family.

For each EIC record, identify:

- lifecycle role;
- current/superseded status;
- evidence relationship;
- proposal relationship;
- decision relationship;
- unresolved issues.

## Stage 6 — Create a compact EIC current-state document

This should be a navigational and state-reconstruction summary, not a replacement for the evidence archive.

It should clearly distinguish:

- human rulings;
- canonical constraints;
- proposed architectures;
- evidence outcomes;
- falsified or weakened hypotheses;
- unresolved questions.

## Stage 7 — Archive classification

Mark older files as:

- superseded GDD;
- completed audit;
- historical proposal;
- evidence record;
- decision record;
- active investigation;
- active proposal;
- current ruling.

Do not delete or silently merge them.

## Stage 8 — Establish the audit protocol

Adopt the deterministic reading order and standard audit-output schema.

This is the primary measure of success: whether different LLMs produce comparable state reports.

# Risks and Trade-offs

## Risk: Too many active control documents

Adding a manifest, registers, and package index increases the active file count.

**Mitigation:** Keep each compact and make the manifest the single entry point. The goal is not minimum file count; it is minimum ambiguity.

## Risk: Metadata bureaucracy

Excessive front matter could slow work and encourage people to bypass the system.

**Mitigation:** Require metadata only for identity, lifecycle, authority, provenance, and relationships. Do not require elaborate templates for every paragraph.

## Risk: Manifest drift

The manifest could become stale and contradict the underlying documents.

**Mitigation:** Update it as a mandatory part of accepted structural or authority changes; include a manifest date and latest change ID.

## Risk: False precision in authority metadata

A field such as `human_approval: true` could conceal uncertainty if the approval scope is unclear.

**Mitigation:** Record approval scope and affected rule IDs, not merely a Boolean.

## Risk: EIC summary becomes a second canonical GDD

A compact EIC package could accidentally become an alternative source of truth.

**Mitigation:** State explicitly that the EIC package cannot override the Master GDD and that only explicit human-ruling entries are authoritative.

## Risk: Archive contamination

LLMs may still inspect archived files and mistake them for current material.

**Mitigation:** Require every archive file to contain prominent `HISTORICAL`, `SUPERSEDED`, or `NON-CANONICAL` metadata and require the manifest to define the active boundary.

## Risk: Conflicts are hidden by consolidation

Merging multiple documents into a “clean” current summary could erase disagreement.

**Mitigation:** Preserve a conflict register and record unresolved conflicts as `CONFLICTED`, rather than selecting a winner without a documented authority rule.

# Human Authorization Gates

The following require explicit project-owner approval before implementation:

1. Creating the Project State Manifest as an official governance document.
2. Selecting and declaring the official active-document set.
3. Declaring v1.1.1 the sole current working canonical GDD entry point.
4. Creating the Decision and Change Register.
5. Creating the Open Questions Register.
6. Creating the System Package Index.
7. Creating an EIC current-state package or summary.
8. Assigning stable document IDs and controlled metadata fields.
9. Recording supersession relationships between existing documents.
10. Marking documents as active, superseded, archived, audit, evidence, proposal, or historical.
11. Establishing the deterministic audit reading order as project procedure.
12. Adopting the standardized audit output format and citation convention.
13. Confirming whether the human-ruling documents authorise only specific constraints or also any broader EIC architecture.
14. Confirming the final disposition of all authority-review items and unresolved EIC questions.
15. Approving any future edits to the Master GDD, including status changes or references to new system packages.

# Target State

A new LLM should be able to read:

1. the Project State Manifest;
2. the current Master GDD;
3. the Authority/Provenance Matrix;
4. the Decision and Change Register;
5. the Open Questions Register;
6. the relevant System Package Index entry;

and reconstruct the current state without reading the historical archive.

The target state should make the following impossible by normal reading order:

- treating an old GDD as current;
- treating repeated proposals as approved;
- treating prototype behaviour as design authority;
- treating a simulation result as a design ruling;
- treating a human decision brief as a human decision;
- treating a historical proposal as reintroduced canon;
- treating `DERIVED CONSTRAINT` as `HUMAN-LOCKED`;
- treating canonical presence as proof of human approval.

# Open Questions / Information Not Established by the Corpus

The corpus does not establish, with sufficient certainty from the material inspected here:

- whether the EIC human-ruling files approved Architecture C as a whole, approved only elements of it, or merely rejected alternatives;
- the exact final status of every EIC candidate, provisional functional shape, and retest result;
- whether EIC simulation version `v0.1.1` formally supersedes `v0.1` or is a parallel revision;
- whether a prototype exists in the repository and, if so, which file or implementation record controls prototype evidence;
- whether all human rulings have been reflected in v1.1.1;
- whether the project owner has formally approved the proposed metadata and audit protocol;
- whether `PROG-002` is intended to remain `DERIVED CONSTRAINT` or `PROPOSED` after later rulings;
- whether the active corpus has a complete rule-ID registry preventing identifier collisions;
- whether the merged document is intended for audit use only or is itself considered a project artifact;
- whether there are additional project files outside the 19 listed Markdown records.

These should remain explicitly unresolved rather than being filled with assumptions.

# Human Authorization Required

No file operation should be performed yet.

Before implementation, the project owner must approve:

- the active-document set;
- the Project State Manifest;
- the Decision and Change Register;
- the Open Questions Register;
- the System Package Index;
- the EIC current-state summary;
- stable document IDs and metadata;
- supersession and archive classifications;
- the deterministic LLM audit protocol;
- the standardized audit output and citation format;
- the precise scope of authority granted by each human-ruling record.

The proposed architecture preserves all existing information, does not promote any proposal into canon, does not demote any human-approved material, and does not require regenerating the Master GDD.