package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"task-week-four/utils"
)

var translator ut.Translator

func translateMessage() {
	universalTranslator := ut.New(zh.New())
	validate := binding.Validator.Engine().(*validator.Validate)
	translator, _ = universalTranslator.GetTranslator("zh")
	if err := zhTranslations.RegisterDefaultTranslations(validate, translator); err != nil {
		utils.Logger().Error(err.Error())
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("json")
	})

	// 注册消息
	translateFn := func(ut ut.Translator, fe validator.FieldError) string {
		msg, err := ut.T(fe.Tag(), fe.Field())
		if err != nil {
			utils.Logger().Warn(err.Error())
			return ""
		}
		return msg
	}
	// 遍历全部的错误消息
	for tag, text := range customMsg {
		// 注册
		// 注册函数，和翻译函数
		if err := validate.RegisterTranslation(tag, translator, func(ut ut.Translator) error {
			if err := ut.Add(tag, text, false); err != nil {
				return err
			}
			return nil
		}, translateFn); err != nil {
			utils.Logger().Warn(err.Error())
		}
	}

}

var customMsg = map[string]string{
	"roleTitleUnique": "{0}角色名称已存在",
}

func init() {
	translateMessage()
}

// 定义翻译方方法
func Translate(err error) gin.H {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}
	msg := gin.H{}
	for _, err := range errs {
		msg[err.Field()] = err.Translate(translator)
	}
	return msg
}
