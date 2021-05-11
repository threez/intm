package merger

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/threez/intm/internal/model"
	"github.com/threez/intm/internal/port"
)

type MergerTestSuite struct {
	Merger port.Merger
	suite.Suite
}

func (suite *MergerTestSuite) TestAll() {
	assert.Equal(suite.T(), []*model.Interval(nil), suite.Merger.Result())
}

func (suite *MergerTestSuite) TestOne() {
	suite.Merger.MergeInterval(model.NewInterval(25, 30))
	assert.Equal(suite.T(), []*model.Interval{
		model.NewInterval(25, 30),
	}, suite.Merger.Result())
}

func (suite *MergerTestSuite) TestSimple() {
	suite.Merger.MergeInterval(model.NewInterval(25, 30))
	suite.Merger.MergeInterval(model.NewInterval(2, 19))
	suite.Merger.MergeInterval(model.NewInterval(14, 23))
	suite.Merger.MergeInterval(model.NewInterval(4, 8))
	assert.Equal(suite.T(), []*model.Interval{
		model.NewInterval(2, 23),
		model.NewInterval(25, 30),
	}, suite.Merger.Result())
}

func (suite *MergerTestSuite) TestExtended() {
	suite.Merger.MergeInterval(model.NewInterval(25, 30))
	suite.Merger.MergeInterval(model.NewInterval(2, 19))
	suite.Merger.MergeInterval(model.NewInterval(14, 23))
	suite.Merger.MergeInterval(model.NewInterval(4, 8))
	suite.Merger.MergeInterval(model.NewInterval(35, 40))
	suite.Merger.MergeInterval(model.NewInterval(1, 50))
	suite.Merger.MergeInterval(model.NewInterval(32, 34))
	assert.Equal(suite.T(), []*model.Interval{
		model.NewInterval(1, 50),
	}, suite.Merger.Result())
}

func (suite *MergerTestSuite) TestBad() {
	suite.Merger.MergeInterval(model.NewInterval(5, 6))
	suite.Merger.MergeInterval(model.NewInterval(3, 4))
	suite.Merger.MergeInterval(model.NewInterval(1, 2))
	suite.Merger.MergeInterval(model.NewInterval(4, 6))
	assert.Equal(suite.T(), []*model.Interval{
		model.NewInterval(1, 2),
		model.NewInterval(3, 6),
	}, suite.Merger.Result())
}

func (suite *MergerTestSuite) TestMoreBad() {
	suite.Merger.MergeInterval(model.NewInterval(50, 60))
	suite.Merger.MergeInterval(model.NewInterval(50, 55))
	suite.Merger.MergeInterval(model.NewInterval(45, 48))
	suite.Merger.MergeInterval(model.NewInterval(10, 24))
	suite.Merger.MergeInterval(model.NewInterval(10, 22))
	suite.Merger.MergeInterval(model.NewInterval(10, 20))
	suite.Merger.MergeInterval(model.NewInterval(30, 60))
	assert.Equal(suite.T(), []*model.Interval{
		model.NewInterval(10, 24),
		model.NewInterval(30, 60),
	}, suite.Merger.Result())
}
