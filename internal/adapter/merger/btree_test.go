package merger

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/threez/intm/internal/model"
)

func TestBtree(t *testing.T) {
	suite.Run(t, &MergerTestSuite{Merger: NewBtree()})
}

func BenchmarkBtree(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		bt := NewBtree()
		for i := 0; i < 1000; i++ {
			bt.MergeInterval(model.NewInterval(25, 30))
			bt.MergeInterval(model.NewInterval(2, 19))
			bt.MergeInterval(model.NewInterval(14, 23))
			bt.MergeInterval(model.NewInterval(4, 8))
		}
		bt.Result()
	}
}
