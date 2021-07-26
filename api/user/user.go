package user

import (
	"MI/models"
	"MI/models/req"
	"MI/pkg/cache"
	"MI/pkg/email"
	"MI/pkg/jwt"
	"MI/pkg/logger"
	"MI/pkg/sms"
	"MI/pkg/validate"
	service "MI/service/user"
	"MI/utils/common"
	"MI/utils/request"
	"MI/utils/response"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//手机号
type Mobile struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
}
//手机号 密码 验证码
type UserMobile struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required,max=20,min=6"`
	Code   string `json:"code" binding:"required,len=6"`
}
//手机号 密码
type MobilePassword struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required,max=20,min=6"`
}
//手机号 验证码
type MobileCode struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
	Code   string `json:"code" binding:"required,len=6"`
}
var UserMobileTrans = map[string]string{"Mobile": "手机号", "Password": "密码", "Code": "验证码"}
var MobileTrans = map[string]string{"Mobile": "手机号"}
var MobileCodeTrans = map[string]string{"Mobile": "手机号","Code":"验证码"}
var EmailTrans = map[string]string{"Email":"邮箱"}
var UserTrans = map[string]string{"Uid":"用户id","NikeName":"昵称","RealName":"真实姓名"}
//手机号、密码登录
func Login(c *gin.Context){
	var userMobile MobilePassword

	if err := c.BindJSON(&userMobile);err != nil {
		msg := validate.TransTagName(&UserMobileTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	//根据手机号 查询是否存在该用户
	u, err := models.GetUserByWhere("mobile=?", userMobile.Mobile)
	if err != nil {
		response.RespError(c,"该手机号不存在")
		return
	}
	//校验密码
	if common.Sha1En(userMobile.Password+u.Salt) != u.Password {
		response.RespError(c,"密码错误")
		return
	}
	//用户密码登录逻辑  分发token
	service.DoLogin(c,u)
}
//手机验证码登录
func LoginByMobileCode(c *gin.Context){
	var userMobile MobileCode
	if err := c.BindJSON(&userMobile);err != nil {
		msg := validate.TransTagName(&UserMobileTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
//	.... 判断验证码是否符合，判断手机号是否符合要求 DoLogin分发token
}

//手机号注册
func SingUpByMobile(c *gin.Context){
	var userMobile UserMobile
	if err := c.BindJSON(&userMobile);err != nil {
		msg := validate.TransTagName(&UserMobileTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	cacheKey := fmt.Sprintf("code:%s",userMobile.Mobile)
	has, err := cache.Get(context.Background(), cacheKey)
	if err != nil {
		logger.Logger.Error()
		return
	}
	if has != userMobile.Code{
		response.RespError(c,"验证码已过期")
		return
	}
	if u, _ := models.GetUserByWhere("mobile=?", userMobile.Mobile); u.Id > 0 {
		response.RespError(c,"该手机号已注册")
		return
	}
	users := models.Users{Mobile: userMobile.Mobile}
	//随机生成盐值
	users.Salt = common.GetRandomBoth(4)
	users.Password = common.Sha1En(userMobile.Password+users.Salt)
	users.Status = 1
	users.NikeName=common.GetRandomBoth(5)

	trace := models.Trace{}
	trace.Ip=common.IpStringToInt(request.GetClientIp(c))
	trace.Type=models.TraceTypeLogin

	device := models.Device{
		Ip: common.IpStringToInt(request.GetClientIp(c)),
		Client: c.GetHeader("User-Agent"),
	}
	//手机号注册 次联保存到数据库中
	if err := users.SingIn(&trace, &device);err != nil {
		logger.Logger.Info(err)
		response.RespError(c,"注册失败")
		return
	}
	response.RespSuccess(c,"注册成功")
	return
}
//发送手机验证码
func SendSms(c *gin.Context){

	var mobile Mobile
	if err := c.BindJSON(&mobile);err != nil {
		msg := validate.TransTagName(&MobileTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	//随机生成一个六位数
	code :=common.GetRandomNum(6)
	err := sms.SmsSet(mobile.Mobile, code)
	if err != nil {
		logger.Logger.Error(err)
	}
	//msg := strings.Replace(sms.SMSTPL, "[code]", code, 1)
	//err := sms.SendSms(p.Mobile, msg)
	//if err != nil {
	//	resp.ShowError(c, "fail")
	//	return
	//}
	//resp.ShowError(c, "success")
	//这里直接使用测试代码
	cacheKey:=fmt.Sprintf("code:%s",mobile.Mobile)
	if err := cache.Set(context.Background(), cacheKey, code,5*60*time.Second);err !=nil{
		logger.Logger.Error(err)
		return
	}
	response.RespData(c,"",code)
}
func CheckCode(c *gin.Context){
	var mobileCode MobileCode
	if err := c.BindJSON(&mobileCode);err != nil {
		msg := validate.TransTagName(&MobileTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	cacheKey := fmt.Sprintf("code:%s",mobileCode.Mobile)
	has, err := cache.Get(context.Background(), cacheKey)
	if err != nil {
		logger.Logger.Error()
		return
	}
	if has != mobileCode.Code{
		response.RespError(c,"验证码错误或已过期")
		return
	}
	//判断手机号是否已注册
	if u, _ := models.GetUserByWhere("mobile=?", mobileCode.Mobile); u.Id > 0 {
		response.RespData(c,"该手机号已注册",map[string]bool{
			"is_register":true,
		})
		return
	}
	response.RespData(c,"该手机号未注册",map[string]bool{
		"is_register":false,
	})

}

func UserInfo(c *gin.Context){

	user, exists := c.Get("user")
	if !exists {
		response.RespError(c,"用户未登录，请登录")
		return
	}
	//从jwt中获取用户常用信息
	userInfo := user.(*jwt.Claims)

	if user, err := models.GetUserByWhere("id=?", userInfo.Id);err == nil {

		response.RespData(c,"",user)
	}

}
func Logout(c *gin.Context){

	user, exists := c.Get("user")
	if !exists {
		response.RespError(c,"用户未登录，请登录")
		return
	}
	//从jwt中获取用户常用信息
	userInfo := user.(*jwt.Claims)
	//将该用户的token加入黑名单 实现退出
	if accessToken, has := request.GetParam(c, "Authorization");has{
		jwt.AddBlack(string(userInfo.Id),accessToken)
		response.RespSuccess(c,"")
	}
	return
}
func ForgetPassword(c *gin.Context){
	var userMobile UserMobile
	if err := c.BindJSON(&userMobile);err != nil {
		msg := validate.TransTagName(&UserMobileTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	cacheKey := fmt.Sprintf("code:%s",userMobile.Mobile)
	has, err := cache.Get(context.Background(), cacheKey)
	if err != nil {
		logger.Logger.Error()
		return
	}
	if has != userMobile.Code{
		response.RespError(c,"验证码错误或已过期")
		return
	}
	//判断手机号是否已注册
	user, err:= models.GetUserByWhere("mobile=?", userMobile.Mobile)
	if err != nil {

		response.RespError(c,"该手机号未注册")
		return
	}

	user.Mobile=userMobile.Mobile
	user.Salt = common.GetRandomBoth(4)
	user.Password = common.Sha1En(userMobile.Password+user.Salt)

	trace := models.Trace{}
	trace.Ip=common.IpStringToInt(request.GetClientIp(c))
	trace.Type=models.TraceTypeEdit

	device := models.Device{
		Ip: common.IpStringToInt(request.GetClientIp(c)),
		Client: c.GetHeader("User-Agent"),

	}
	if err := user.Update(&trace, &device);err != nil{
		response.RespError(c,"更新失败")
		return
	}
	response.RespSuccess(c,"")
}
func BindEmail(c *gin.Context){
	var emailReq req.EmailReq
	if err := c.BindJSON(&emailReq);err != nil {
		msg := validate.TransTagName(&EmailTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	token, err := jwt.GenerateEmailToken(emailReq)
	if err != nil {
		return
	}
	//获取 邮件模板
	notice, err := models.GetNoticeByWhere("operation_type=?", emailReq.OperationType)
	if err != nil {
		logger.Logger.Info("查询notice：",err)
		return
	}
	emailStr := fmt.Sprintf("http://localhost:8081/#/validate/email/%s",
		token)
	emailContent := strings.Replace(notice.Ext,"validate",emailStr,-1)
	if err := email.SendEmail(emailContent, emailReq.Email);err != nil {
		logger.Logger.Error("send email err:",err)
		response.RespError(c,"邮件发送失败")
		return
	}
	response.RespSuccess(c,"邮件发送成功")
}

func ValidateEmail(c *gin.Context){

	token := c.Query("token")

	if token == "" {
		response.RespError(c,"token不存在")
		return
	}
	//解析token 验证token
	emailClaims, err := jwt.ParseEmailToken(token)
	if err != nil {
		response.RespError(c,"token解析错误")
		return
	}
	if time.Now().Unix() >emailClaims.ExpiresAt {
		response.RespError(c,"token已过期")
		return
	}
	//绑定邮箱信息
	if emailClaims.OperationType == 1 {
		if err := models.BindEmail(emailClaims.UserID, emailClaims.Email);err != nil{
			response.RespError(c,"绑定失败")
			return
		}
	}
	if emailClaims.OperationType ==2 {
		if err := models.BindEmail(emailClaims.UserID, "");err != nil{
			response.RespError(c,"解绑失败")
			return
		}
	}
	response.RespSuccess(c,"")
}

func UpdateUserInfo(c *gin.Context){
	var userReq req.UserReq
	if err := c.BindJSON(&userReq);err != nil {
		msg := validate.TransTagName(&UserTrans,err)
		response.RespValidatorError(c,msg)
		return
	}
	if err := models.UpdateUser(userReq);err != nil{
		logger.Logger.Error("update user info err:",err)
		response.RespError(c,"更新失败")
		return
	}
	response.RespSuccess(c,"更新成功")
}