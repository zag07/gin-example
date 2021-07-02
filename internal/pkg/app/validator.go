package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

func (v *ValidError) Error() string {
	return v.Message
}

type ValidErrors []*ValidError

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) ValidErrors {
	var (
		err  = c.ShouldBind(v)
		errs ValidErrors
	)

	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		vErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return errs
		}

		for key, value := range vErrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return errs
	}

	return nil
}
