package api

import (
	"bbs-go/internal/controllers/render"
	// Swagger 模型导入，用于文档生成
	_ "bbs-go/internal/models"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"

	captcha2 "bbs-go/internal/pkg/captcha"
	"bbs-go/internal/pkg/errs"
	"bbs-go/internal/services"
)

// LoginController 处理用户登录相关的请求
type LoginController struct {
	Ctx iris.Context
}

// PostSignup 用户注册
// @Summary 用户注册
// @Description 用户通过邮箱、用户名和密码进行注册
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param captchaId formData string true "验证码ID"
// @Param captchaCode formData string true "验证码"
// @Param captchaProtocol formData int false "验证码协议版本" default(1)
// @Param email formData string true "邮箱"
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param rePassword formData string true "确认密码"
// @Param nickname formData string true "昵称"
// @Param redirect formData string false "登录成功后的重定向地址"
// @Success 200 {object} models.JsonResult{data=string} "注册成功"
// @Failure 400 {object} models.JsonResult{data=string} "注册失败"
// @Router /login/signup [post]
func (c *LoginController) PostSignup() *web.JsonResult {
	var (
		captchaId          = c.Ctx.PostValueTrim("captchaId")
		captchaCode        = c.Ctx.PostValueTrim("captchaCode")
		captchaProtocol, _ = params.GetInt(c.Ctx, "captchaProtocol")
		email              = c.Ctx.PostValueTrim("email")
		username           = c.Ctx.PostValueTrim("username")
		password           = c.Ctx.PostValueTrim("password")
		rePassword         = c.Ctx.PostValueTrim("rePassword")
		nickname           = c.Ctx.PostValueTrim("nickname")
		redirect           = c.Ctx.FormValue("redirect")
	)
	loginMethod := services.SysConfigService.GetLoginMethod()
	if !loginMethod.Password {
		return web.JsonErrorMsg("账号密码登录/注册已禁用")
	}
	// 根据验证码协议版本校验验证码
	if captchaProtocol == 2 {
		if !captcha2.Verify(captchaId, captchaCode) {
			return web.JsonError(errs.CaptchaError)
		}
	} else {
		if !captcha.VerifyString(captchaId, captchaCode) {
			return web.JsonError(errs.CaptchaError)
		}
	}
	user, err := services.UserService.SignUp(username, email, nickname, password, rePassword)
	if err != nil {
		return web.JsonError(err)
	}
	return render.BuildLoginSuccess(c.Ctx, user, redirect)
}

// PostSignin 用户登录
// @Summary 用户登录
// @Description 用户通过用户名和密码进行登录
// @Tags 用户认证
// @Accept json
// @Produce json
// @Param captchaId formData string true "验证码ID"
// @Param captchaCode formData string true "验证码"
// @Param captchaProtocol formData int false "验证码协议版本" default(1)
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param redirect formData string false "登录成功后的重定向地址"
// @Success 200 {object} models.JsonResult{data=string} "登录成功"
// @Failure 400 {object} models.JsonResult{data=string} "登录失败"
// @Router /login/signin [post]
func (c *LoginController) PostSignin() *web.JsonResult {
	var (
		captchaId          = c.Ctx.PostValueTrim("captchaId")
		captchaCode        = c.Ctx.PostValueTrim("captchaCode")
		captchaProtocol, _ = params.GetInt(c.Ctx, "captchaProtocol")
		username           = c.Ctx.PostValueTrim("username")
		password           = c.Ctx.PostValueTrim("password")
		redirect           = c.Ctx.FormValue("redirect")
	)
	loginMethod := services.SysConfigService.GetLoginMethod()
	if !loginMethod.Password {
		return web.JsonErrorMsg("账号密码登录/注册已禁用")
	}

	// 根据验证码协议版本校验验证码
	if captchaProtocol == 2 {
		if !captcha2.Verify(captchaId, captchaCode) {
			return web.JsonError(errs.CaptchaError)
		}
	} else {
		if !captcha.VerifyString(captchaId, captchaCode) {
			return web.JsonError(errs.CaptchaError)
		}
	}

	user, err := services.UserService.SignIn(username, password)
	if err != nil {
		return web.JsonError(err)
	}
	return render.BuildLoginSuccess(c.Ctx, user, redirect)
}

// GetSignout 退出登录
// @Summary 退出登录
// @Description 用户退出登录，清除登录状态
// @Tags 用户认证
// @Accept json
// @Produce json
// @Success 200 {object} models.JsonResult{data=string} "退出成功"
// @Failure 400 {object} models.JsonResult{data=string} "退出失败"
// @Router /login/signout [get]
func (c *LoginController) GetSignout() *web.JsonResult {
	err := services.UserTokenService.Signout(c.Ctx)
	if err != nil {
		return web.JsonError(err)
	}
	return web.JsonSuccess()
}
