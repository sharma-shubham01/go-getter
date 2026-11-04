# Go Fundamentals — Learning Notes & Examples

A hands-on, topic-by-topic walkthrough of the Go programming language, from arrays and slices up to concurrency patterns and an HTTP rate-limiting middleware.

Every topic is a small, self-contained `main.go` program. The explanation lives **inside the code as comments**, next to the line it describes, so each file works as both notes and a runnable example.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [How to Run the Examples](#how-to-run-the-examples)
- [Directory Structure](#directory-structure)
- [Topics](#topics)
  - [a. Arrays](#a-arrays)
  - [b. Slices](#b-slices)
  - [c. Strings](#c-strings)
  - [d. Pointers](#d-pointers)
  - [e. Functions](#e-functions)
  - [f. Defer](#f-defer)
  - [g. Types](#g-types)
  - [h. Methods](#h-methods)
  - [i. Interfaces](#i-interfaces)
  - [j. Errors](#j-errors)
  - [k. Standard Library APIs](#k-standard-library-apis)
  - [l. Context](#l-context)
  - [m. Concurrency](#m-concurrency)
  - [n. Sync Package](#n-sync-package)
  - [o. Rate Limiting Middleware](#o-rate-limiting-middleware)
- [Suggested Learning Path](#suggested-learning-path)
- [Roadmap](#roadmap)

---

## Prerequisites

- **Go 1.22 or newer.** The `serve mux` example uses method-based routing patterns (`"GET /hello"`) that were added in Go 1.22. Everything else works on older versions.
- No third-party dependencies. Only the standard library is used.

Check your version:

```bash
go version
```

---

## How to Run the Examples

The `golang/` directory has **no `go.mod`**. Each example is a standalone `package main`, so you run it by pointing `go run` at the file.

> **Heads-up:** directory names contain spaces and dots (e.g. `d. pointers/1. basics`), so **always quote the path**.

```bash
cd golang
go run "a. arrays/main.go"
go run "e. functions/5. closures/main.go"
go run "m. concurrency/6. workerpool/main.go"
```

Or `cd` into the directory first:

```bash
cd "j. errors/3. errors - is and as"
go run main.go
```

### Running the tests

One example ships with a unit test (a mocked `Writer` interface):

```bash
cd "i. interfaces/practice"
go test main.go main_test.go
```

### Detecting data races

Run any concurrency example with the race detector enabled:

```bash
go run -race "m. concurrency/1. basics/main.go"
```

### Examples that start a server

These block and listen on port `8080`. Run them in one terminal and send requests from another:

| Example | Try it |
|---|---|
| `k. api/6. http server` | `curl localhost:8080/hello` |
| `k. api/7. serve mux` | `curl localhost:8080/hello` |
| `o. rate limiting middleware` | `for i in {1..10}; do curl -s localhost:8080/api; echo; done` |

### Examples that touch files

Run these **from inside their own directory**, because they use relative file paths:

| Example | Reads | Writes |
|---|---|---|
| `k. api/1. io` | `letters.txt` | `writing.txt` |
| `k. api/2. encoding` | `output.json` | `output.json` |

---

## Directory Structure

```
golang/
├── a. arrays/
├── b. slices/
│   ├── basics/
│   └── slicing/
├── c. strings/
├── d. pointers/
│   ├── 1. basics/
│   ├── 2. uses/
│   └── 3. parameters/
├── e. functions/
│   ├── 1. basics/
│   ├── 2. variadic functions/
│   ├── 3. first class functions/
│   ├── 4. function type/
│   ├── 5. closures/
│   └── 6. anonymous functions/
├── f. defer/
│   ├── 1. basics/
│   ├── 2. rules/
│   └── 3. recover and defer/
├── g. types/
│   ├── 1. basics/
│   ├── 2. inheritance/
│   └── 3. iota/
├── h. methods/
│   ├── 1. basics/
│   ├── 2. methods/
│   └── 3. method-expression/
├── i. interfaces/
│   ├── basics/
│   ├── empty-interface/
│   ├── nil-interface/
│   └── practice/            # includes main_test.go
├── j. errors/
│   ├── 1. basics/
│   ├── 2. sentinel errors/
│   ├── 3. errors - is and as/
│   └── 4. panic/
├── k. api/
│   ├── 1. io/
│   ├── 2. encoding/
│   ├── 3. custom/
│   ├── 4. time/
│   ├── 5. http client/
│   ├── 6. http server/
│   └── 7. serve mux/
├── l. context/
├── m. concurrency/
│   ├── 1. basics/
│   ├── 2. goroutines/
│   ├── 3. channels/
│   ├── 4. deadlocks/
│   ├── 5. waitgroups/
│   ├── 6. workerpool/
│   ├── practice-buffered-1/
│   ├── practice-buffered-2/
│   ├── practice-buffered-unbuffered/
│   ├── practice-mutex/
│   ├── practice-semaphore/
│   └── practice-unbuffered-1/
├── n. sync package/
│   ├── 1. waitgroup/
│   ├── 2. mutex/
│   └── 3. others/
├── o. rate limiting middleware/
└── pending_items            # scratch list of topics still to cover
```

The letter prefixes (`a.`, `b.`, …) give the intended reading order, and numbered subfolders (`1.`, `2.`, …) give the order within a topic.

---

## Topics

### a. Arrays

📁 `a. arrays/`

- Arrays as fixed-size, contiguous, same-type collections
- Four ways to declare: basic, with elements, **sparse** (`[5]int{1, 2: 10, 50}`), and **implicit length** (`[...]int{...}`)
- 2D arrays, indexing, and mutation
- **Limitations:** the size is part of the type and cannot change, and it must be a constant. This is why slices are usually preferred.

### b. Slices

| Example | What it covers |
|---|---|
| `basics/` | Slices as wrappers over arrays; `len` vs `cap`; `make([]T, len, cap)`; `append` and the spread operator `...`; **nil vs empty slices** |
| `slicing/` | `s[x:y]` notation; the **ripple effect** of shared backing arrays; how `append` can overwrite the parent or reallocate; full slice expression `s[low:high:max]`; `copy(dst, src)` |

### c. Strings

📁 `c. strings/`

- Strings are immutable sequences of bytes
- Indexing returns a **byte**, and slicing works on strings
- Conversions between `string`, `rune`, and `byte`

### d. Pointers

| Example | What it covers |
|---|---|
| `1. basics/` | Stack vs heap, the garbage collector, the `*` and `&` operators, dereferencing, the nil zero value |
| `2. uses/` | Three reasons to use pointers: **mutability**, signalling **"no value"** with `nil`, and **performance** with large structs |
| `3. parameters/` | Go is **strictly call-by-value**: even a pointer is copied. Why maps and slices *look* like they're passed by reference, and why `append` inside a function doesn't affect the caller's slice |

### e. Functions

| Example | What it covers |
|---|---|
| `1. basics/` | Declaring functions, multiple parameters, multiple return values, void functions |
| `2. variadic functions/` | `func f(nums ...int)`, passing a slice with `nums...`, and why Go has no default parameters |
| `3. first class functions/` | Functions as values: assigning them to variables, passing them as arguments, returning them |
| `4. function type/` | Naming a signature with `type operation func(int, int) int` to improve readability |
| `5. closures/` | Capturing surrounding variables; a stateful `counter()` closure |
| `6. anonymous functions/` | Immediately invoked function literals and inline callbacks |

### f. Defer

| Example | What it covers |
|---|---|
| `1. basics/` | When deferred calls run (on return, at the end of the body, or on panic); **LIFO** order; arguments are evaluated immediately but the call runs later; typical cleanup uses |
| `2. rules/` | **Rule 1:** arguments are captured when the `defer` line runs, not when the call executes. **Rule 2:** a deferred function can modify **named return values**, but not unnamed ones (`namedReturn()` returns 11, `unnamedReturn()` returns 1) |
| `3. recover and defer/` | Deferred functions still run while a panic unwinds the stack; `recover()` only works when called **directly** inside a deferred function |

### g. Types

| Example | What it covers |
|---|---|
| `1. basics/` | User-defined types with `type`, abstract vs concrete types, derived types (`type Age int`) |
| `2. inheritance/` | Go has **no inheritance**. A derived type and its base type are distinct, and code reuse comes from composition |
| `3. iota/` | Enumerated constants with `iota` |

### h. Methods

| Example | What it covers |
|---|---|
| `1. basics/` | Method syntax; **value receivers vs pointer receivers**; no method overloading |
| `2. methods/` | Method values: passing `c.Add` where a function is expected, and storing a method in a variable |
| `3. method-expression/` | Method expressions such as `Person.GetDetails` and `(*Rectangle).SetLength`, where the receiver becomes the first argument |

### i. Interfaces

| Example | What it covers |
|---|---|
| `basics/` | Implicit interface satisfaction, polymorphism through `[]Shape`, **interface embedding** |
| `empty-interface/` | `interface{}` / `any`; **type assertion** with comma-ok; **type switches** |
| `nil-interface/` | Why an interface holding a nil `*int` is **not** `== nil` (the interface stores a type and a value pointer) |
| `practice/` | Polymorphism with an `Animal` interface, a custom `error` type, and a `Writer` interface **mocked in a unit test** |

### j. Errors

| Example | What it covers |
|---|---|
| `1. basics/` | The `error` interface, Go conventions for errors, custom error structs |
| `2. sentinel errors/` | Sentinel errors, wrapping with `fmt.Errorf("%w")`, `errors.Join`, `errors.Unwrap` |
| `3. errors - is and as/` | `errors.Is` to match an error **instance** and `errors.As` to match an error **type** (e.g. `*fs.PathError`) |
| `4. panic/` | `panic`, how deferred functions still run during a panic, and `recover` |

### k. Standard Library APIs

| Example | What it covers |
|---|---|
| `1. io/` | `io.Reader` and `io.Writer`; one generic function that reads from both a file and a `strings.Reader` |
| `2. encoding/` | `json.Marshal`/`Unmarshal`; struct tags (`omitempty`, `-`); streaming with `json.Encoder`/`Decoder` |
| `3. custom/` | Implementing `json.Marshaler`/`Unmarshaler` to serialize an `iota` enum as a string |
| `4. time/` | Wall-clock vs monotonic time; `time.Time` and `time.Duration`; `Add`, `Since`, and comparisons |
| `5. http client/` | `http.Client` with a timeout, `NewRequestWithContext`, decoding a JSON response |
| `6. http server/` | `http.ListenAndServe`, the `Handler` interface, `ResponseWriter`, and `Request` |
| `7. serve mux/` | `http.NewServeMux`, method-and-path patterns, path precedence rules, `PathValue` |

### l. Context

📁 `l. context/`

A notes file on `context.Context`:

- Its four methods: `Deadline`, `Done`, `Err`, and `Value`
- Common uses: cancellation, deadlines, and request-scoped values
- Conventions: create it at the root, pass it as the first parameter, and never store or mutate it

### m. Concurrency

**Core lessons**

| Example | What it covers |
|---|---|
| `1. basics/` | Concurrency vs parallelism, race conditions, atomicity, critical sections, `sync.Mutex` |
| `2. goroutines/` | Notes on OS threads, green threads, goroutines, and the Go scheduler |
| `3. channels/` | Unbuffered, buffered, send-only, receive-only, and bidirectional channels |
| `4. deadlocks/` | What causes deadlocks; closing channels and ranging over them |
| `5. waitgroups/` | `sync.WaitGroup` with `Add`, `Done`, and `Wait` |
| `6. workerpool/` | The **worker pool** pattern: a job queue, N workers, a results channel, and a WaitGroup |

**Practice problems**

| Example | Problem |
|---|---|
| `practice-buffered-1/` | Merge two closed channels into one channel |
| `practice-buffered-2/` | Square integers using a fixed pool of workers |
| `practice-buffered-unbuffered/` | Print `ABCABCABC` with three goroutines, each printing one letter, without sleeps or counter checks |
| `practice-unbuffered-1/` | Two goroutines print even and odd numbers in strict order, using channels and also using `select` |
| `practice-mutex/` | A thread-safe **round-robin load balancer**, built with `sync.RWMutex` and with `sync/atomic` |
| `practice-semaphore/` | Process 50 jobs while allowing at most 5 to run at once, using a buffered channel as a **semaphore** |

### n. Sync Package

A closer look at the synchronization primitives in the `sync` package.

| Example | What it covers |
|---|---|
| `1. waitgroup/` | `Add`, `Done`, and `Wait`. Common pitfalls: calling `Add` inside the goroutine, copying a `WaitGroup` by value, and a negative counter panic |
| `2. mutex/` | `sync.Mutex` with a `SafeCounter` and the idiomatic `defer Unlock()`; deadlock causes; normal vs. starvation mode. `sync.RWMutex` with a read-heavy `Cache`, when to use `RLock` vs `Lock`, and writer priority |
| `3. others/` | `sync.Map` and when it beats a mutex-guarded map; `sync.Once` for lazy singletons; an introduction to `sync.Pool` and `sync.Cond` |

### o. Rate Limiting Middleware

📁 `o. rate limiting middleware/`

A small but complete HTTP server that applies the **token bucket** algorithm as middleware:

- The bucket holds at most `maxTokens` tokens (a burst of 5)
- A background `time.Ticker` goroutine adds one token every `refillHz` (500 ms)
- Each request uses one token. When the bucket is empty, the server responds with **`429 Too Many Requests`**
- A `sync.Mutex` protects the shared token count

```bash
go run "o. rate limiting middleware/main.go"
# in another terminal: the first 5 requests succeed, the rest get 429
for i in {1..10}; do curl -s localhost:8080/api; echo; done
```

This example brings together earlier topics: closures, methods, goroutines, `sync.Mutex`, `time`, and `net/http`.

---

## Suggested Learning Path

```
Basics          a. arrays → b. slices → c. strings → d. pointers
Functions       e. functions → f. defer
Type system     g. types → h. methods → i. interfaces → j. errors
Std library     k. api → l. context
Concurrency     m. concurrency (core 1–6, then practice-*) → n. sync package
Capstone        o. rate limiting middleware
```

---

## Roadmap

Topics still to cover, from `pending_items`:

- [ ] Semaphores (deep dive)
- [x] Mutex (deep dive) — covered in `n. sync package/2. mutex`

Other natural next steps: maps and structs, generics, `select` with timeouts, `context` cancellation in practice, `errgroup`, and table-driven tests.
