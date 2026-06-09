# Changelog

## [Stage 5] — 2026-06-09

### Changed
- Replaced GPA with 4 exam scores per applicant (physics, chemistry, math, cs)
- Each department now ranks candidates by a single department-specific exam
- Updated `application_list.py` to Stage 5 format (98 students with exam scores)

## [Stage 4] — 2026-06-09

### Changed
- Full rewrite: reads `applicants.txt` instead of stdin for applicant data
- Implements 3-wave priority-based admission across 5 departments (Biotech, Chemistry, Engineering, Mathematics, Physics)
- Outputs each department (alphabetical order) with accepted applicants and GPA, sorted by GPA desc / name asc
- Updated `.gitignore` to exclude yaml, yml, gradle, and XML config files

## [Stage 3] — 2026-06-09

### Changed
- Full rewrite: reads N applicants with GPAs, ranks top M by GPA (desc), ties broken alphabetically by full name
- Output: `Successful applicants:` header followed by M names (no GPA)

## [Stage 2] — 2026-06-09

### Changed
- Added threshold check: mean ≥ 60 accepts, mean < 60 rejects with a different message
- Updated `tests.py` to Stage 2 test suite (6 cases covering both outcomes and boundary)

## [Stage 1] — 2026-06-09

### Added
- Initial solution: reads 3 integer exam scores, computes mean, prints result and acceptance message
- `CLAUDE.md` with project guidance and stage reference
- `README.md` with project overview, setup, and input/output documentation
- `.gitignore` covering Go binaries, IDE files, Python cache, and generated department output files
- Git repository initialized
