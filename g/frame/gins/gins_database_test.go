// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gins_test

import (
    "github.com/gogf/gf/g/frame/gins"
    "github.com/gogf/gf/g/os/gfile"
    "github.com/gogf/gf/g/test/gtest"
    "testing"
)

func Test_Database(t *testing.T) {
    config := `
# 模板引擎目录
viewpath = "/home/www/templates/"
test = "v=1"
# MySQL数据库配置
[database]
    [[database.default]]
        host     = "127.0.0.1"
        port     = "3306"
        user     = "root"
        pass     = ""
        # pass     = "12345678"
        name     = "test"
        type     = "mysql"
        role     = "master"
        charset  = "utf8"
        priority = "1"
    [[database.test]]
        host     = "127.0.0.1"
        port     = "3306"
        user     = "root"
        pass     = ""
        # pass     = "12345678"
        name     = "test"
        type     = "mysql"
        role     = "master"
        charset  = "utf8"
        priority = "1"
# Redis数据库配置
[redis]
    default = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
    path := "config.toml"
    err  := gfile.PutContents(path, config)
    gtest.Assert(err, nil)
    defer gfile.Remove(path)
    defer gins.Config().Reload()

    gtest.Case(t, func() {
        dbDefault := gins.Database()
        dbTest    := gins.Database("test")
        gtest.AssertNE(dbDefault, nil)
        gtest.AssertNE(dbTest,    nil)

        gtest.Assert(dbDefault.PingMaster(), nil)
        gtest.Assert(dbDefault.PingSlave(),  nil)
        gtest.Assert(dbTest.PingMaster(),    nil)
        gtest.Assert(dbTest.PingSlave(),     nil)
    })
}

