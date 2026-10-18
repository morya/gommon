package eagle_validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type IValidator interface {
	Validate(i interface{}) error
}

type wrapValidator struct {
	e *EagleValidator
}

func (w *wrapValidator) Validate(i interface{}) error {
	return w.e.ValidateStruct(i)
}

func WrapEchoValidator() IValidator {
	return &wrapValidator{e: Validator}
}

var Validator = &EagleValidator{}

type EagleValidator struct {
	once sync.Once
	*validator.Validate

	translator *ut.UniversalTranslator
	ZhTrans    ut.Translator
}

func (v *EagleValidator) Engine() *validator.Validate {
	v.lazyinit()
	return v.Validate
}

func (v *EagleValidator) lazyinit() {
	v.once.Do(func() {
		v.Validate = validator.New()
		// v.validate.SetTagName("binding")

		zh := zh.New()
		v.translator = ut.New(zh, zh)
		v.ZhTrans, _ = v.translator.GetTranslator("zh")

		if err := zh_translations.RegisterDefaultTranslations(v.Validate, v.ZhTrans); err != nil {
			panic(err)
		}
	})
}

func (v *EagleValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.Validate.Struct(obj); err != nil {

			if errs, ok := err.(validator.ValidationErrors); ok {
				sliceErrs := []string{}
				for _, e := range errs {
					s := fmt.Sprintf("%v:%v %v",
						e.StructNamespace(),
						e.Tag(),
						e.Translate(v.ZhTrans),
					)
					sliceErrs = append(sliceErrs, s)
				}
				return errors.New(strings.Join(sliceErrs, ","))

			}

			return err
		}
	}
	return nil
}
