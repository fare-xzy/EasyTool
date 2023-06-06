package ans1

import (
	"os"
	"testing"
)

func TestParseAsn1(t *testing.T) {
	file, _ := os.ReadFile("E:\\IDEA\\Web\\Mine\\EasyTool\\server\\EasyToolServer\\resources\\GB.dat")
	ParseAsn1(file)
}
