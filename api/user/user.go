package user

import (
	"MI/models"
	"MI/pkg/jwt"
	"MI/pkg/logger"
	"MI/pkg/sms"
	"MI/pkg/validate"
	service "MI/service/user"
	"MI/utils/common"
	"MI/utils/request"
	"MI/utils/response"
	"github.com/gin-gonic/gin"
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
	users := models.Users{Mobile: userMobile.Mobile}
	//判断手机号是否已注册
	if u, _ := models.GetUserByWhere("mobile=?", users.Mobile); u.Id > 0 {
		response.RespError(c,"该手机号已注册")
		return
	}

	//验证code,从redis中获取手机号的验证码
	//if sms.SmsCheck(userMobile.Mobile,userMobile.Code) {
	//	resp.RespError(c, "验证码已失效")
	//	return
	//}
	//随机生成盐值
	users.Salt = common.GetRandomBoth(4)
	users.Password = common.Sha1En(userMobile.Password+users.Salt)
	users.Status = 1

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
	return
}

func UserInfo(c *gin.Context){

	user, exists := c.Get("user")
	if !exists {
		response.RespError(c,"用户没有登录，请登录")
	}
	//从jwt中获取用户常用信息
	userInfo := user.(*jwt.Claims)

	if user, err := models.GetUserByWhere("id=?", userInfo.Id);err == nil {

		response.RespData(c,"",user)
	}


}