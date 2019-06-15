package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

type User2 struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=30"`
	Email     string `validate:"required,email"`
}

func main() {
	user := &User2{
		FirstName: "Benson",
		LastName:  "Max",
		Age:       199,
		Email:     "particularkid@163.com",
	}

	validate := validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)
	translator, found := uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("not found")
	}

	err := validate.Struct(user) //方法返回一个结果
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			fmt.Println(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err.Translate(translator))
			}
		}
	}
}
