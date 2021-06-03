package utils

import (
	"github.com/kataras/iris/v12/context"
	"net/http"
)

func IndexOf(array []string, search string) int {
	for index, element := range array {
		if element == search {
			return index
		}
	}
	return -1
}

func IsArrayContains(array []string, search string) bool {
	return IndexOf(array, search) != -1
}

func ShouldAcceptJson(ctx *context.Context) bool {
	if ctx.GetHeader("Accept") != "application/json" {
		ctx.StatusCode(http.StatusForbidden)
		return false
	}
	return true
}
