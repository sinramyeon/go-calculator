package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// 각 테스트를 위한 구조체
type testCase struct {
	value1Float32, value2Float32, expectedResultFloat32 float32
}

// 테스트를 위한 테스트 수트
type CalculatorTestSuite struct {
	suite.Suite
	testSuccessCases []testCase
	testFailureCases []testCase
}

// 관련된 모든 테스트를 실행
func Test_add_suite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

// 테스트를 위한 값 설정
func (suite *CalculatorTestSuite) SetupTest() {
	suite.testSuccessCases = []testCase{
		{10, 20, 30},
		{20, 30, 50},
	}
	suite.testFailureCases = []testCase{
		{30, 40, 60},
		{70, 80, 110},
	}
}

//AddtionOfTwoRationalNumber will take two formal parameters - should return addition value
func (suite *CalculatorTestSuite) Test_given_two_rational_number_function_AddtionOfTwoRationalNumber_should_return_addition_value_on_success_case() {
	for _, tc := range suite.testSuccessCases {
		actualResultFloat32 := AddtionOfTwoRationalNumber(tc.value1Float32, tc.value2Float32)
		if returnBooleanValue := assert.Equal(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should be equal"); !returnBooleanValue {
			require.Equal(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should be equal")
		}
	}
}

func (suite *CalculatorTestSuite) Test_given_two_rational_number_function_AddtionOfTwoRationalNumber_should_return_addition_value_on_failure_case() {
	for _, tc := range suite.testFailureCases {
		actualResultFloat32 := AddtionOfTwoRationalNumber(tc.value1Float32, tc.value2Float32)
		if returnBooleanValue := assert.NotEqual(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should not be equal"); !returnBooleanValue {
			require.NotEqual(suite.T(), tc.expectedResultFloat32, actualResultFloat32, "Expected and Calculated value should not be equal")
		}
	}
}

//TearDownTest will run after all the tests in the suite have been run
func (suite *CalculatorTestSuite) TearDownTest() {
	suite.testSuccessCases = nil
	suite.testFailureCases = nil
}
