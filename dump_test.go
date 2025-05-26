package debug

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// stripANSI removes ANSI color codes for testable output.
func stripANSI(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(s, "")
}

func TestSimpleStruct(t *testing.T) {
	type Profile struct {
		Age   int
		Email string
	}
	type User struct {
		Name    string
		Profile Profile
	}

	user := User{Name: "Alice", Profile: Profile{Age: 30, Email: "alice@example.com"}}
	out := stripANSI(DumpStr(user))

	assert.Contains(t, out, "#debug.User")
	assert.Contains(t, out, "+Name")
	assert.Contains(t, out, "\"Alice\"")
	assert.Contains(t, out, "+Profile")
	assert.Contains(t, out, "#debug.Profile")
	assert.Contains(t, out, "+Age")
	assert.Contains(t, out, "30")
	assert.Contains(t, out, "+Email")
	assert.Contains(t, out, "alice@example.com")
}

func TestNilPointer(t *testing.T) {
	var s *string
	out := stripANSI(DumpStr(s))
	assert.Contains(t, out, "(nil)")
}

func TestCycleReference(t *testing.T) {
	type Node struct {
		Next *Node
	}
	n := &Node{}
	n.Next = n
	out := stripANSI(DumpStr(n))
	assert.Contains(t, out, "↩︎ &1")
}

func TestMaxDepth(t *testing.T) {
	type Node struct {
		Child *Node
	}
	n := &Node{}
	curr := n
	for i := 0; i < 20; i++ {
		curr.Child = &Node{}
		curr = curr.Child
	}
	out := stripANSI(DumpStr(n))
	assert.Contains(t, out, "... (max depth)")
}

func TestMapOutput(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	out := stripANSI(DumpStr(m))

	assert.Contains(t, out, "a => 1")
	assert.Contains(t, out, "b => 2")
}

func TestSliceOutput(t *testing.T) {
	s := []string{"one", "two"}
	out := stripANSI(DumpStr(s))

	assert.Contains(t, out, "0 => \"one\"")
	assert.Contains(t, out, "1 => \"two\"")
}

func TestAnonymousStruct(t *testing.T) {
	out := stripANSI(DumpStr(struct{ ID int }{ID: 123}))

	assert.Contains(t, out, "+ID")
	assert.Contains(t, out, "123")
}

func TestControlCharsEscaped(t *testing.T) {
	s := "line1\nline2\tok"
	out := stripANSI(DumpStr(s))
	assert.Contains(t, out, `\n`)
	assert.Contains(t, out, `\t`)
}

func TestFuncPlaceholder(t *testing.T) {
	fn := func() {}
	out := stripANSI(DumpStr(fn))
	assert.Contains(t, out, "func(...) {...}")
}
