package ecode

import (
	"Week04/pkg"
)

var ParameterError = pkg.NewErrCode("参数错误", 400)

var AirticleNotFound = pkg.NewErrCode("文章不存在", 100)
var AirticleExits = pkg.NewErrCode("文章已存在", 101)
var AirticleUpdateFail = pkg.NewErrCode("文章更新失败", 102)
var AirticleListFail = pkg.NewErrCode("文章列表获取失败", 103)
var AirtilceIdIsError = pkg.NewErrCode("文章ID有误", 104)
