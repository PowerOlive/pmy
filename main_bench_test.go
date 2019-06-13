package main

import (
	"log"
	"os"
	"testing"

	pmy "github.com/relastle/pmy/src"
)

// BenchmarkLoadLargeRules2 calculates time elapsed to load large rule file
// and fetch them all
func BenchmarkLoadLargeRules2(b *testing.B) {
	RULE_COUNT := 10000
	filePath := "./test_pmy_rules_large.json"
	pmy.DumpDummyRulesJSON(filePath, RULE_COUNT)
	b.ResetTimer()
	out := pmy.Run(
		filePath,
		pmy.Input{
			BufferLeft:  "git abcdef tes abcdef  ",
			BufferRight: "",
		},
	)
	b.StopTimer()
	os.Remove(filePath)
	if out != "" {
		log.Fatal("output is not empty")
	}
	return
}
