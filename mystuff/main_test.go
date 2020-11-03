package mystuff_test

import (
	"testing"

	"github.com/yitsushi/dde-experiment/dde"
	"github.com/yitsushi/dde-experiment/mystuff"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")

	env := dde.NewEnvironment()

	env.Register("FireTurtleV1", mystuff.FireTurtleV1{})
	env.Register("FireUnicornV1", mystuff.FireUnicornV1{})

	env.Create("scenario")

	return func(t *testing.T) {
		t.Log("teardown test case")
		env.Destroy()
	}
}

func TestSomething(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
}
