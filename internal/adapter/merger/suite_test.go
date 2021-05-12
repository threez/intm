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

func (suite *MergerTestSuite) TestEmpty() {
	assert.Equal(suite.T(), []*model.Interval(nil), suite.Merger.Result())
}

func (suite *MergerTestSuite) test(in, out [][2]int) {
	for _, v := range in {
		suite.Merger.MergeInterval(model.NewInterval(v[0], v[1]))
	}
	expected := make([]string, len(out))
	for i, v := range out {
		expected[i] = model.NewInterval(v[0], v[1]).String()
	}
	actualResult := suite.Merger.Result()
	actual := make([]string, len(actualResult))
	for i, v := range actualResult {
		actual[i] = v.String()
	}
	assert.Equal(suite.T(), expected, actual)
}

func (suite *MergerTestSuite) TestOne() {
	suite.test([][2]int{
		{25, 30},
	}, [][2]int{
		{25, 30},
	})
}

func (suite *MergerTestSuite) TestSimple() {
	suite.test([][2]int{
		{25, 30},
		{2, 19},
		{14, 23},
		{4, 8},
	}, [][2]int{
		{2, 23},
		{25, 30},
	})
}

func (suite *MergerTestSuite) TestExtended() {
	suite.test([][2]int{
		{25, 30},
		{2, 19},
		{14, 23},
		{4, 8},
		{35, 40},
		{1, 50},
		{32, 34},
	}, [][2]int{
		{1, 50},
	})
}

func (suite *MergerTestSuite) TestBad() {
	suite.test([][2]int{
		{5, 6},
		{3, 4},
		{1, 2},
		{4, 6},
	}, [][2]int{
		{1, 2},
		{3, 6},
	})
}

func (suite *MergerTestSuite) TestMoreBad() {
	suite.test([][2]int{
		{50, 60},
		{10, 24},
		{50, 55},
		{10, 20},
		{45, 48},
		{10, 22},
		{30, 60},
	}, [][2]int{
		{10, 24},
		{30, 60},
	})
}
