package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateValidateToken(t *testing.T) {
	// Testcases list
	testCases := []struct {
		testName string
		id       string
		exp      int64
	}{
		{
			testName: "success",
			id:       "3aaaaaa",
		},
		{
			testName: "success",
			id:       "4babaaaahaha",
		}, {
			testName: "success",
			id:       "1ajdhdgdhd",
		}, {
			testName: "success",
			id:       "1ajdd",
		}, {
			testName: "success",
			id:       "hdgdhd",
		}, {
			testName: "success",
			id:       "1ajdhdgd",
		}, {
			testName: "success",
			id:       "2ajgdhd",
		}, {
			testName: "success",
			id:       "3ajgdhd",
		}, {
			testName: "success",
			id:       "4ajgdhd",
		}, {
			testName: "success",
			id:       "5ajgdhd",
		}, {
			testName: "success",
			id:       "6ajgdhd",
		},
	}

	for _, testCase := range testCases {
		newToken, err := NewAuth("coba", "test").CreateToken(&testCase.id)
		assert.Nil(t, err)
		fmt.Println(*newToken)

		tokenAdmin, err := NewAuth("coba", "test").TokenAdmin(&testCase.id)
		assert.Nil(t, err)
		fmt.Println(*tokenAdmin)

		user_id, er := NewAuth("coba", "test").ValidateToken(newToken)
		assert.Nil(t, er)

		admin_id, er := NewAuth("coba", "test").ValidateAdmin(tokenAdmin)
		assert.Nil(t, er)

		assert.Equal(t, testCase.id, *user_id)
		assert.Equal(t, testCase.id, *admin_id)

	}
}
