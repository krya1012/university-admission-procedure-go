# University Admission Procedure (Go)

A [Hyperskill](https://hyperskill.org/projects/234) project implementing a university admission algorithm in Go. The algorithm is built incrementally across 7 stages, growing from a simple score calculator to a full multi-department admission system with priority queues, department-specific exam weighting, and file output.

## Project Structure

```
University Admission Procedure (Go)/
└── task/
    ├── main.go       # solution (edit this)
    └── tests.py      # stage test suite
```

Each stage has its own subfolder with a `task.html` describing the requirements:

| # | Stage | Description |
|---|-------|-------------|
| 1 | No one is left behind | Read 3 exam scores → print mean + always accept |
| 2 | Raising the bar | Accept if mean ≥ 60, reject otherwise |
| 3 | Going big | Rank N applicants by GPA, admit top M |
| 4 | Choose your path | 5 departments, 3-priority wave admission from `applicants.txt` |
| 5 | Special knowledge | Department-specific exam scores instead of GPA |
| 6 | Extensive training | Multi-exam means per department; output to `.txt` files |
| 7 | Something special | Special admission exam; use `max(finals mean, special score)` |

## Requirements

- Go 1.18+
- Python 3 (for tests)
- `hs-test-python` test framework:
  ```bash
  pip install https://github.com/hyperskill/hs-test-python/archive/release.tar.gz
  ```

## Running

```bash
# Run the current solution (Stage 1 example)
echo -e "70\n90\n60" | go run "University Admission Procedure (Go)/task/main.go"
```

## Testing

```bash
cd "University Admission Procedure (Go)/task"
python3 tests.py
```

## Input / Output

### Stages 1–2 (stdin)
```
75
90
68
```
```
77.66666666666667
Congratulations, you are accepted!
```

### Stages 4–7 (reads `applicants.txt`)

Applicant file format evolves per stage:

| Stage | Format |
|-------|--------|
| 4 | `FirstName LastName GPA Dept1 Dept2 Dept3` |
| 5–6 | `FirstName LastName physics chemistry math cs Dept1 Dept2 Dept3` |
| 7 | `FirstName LastName physics chemistry math cs specialExam Dept1 Dept2 Dept3` |

Departments: `Biotech`, `Chemistry`, `Engineering`, `Mathematics`, `Physics`

Stages 6–7 write results to: `biotech.txt`, `chemistry.txt`, `engineering.txt`, `mathematics.txt`, `physics.txt`
