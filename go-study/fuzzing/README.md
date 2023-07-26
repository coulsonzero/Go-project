# Go Fuzzing

```shell
$ go test

$ go test -fuzz=Fuzz

$ go test -fuzz=Fuzz -fuzztime 30s
```

## fuzz

```shell
$ go test
PASS
ok      go-study/src/fuzzing/fuzz       0.084s

$ go test -fuzz=-Fuzz
PASS
ok      go-study/src/fuzzing/fuzz       0.345s
```


## fuzz2
```shell
$ go test
input: "Hello, world"
runes: ['H' 'e' 'l' 'l' 'o' ',' ' ' 'w' 'o' 'r' 'l' 'd']
input: "dlrow ,olleH"
runes: ['d' 'l' 'r' 'o' 'w' ' ' ',' 'o' 'l' 'l' 'e' 'H']
input: " "
runes: [' ']
input: " "
runes: [' ']
input: "!12345"
runes: ['!' '1' '2' '3' '4' '5']
input: "54321!"
runes: ['5' '4' '3' '2' '1' '!']
PASS
ok      go-study/src/fuzzing/fuzz2      0.459s


$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
fuzz: minimizing 31-byte failing input file
fuzz: elapsed: 0s, minimizing
--- FAIL: FuzzReverse (0.03s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=1, doubleRev=1
        reverse_test.go:18: Before: "\xd1", after: "�"
    
    Failing input written to testdata/fuzz/FuzzReverse/9a37ba4994d2f5a4db95868e179b29266d973c2ec7dbe57c44593ac907920eb5
    To re-run:
    go test -run=FuzzReverse/9a37ba4994d2f5a4db95868e179b29266d973c2ec7dbe57c44593ac907920eb5
FAIL
exit status 1
FAIL    go-study/src/fuzzing/fuzz2      0.473s
```

## fuzz3

```shell
$ go test
PASS
ok      go-study/src/fuzzing/fuzz3      0.305s

$ go test -fuzz=Fuzz -fuzztime 30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/39 completed
fuzz: elapsed: 0s, gathering baseline coverage: 39/39 completed, now fuzzing with 8 workers
fuzz: elapsed: 3s, execs: 987468 (329106/sec), new interesting: 6 (total: 45)
fuzz: elapsed: 6s, execs: 2014684 (342432/sec), new interesting: 6 (total: 45)
fuzz: elapsed: 9s, execs: 2994630 (326643/sec), new interesting: 7 (total: 46)
fuzz: elapsed: 12s, execs: 3982486 (329272/sec), new interesting: 7 (total: 46)
fuzz: elapsed: 15s, execs: 4905396 (307621/sec), new interesting: 7 (total: 46)
fuzz: elapsed: 18s, execs: 5872856 (322506/sec), new interesting: 7 (total: 46)
fuzz: elapsed: 21s, execs: 6872904 (333250/sec), new interesting: 8 (total: 47)
fuzz: elapsed: 24s, execs: 7897594 (341706/sec), new interesting: 8 (total: 47)
fuzz: elapsed: 27s, execs: 8914129 (338827/sec), new interesting: 8 (total: 47)
fuzz: elapsed: 30s, execs: 9936177 (340639/sec), new interesting: 8 (total: 47)
fuzz: elapsed: 30s, execs: 9936177 (0/sec), new interesting: 8 (total: 47)
PASS
ok      go-study/src/fuzzing/fuzz3      30.193s
```