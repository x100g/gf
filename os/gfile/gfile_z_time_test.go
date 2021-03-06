// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gfile_test

import (
	"os"
	"testing"
	"time"

	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/test/gtest"
)

func Test_MTime(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {

		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.Assert(err, nil)

		t.Assert(gfile.MTime(testpath()+file1), fileobj.ModTime().Unix())
		t.Assert(gfile.MTime(""), 0)
	})
}

func Test_MTimeMillisecond(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		var (
			file1   = "/testfile_t1.txt"
			err     error
			fileobj os.FileInfo
		)

		createTestFile(file1, "")
		defer delTestFiles(file1)
		fileobj, err = os.Stat(testpath() + file1)
		t.Assert(err, nil)

		time.Sleep(time.Millisecond * 100)
		t.AssertGE(
			gfile.MTimeMillisecond(testpath()+file1),
			fileobj.ModTime().UnixNano()/1000000,
		)
		t.Assert(gfile.MTimeMillisecond(""), 0)
	})
}
