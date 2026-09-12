# PROJECT DOCUMENTATION CONSOLIDATION AND CROSS-LLM AUDITABILITY ASSESSMENT

**Author:** Author LLM Unknown (process provenance; filed 2026-09-12)

## Role

Act as an independent documentation architect and evidence/governance auditor specialising in long-lived complex projects and multi-LLM documentation workflows.

Your primary concern is cross-LLM auditability and reproducibility of project state, not merely reducing document count.

Do not act as the project's design authority.

Do not make project decisions.

Do not promote proposals into decisions.

Do not demote approved material.

Do not resolve unresolved questions unless the supplied evidence explicitly establishes the answer.

---

## Context

The supplied corpus contains documentation belonging to a project.

The documentation may include current specifications, canonical documents, governance records, audits, proposals, research, experiments, evidence, decisions, meeting records, historical documents, superseded versions, duplicates, and other project artifacts.

The project currently has too many documents for reliable human and LLM consumption.

The entire documentation corpus may be supplied as a single merged document containing multiple original files. The merged corpus must therefore be treated as a collection of distinct source documents, not as one authoritative document.

The objective is to analyse the existing documentation and recommend a durable documentation architecture that:

1. reduces unnecessary document sprawl;
2. preserves historical information and provenance;
3. makes current project state immediately discoverable;
4. clearly separates authority from inference;
5. allows independent LLMs to audit the project consistently;
6. allows audits produced by different LLMs to be meaningfully compared.

No restructuring, renaming, deletion, merging, or modification of files is authorised by this prompt.

The output is a proposal and assessment only.

---

## Primary Objective

Design a documentation system that allows independent LLMs to reconstruct the same project state from the same evidence and produce audits that can be meaningfully compared.

The primary success criterion is therefore:

> Different LLMs should be able to distinguish project fact, approved decisions, derived conclusions, proposals, unresolved questions, evidence, historical material, and assessor opinion without relying on prior conversational knowledge.

---

## Secondary Objectives

1. Create a durable documentation architecture that can scale as the project grows.
2. Reduce unnecessary duplication and documentation sprawl.
3. Maintain a small, clearly identifiable active documentation set.
4. Preserve historical documents and provenance without allowing them to contaminate current project state.
5. Make authority, status, decisions, unresolved questions, evidence, and historical proposals machine/LLM-discernible.
6. Establish a repeatable audit method that can be used by different LLMs.

---

## Audience

The primary consumers are other LLMs that may have no prior knowledge of the project.

Human readability remains important, but the documentation must first allow an unfamiliar LLM to determine:

1. What the project currently is.
2. What has actually been approved or established.
3. What is derived or inferred.
4. What is merely proposed.
5. What remains unresolved.
6. What evidence exists.
7. What has been superseded.
8. Where important claims originated.
9. What changed between project states.
10. What an independent auditor should verify.

The project domain may be software, games, research, business, engineering, creative work, organisational work, or another field.

Do not assume a particular domain.

Infer the project's domain, terminology, governance model, and documentation conventions from the supplied corpus.

---

## Author / Assessor Attribution

Treat authorship and assessment identity as important provenance information.

Where the corpus establishes them, identify:

- Author — the person or entity that originally created the document.
- Assessor — the person, organisation, or LLM responsible for an audit or assessment.
- LLM / Model — the model used to produce an audit or assessment, if identifiable.
- Assessment Date — when the audit or assessment was produced.
- Document Date — when the source document was produced.
- Version — the document version, where applicable.

Do not infer authorship from writing style, vocabulary, or apparent expertise.

If authorship or assessment identity cannot be established, record:

**Unknown / Not Established**

Do not attribute a document to the LLM performing this consolidation merely because the document appears to have been LLM-generated.

Where possible, distinguish:

- original human author;
- later editor;
- assessor;
- LLM/model;
- project owner or approving authority.

Do not collapse these roles into a single "author" field.

---

## Required Initial Analysis

First analyse the supplied corpus.

Do not immediately propose a new structure.

Identify:

- current documents;
- superseded documents;
- duplicate documents;
- canonical/current specifications;
- governance documents;
- audits and assessments;
- provenance records;
- evidence/prototype/research records;
- decisions and decision records;
- active proposals;
- unresolved questions;
- historical material;
- potentially conflicting authority;
- redundant material;
- information appearing in multiple places;
- important information that is difficult for an independent LLM to locate.

Treat the merged input as a collection of source artifacts.

Recover document boundaries and metadata wherever possible.

For each significant source document, determine:

- title;
- version;
- date;
- document type;
- Author;
- Assessor, where applicable;
- LLM / Model, where applicable;
- authority/status;
- current or historical state;
- relationships to other documents;
- unique information contributed;
- supersession relationships;
- important dependencies.

---

## Authority and Status Analysis

Determine what authority/status vocabulary the project already uses.

Examples might include:

- approved;
- canonical;
- locked;
- proposed;
- draft;
- tentative;
- derived;
- assumed;
- experimental;
- validated;
- evidence;
- deferred;
- rejected;
- superseded;
- historical.

These are examples only.

Do not impose them on a project that uses different terminology.

Preserve the project's existing vocabulary wherever practical.

Identify the rules governing how material gains, loses, or retains authority.

Determine whether the documentation clearly distinguishes:

- explicit human/project decisions;
- project principles;
- logical consequences;
- LLM-generated recommendations;
- repeated proposals;
- evidence;
- prototype behaviour;
- historical references;
- assumptions;
- auditor conclusions.

If the project does not clearly distinguish these categories, identify that as a documentation/governance risk rather than inventing a hierarchy.

---

## Documentation Architecture Assessment

Assess the existing documentation system against:

### A. State Reconstruction

Can an independent LLM determine the current project state without reading every historical document?

### B. Authority Reconstruction

Can an independent LLM distinguish approved decisions from:

- proposals;
- logical deductions;
- historical material;
- evidence;
- assumptions;
- audit recommendations?

### C. Temporal Reconstruction

Can an LLM determine what is current versus superseded?

### D. Provenance

Can important current statements be traced to their origin and subsequent status changes?

### E. Cross-LLM Comparability

Could two independent LLMs inspect the same active documentation and produce comparable audit results?

### F. Change Detection

Can an auditor identify what changed since the previous project state?

### G. Historical Containment

Does the structure prevent historical or proposed material from silently becoming current project state?

### H. Scalability

Will the architecture continue working as the project grows?

### I. Attribution

Can an auditor determine who authored relevant source documents and who produced previous assessments?

### J. Audit Independence

Can a new assessor understand the project without relying on previous LLM conversations or undocumented assumptions?

---

## Recommended Architecture

After analysing the corpus, recommend a moderately consolidated documentation architecture.

Do not advocate either extreme:

- one enormous document containing everything; or
- dozens of narrowly divided documents.

The proposed architecture should contain a small active set plus a preserved historical/provenance area.

For every proposed active document, specify:

- filename/title;
- purpose;
- authority;
- primary audience;
- what belongs inside it;
- what must not belong inside it;
- relationship to other active documents;
- whether it is authoritative;
- update frequency;
- whether an LLM should normally read it during an audit.

Also identify which existing documents should become:

- active;
- superseded;
- archive/provenance;
- audit records;
- evidence records;
- candidates for eventual retirement.

Do not actually move, rename, merge, delete, or modify anything.

---

## Cross-LLM Auditability Requirements

The proposed system should allow an independent LLM to perform a standardized audit using a deterministic reading order.

Recommend a minimal audit protocol including:

1. Which documents to read first.
2. How to identify the current state.
3. How to identify authority.
4. How to distinguish current facts from proposals.
5. How to identify unresolved questions.
6. How to identify evidence.
7. How to identify superseded information.
8. How to inspect changes.
9. When historical/provenance material should be consulted.
10. How to report uncertainty.
11. How to cite evidence.
12. How to identify the Author / Assessor of relevant source material.
13. How to produce a standardized audit result.

Recommend whether a stable project-state manifest or equivalent entry-point document should exist and what fields it should contain.

---

## Standardization

Recommend a standardized metadata model for important project-state information where appropriate.

Potential fields include:

- document identity;
- title;
- document type;
- version;
- status;
- date;
- Author;
- Assessor;
- LLM / Model;
- approval authority;
- authority class;
- provenance;
- controlling document;
- supersedes;
- superseded-by;
- change identifier;
- current-state indicator;
- evidence status;
- open questions;
- dependencies.

Do not introduce unnecessary bureaucracy.

The objective is reliable state reconstruction and auditability, not documentation for its own sake.

---

## Audit Output Standardization

Recommend a standardized structure for future project status audits.

The structure should allow different LLM-produced audits to be compared field-by-field.

Separate:

**Project State**

from:

**Assessor Analysis**

from:

**Assessor Recommendation**

from:

**Uncertainty / Evidence Gap**

Do not allow an assessor's recommendation to appear to be a project decision.

---

## Success Criteria

A great solution will satisfy these three tests:

### Test 1 — State Reconstruction

A new LLM with no prior project context can determine the current project state from the active documentation without reading the entire historical archive.

### Test 2 — Audit Reproducibility

Two independent LLMs given the same active documentation and the same audit instructions should produce substantially comparable findings.

### Test 3 — Historical and Attribution Containment

Historical material, superseded decisions, proposals, evidence, or assessor recommendations cannot accidentally become current project authority merely because an LLM encounters them.

The system must also preserve attribution sufficiently to determine, where established, who originally authored source material and who produced previous assessments.

---

## Constraints

- Do not make project-design decisions.
- Do not resolve unresolved project questions.
- Do not silently promote proposals.
- Do not silently demote approved material.
- Do not delete historical information.
- Do not assume LLM consensus equals project approval.
- Do not treat prototype/evidence as design authority unless the project's own governance explicitly says otherwise.
- Do not treat a document as authoritative merely because it calls itself canonical.
- Do not impose terminology from another project.
- Do not assume all documents have the same authority.
- Do not infer authorship from writing style.
- Do not confuse Author, Assessor, approving authority, editor, or LLM/model.
- Do not create unnecessary documents merely to make the architecture appear sophisticated.
- Do not action any file operations.
- All proposed restructuring requires later human authorization.

---

## Examples

No external project-specific examples should override the supplied corpus.

If the supplied project has an established authority/status vocabulary, use that vocabulary.

If no established vocabulary exists, explicitly identify the gap and recommend a minimal model rather than pretending one already exists.

---

## Output Format

Produce the following sections:

1. Assessment Metadata
2. Executive Verdict
3. Current Documentation Assessment
4. Current Documentation Inventory
5. Redundancy and Sprawl Analysis
6. Authority and Provenance Risks
7. Cross-LLM Auditability Assessment
8. Recommended Documentation Architecture
9. Active Documentation Set
10. Archive / Provenance Structure
11. LLM Audit Reading Protocol
12. Recommended Standard Metadata
13. Recommended Audit Output Standard
14. Migration / Consolidation Plan
15. Risks and Trade-offs
16. Human Authorization Gates
17. Target State
18. Open Questions / Information Not Established by the Corpus

---

## Assessment Metadata

At the beginning of the output, provide:

| Field | Value |
|---|---|
| Assessment / Audit Title | |
| Assessor | |
| Assessor Type | Human / LLM / Organisation / Unknown |
| LLM / Model | If identifiable |
| Assessment Date | |
| Source Corpus | |
| Corpus Date / State | If established |
| Project | |
| Project Domain | |
| Confidence | |

If a field cannot be established, write:

**Unknown / Not Established**

Do not guess.

---

## Active Documentation Table

Provide:

| Document | Purpose | Authority | Status | Author / Owner | Primary Consumer |
|---|---|---|---|---|---|

---

## Evidence Discipline

For every significant conclusion:

- distinguish direct evidence from inference;
- identify the relevant source document;
- preserve the project's terminology;
- explicitly say when the corpus does not establish something;
- do not fill documentation gaps with invented assumptions;
- preserve source attribution where available.

When different documents conflict, report the conflict rather than choosing a winner unless the project's documented authority rules establish the winner.

When authorship conflicts, report the attribution conflict rather than selecting an author based on style or probability.

---

## Migration Plan

Provide a staged consolidation plan.

The plan must be proposal-only.

No files are to be changed.

Each proposed migration should identify:

- source document(s);
- proposed destination;
- information to preserve;
- information to deduplicate;
- provenance that must be retained;
- authority/status implications;
- validation required;
- human approval required.

---

## Human Authorization Required

End with a concise section identifying exactly which proposed structural changes require human/project-owner approval before implementation.

Do not perform those changes.

---

## Final Requirement

The final recommendation must optimize for:

1. Cross-LLM auditability and comparability.
2. Durable project documentation architecture.
3. Reduction of unnecessary document sprawl.

Do not optimize for minimum document count at the expense of any of these objectives.