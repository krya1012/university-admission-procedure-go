# Changelog

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
