package tests

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCallFCAPI(t *testing.T) {
	testTools := SetupSuite(t, SetupOptions{})
	defer testTools.Teardown(t)

	type input struct {
		pass bool
	}

	type expected struct {
		Expected
	}
	suite := Testcase{
		"001.Call EOF(LP-620)": {
			Mock: &Mock{},
			Input: input{
				pass: false,
			},
			Expected: expected{
				Expected: Expected{
					Error: io.EOF,
				},
			},
		},
		"002.Call Success(LP-620)": {
			Mock: &Mock{},
			Input: input{
				pass: true,
			},
			Expected: expected{
				Expected: Expected{
					Error: nil,
				},
			},
		},
	}

	for name, tc := range suite {
		t.Run(name, func(t *testing.T) {
			testTools := SetupTest(t, SetupOptions{Mock: tc.Mock})
			defer testTools.Teardown(t)
			input := tc.Input.(input).pass
			err := testTools.FC.CallFcApiEOF(input)
			expected := tc.Expected.(expected)
			if expected.Error != nil {
				require.Error(t, err)
				assert.EqualError(t, err, expected.Error.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
