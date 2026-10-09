# Go with GenAI — Lesson 6 homework

**Lesson 6: File I/O, JSON, and testing — a triad for production-ready Go code**

This repository implements three homework tasks. The package tests specify the expected behavior; the AI blind-testing report records the generated cases and their review. The mentor-review status is kept explicit in `ai-blind-testing/REPORT.md`.

| Folder | Task | Deliverable |
|---|---:|---|
| [`todo/`](todo) | 1 | `SaveTodos` / `LoadTodos` JSON persistence; at least 80% package coverage |
| [`validate/`](validate) | 2 | `ValidateEmail` / `ValidatePhone` with table-driven tests (at least 8 cases each) |
| [`ai-blind-testing/`](ai-blind-testing) | 3 | AI blind-test report, with mentor review tracked separately |

## Run and verify

Requires Go 1.22 or newer (`go version`). From the repository root, run:

```bash
go test -v -count=1 ./...  # all tests and subtests
go test -count=1 -cover ./...  # coverage by package
go vet ./...
```

Or use the Makefile:

```bash
make test-verbose
make cover
make vet
```

`todo.SaveTodos` writes indented JSON and creates or overwrites the target file. `todo.LoadTodos` returns contextual errors that preserve causes, including `os.ErrNotExist` and JSON syntax errors. The package tests exercise round trips, nil and empty slices, file creation/overwrite, and failure cases.

The validators intentionally implement scoped syntax checks, not full standards or service verification:

- Email accepts ASCII dot-atom local parts (maximum 64 bytes) and multi-label domains (maximum 63 bytes per label) with a final label of at least two letters; the full address is limited to 254 bytes. Whitespace, Unicode, quoted local parts, address literals, and malformed dots/hyphens are rejected.
- Phone accepts 8–15 ASCII digits with an optional leading `+`; separators and other characters are rejected.

See `validate/validate.go` for the exact rules and `validate/validate_test.go` for table-driven examples.

## Continuous integration

`.github/workflows/check-homework.yml` runs on pushes, pull requests, and manual dispatch. Task 1 and Task 2 test steps run independently so failures in one do not hide the other. CI also runs `go vet`, enforces at least 80% coverage for `todo`, checks that the AI report is at least 80 words, writes a job summary, and fails the overall job when a required check is incomplete. The AI report's word-count check does not constitute mentor approval; the mentor reviews its contents separately.

## Publish to your own GitHub repository

**Option A — GitHub CLI**

```bash
gh repo create <your-username>/go-with-genai-lesson6-homework --private --source=. --remote=origin --push
```

**Option B — manual setup**

```bash
git init
git add .
git commit -m "Complete Lesson 6 homework"
git branch -M main
git remote add origin https://github.com/<your-username>/go-with-genai-lesson6-homework.git
git push -u origin main
```

Then open the **Actions** tab to review the workflow run and its summary.

## Repository layout

```
.
├── .github/workflows/check-homework.yml
├── todo/
│   ├── todo.go
│   ├── todo_test.go
│   └── README.md
├── validate/
│   ├── validate.go
│   ├── validate_test.go
│   └── README.md
├── ai-blind-testing/
│   ├── REPORT.md
│   └── README.md
├── go.mod
├── Makefile
└── README.md
```

---

# Go з GenAI — домашнє завдання до уроку 6

**Урок 6: Робота з файлами, JSON і тестування — тріада для коду на Go, готового до промислової експлуатації**

У цьому репозиторії реалізовано три завдання. Тести пакетів визначають очікувану поведінку, а звіт про сліпу генерацію тестів ШІ фіксує згенеровані випадки та їх перевірку. Статус перевірки ментором зазначається окремо у `ai-blind-testing/REPORT.md`.

| Папка | Завдання | Результат |
|---|---:|---|
| [`todo/`](todo) | 1 | JSON-збереження через `SaveTodos` / `LoadTodos`; покриття пакета не нижче 80% |
| [`validate/`](validate) | 2 | `ValidateEmail` / `ValidatePhone` із табличними тестами (щонайменше 8 випадків для кожної функції) |
| [`ai-blind-testing/`](ai-blind-testing) | 3 | Звіт про сліпу генерацію тестів ШІ; перевірка ментором відстежується окремо |

## Запуск і перевірка

Потрібна версія Go 1.22 або новіша (`go version`). У корені репозиторію виконайте:

```bash
go test -v -count=1 ./...  # усі тести й підтести
go test -count=1 -cover ./...  # покриття для кожного пакета
go vet ./...
```

Або скористайтеся Makefile:

```bash
make test-verbose
make cover
make vet
```

`todo.SaveTodos` записує форматований JSON і створює або перезаписує файл. `todo.LoadTodos` повертає контекстні помилки, зберігаючи їхні першопричини, зокрема `os.ErrNotExist` та помилки синтаксису JSON. Тести пакета перевіряють round-trip, nil- і порожні списки, створення/перезапис файла та помилкові сценарії.

Валідатори навмисно перевіряють визначений набір синтаксичних правил, а не повну відповідність стандартам чи реальність даних:

- Email: ASCII local-part у форматі dot-atom (не більше 64 байтів) і домен із кількох частин (не більше 63 байтів на label), остання з яких має щонайменше дві літери; загальна довжина — до 254 байтів. Пробіли, Unicode, quoted local-part, адресні літерали та некоректні крапки/дефіси відхиляються.
- Телефон: 8–15 ASCII-цифр із необов'язковим `+` на початку; роздільники й інші символи відхиляються.

Точні правила наведено у `validate/validate.go`, приклади табличних тестів — у `validate/validate_test.go`.

## Безперервна інтеграція

`.github/workflows/check-homework.yml` запускається під час push, pull request і вручну. Тести Завдання 1 та Завдання 2 запускаються незалежно, тому помилки одного не приховують результати іншого. CI також запускає `go vet`, вимагає покриття `todo` щонайменше 80%, перевіряє, що звіт про ШІ містить не менше 80 слів, формує підсумок і завершує процес помилкою, якщо обов'язкова перевірка не пройдена. Перевірка кількості слів не є підтвердженням схвалення ментора — ментор окремо оцінює зміст звіту.

## Публікація у власному репозиторії GitHub

**Варіант А — GitHub CLI**

```bash
gh repo create <your-username>/go-with-genai-lesson6-homework --private --source=. --remote=origin --push
```

**Варіант Б — вручну**

```bash
git init
git add .
git commit -m "Complete Lesson 6 homework"
git branch -M main
git remote add origin https://github.com/<your-username>/go-with-genai-lesson6-homework.git
git push -u origin main
```

Після цього відкрийте вкладку **Actions**, щоб переглянути запуск workflow та його підсумок.

## Структура репозиторію

```
.
├── .github/workflows/check-homework.yml
├── todo/
│   ├── todo.go
│   ├── todo_test.go
│   └── README.md
├── validate/
│   ├── validate.go
│   ├── validate_test.go
│   └── README.md
├── ai-blind-testing/
│   ├── REPORT.md
│   └── README.md
├── go.mod
├── Makefile
└── README.md
```
