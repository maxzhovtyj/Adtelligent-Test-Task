package concat

import (
	"testing"
)

var loremIpsum = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Maecenas non odio eget quam gravida laoreet vitae id est. 
Cras sit amet porta dui. Pellentesque at pulvinar ante. 
Pellentesque leo dolor, tristique a diam vel, posuere rhoncus ex. 
Mauris gravida, orci eu molestie pharetra, mi nibh bibendum arcu, in bibendum augue neque ac nulla. 
Phasellus consectetur turpis et neque tincidunt molestie. 
Vestibulum diam quam, sodales quis nulla eget, volutpat euismod mauris.
`

const SIZE = 10000

var strSlice = make([]string, SIZE)

func init() {
	for i := 0; i < len(strSlice); i++ {
		strSlice[i] = loremIpsum
	}
}

func BenchmarkConcatenationOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Basic(strSlice)
	}
	b.ReportAllocs()
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithJoin(strSlice)
	}
	b.ReportAllocs()
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithStringBuilder(strSlice, len(loremIpsum))
	}
	b.ReportAllocs()
}
