```
❯ make
go test -v ./mystuff
=== RUN   TestSomething
    main_test.go:11: setup test case
2020/11/03 14:08:43 Create from model: <*mystuff.FireTurtleV1> &{Name:Bob Shell:{Color:Blue}}
2020/11/03 14:08:43 Create from model: <*mystuff.FireUnicornV1> &{Name:Caroline}
    main_test.go:21: teardown test case
--- PASS: TestSomething (0.00s)
PASS
ok      github.com/yitsushi/dde-experiment/mystuff      (cached)
```
