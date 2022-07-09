package mFiber

import (
	"github.com/EasyGolang/goTools/mUrl"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

/*
	var data struct {
		FullPath string
	}
	json := mFiber.DataParser(c, &data)

*/

func Parser(c *fiber.Ctx, con ...any) map[string]any {
	json := make(map[string]any)

	// 1 解析 链接参数
	fullPath := c.BaseURL() + c.OriginalURL()
	urlData := mUrl.InitUrl(fullPath).ParseQuery()
	if len(urlData) > 0 {
		for key, val := range urlData {
			if len(val) > 0 {
				json[key] = val[0]
			}
		}
	}

	// 2 解析 fromData
	fromData, _ := c.MultipartForm()
	if fromData != nil {
		data := fromData.Value
		for key, val := range data {
			json[key] = val[0]
		}
	}

	// 3 解析 jsonData
	c.BodyParser(&json)

	if len(con) > 0 {
		mapstructure.Decode(json, con[0])
	}

	return json
}
