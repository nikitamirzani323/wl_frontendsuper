package controller

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type responsebanktype struct {
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	Listcatebank interface{} `json:"listcatebank"`
	Record       interface{} `json:"record"`
}

func Banktype(c *fiber.Ctx) error {
	type payload_catebank struct {
		Master string `json:"master"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_catebank)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsebanktype{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"master":          client.Master,
		}).
		Post(PATH + "api/banktype")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsebanktype)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":       result.Status,
			"message":      result.Message,
			"record":       result.Record,
			"listcatebank": result.Listcatebank,
			"time":         time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Savebanktype(c *fiber.Ctx) error {
	type payload_savecatebank struct {
		Sdata               string `json:"sdata" `
		Page                string `json:"page" `
		Banktype_idbanktype string `json:"idbanktype"`
		Banktype_idcatebank int    `json:"idcatebank"`
		Banktype_name       string `json:"name" `
		Banktype_img        string `json:"img" `
		Banktype_status     string `json:"status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_savecatebank)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":     hostname,
			"sdata":               client.Sdata,
			"page":                client.Page,
			"banktype_id":         client.Banktype_idbanktype,
			"banktype_idcatebank": client.Banktype_idcatebank,
			"banktype_name":       client.Banktype_name,
			"banktype_img":        client.Banktype_img,
			"banktype_status":     client.Banktype_status,
		}).
		Post(PATH + "api/savebanktype")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
