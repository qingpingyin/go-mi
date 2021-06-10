package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"regexp"
	"strings"
)
//初始化一个翻译器
var Trans ut.Translator

func InitValidate()  {
	uni := ut.New(zh.New())
	//直接使用中文翻译器
	Trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterValidation("mobile",checkMobile)
		//给自定义字段添加中文翻译
		v.RegisterTranslation("mobile",Trans,registerTranslator("mobile","{0}格式错误"),translate)
		//注册翻译器
		_= zh_translations.RegisterDefaultTranslations(v, Trans)
	}
}
//map 自定义Tag
func TransTagName( langs *map[string]string,err error) interface{} {
	for _, e := range err.(validator.ValidationErrors) {
		v:=e.Translate(Trans)//翻译错误
		for key,value:=range *langs  {
			v=strings.Replace(v,key,value,-1)
		}
		return v
	}
	return err
}
// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func checkMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)
}