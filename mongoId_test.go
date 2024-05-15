/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-15 10:17:28
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-05-15 10:17:44
 * @FilePath     : /v2/go-common-v2-dh-validator/mongoId_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhvalidator

import "testing"

func Test_validateMongoId(t *testing.T) {
	testCases := []struct {
		description string
		mongoId     string
		expected    bool
	}{
		// 有效ObjectId的测试用例
		{
			description: "Valid ObjectId",
			mongoId:     "5e3f7c2f9b0b6527f34f0f4e",
			expected:    true,
		},
		// 无效ObjectId的测试用例
		{
			description: "Invalid ObjectId - wrong length",
			mongoId:     "5e3f7c2f9b0b6527f34f0f4e1",
			expected:    false,
		},
		{
			description: "Invalid ObjectId - contains characters other than hex",
			mongoId:     "5e3f7c2f9b0b6527f34f0f4eg",
			expected:    false,
		},
		// 空字符串的测试用例
		{
			description: "Empty string",
			mongoId:     "",
			expected:    false,
		},
		// 包含空格的ObjectId字符串的测试用例
		{
			description: "ObjectId with space",
			mongoId:     "5e3f7c2f 9b0b6527 f34f0f4e",
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := validateMongoId(tc.mongoId)
			if result != tc.expected {
				t.Errorf("Expected validateMongoId(%s) to be %v, but got %v", tc.mongoId, tc.expected, result)
			}
		})
	}
}
