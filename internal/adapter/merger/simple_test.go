package merger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSimple(t *testing.T) {
	suite.Run(t, &MergerTestSuite{Merger: NewSimple()})
}
