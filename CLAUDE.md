# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a [Hyperskill](https://hyperskill.org/projects/234) educational Go project: **University Admission Procedure**. It is structured as a 7-stage progressive project where each stage builds on the previous one.

The solution file lives at:
```
University Admission Procedure (Go)/task/main.go
```

## Running and Testing

Run the solution:
```bash
go run "University Admission Procedure (Go)/task/main.go"
```

Run the stage tests (uses Python `hstest` framework):
```bash
cd "University Admission Procedure (Go)/task"
python tests.py
```

Install the test framework (required once):
```bash
pip install https://github.com/hyperskill/hs-test-python/archive/release.tar.gz
```

## Stage Progression

Each stage is described in `task.html` within its subfolder. The stages build complexity incrementally:

| Stage | Folder | Core change |
|-------|--------|-------------|
| 1 | `No one is left behind` | Read 3 exam scores → print mean + always accept |
| 2 | `Raising the bar` | Add threshold: accept if mean ≥ 60, else reject |
| 3 | `Going big` | Read N applicants with GPA from stdin; accept top M; sort by GPA desc, then name asc |
| 4 | `Choose your path` | Read `applicants.txt`; 5 departments; 3-priority wave admission; sort by GPA |
| 5 | `Special knowledge` | Replace GPA with department-specific exam scores (physics/chemistry/math/cs) |
| 6 | `Extensive training` | Some departments use mean of multiple exams; write results to per-department `.txt` files |
| 7 | `Something special` | Add special admission exam column; rank by `max(mean finals score, special exam score)`; output to files |

## Applicants File Format

Stages 4–6 read from `applicants.txt` in the working directory.

- **Stages 4:** `FirstName LastName GPA Dept1 Dept2 Dept3`
- **Stages 5–6:** `FirstName LastName physics chemistry math cs Dept1 Dept2 Dept3`
- **Stage 7:** `FirstName LastName physics chemistry math cs specialExam Dept1 Dept2 Dept3`

## Department Exam Mapping (Stages 5–7)

| Department | Exams used |
|-----------|------------|
| Physics | physics + math (mean) |
| Chemistry | chemistry |
| Mathematics | math |
| Engineering | cs + math (mean) |
| Biotech | chemistry + physics (mean) |

## Admission Algorithm (Stages 4–7)

Three-wave admission: in each wave, only applicants with that wave's priority for a department compete. Accepted applicants are removed from the pool before the next wave. Each department accepts at most N students total across all waves. Ties in score are broken alphabetically by full name.

## Output Files (Stages 6–7)

Results are written to: `biotech.txt`, `chemistry.txt`, `engineering.txt`, `mathematics.txt`, `physics.txt`. Each line: `FirstName LastName score` (score as float with one decimal, e.g. `83.5`).
