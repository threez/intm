package merger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestBtree(t *testing.T) {
	suite.Run(t, &MergerTestSuite{Merger: NewBtree()})
}
