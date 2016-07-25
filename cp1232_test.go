package cp1232

import (
	"testing"
)

func Test1(t *testing.T) {
	text := FromUtf8("“Não existe almoço grátis”")
	expected := "\x93N\xE3o existe almo\xE7o gr\xE1tis\x94"

	if text != expected {
		t.Fatalf("\nexpected [%s]\n got     [%s].", expected, text)
	}
}
