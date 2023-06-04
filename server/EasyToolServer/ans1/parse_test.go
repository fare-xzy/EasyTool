package ans1

import (
	"os"
	"testing"
)

func TestParseAsn1(t *testing.T) {
	file, _ := os.ReadFile("C:\\Users\\12691\\Desktop\\1.dat")
	ParseAsn1(file)
}
