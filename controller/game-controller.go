package controller

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type responsegame struct {
	Status           int         `json:"status"`
	Message          string      `json:"message"`
	Listcategorygame interface{} `json:"listcategorygame"`
	Listprovidergame interface{} `json:"listprovidergame"`
	Record           interface{} `json:"record"`
}

func Game(c *fiber.Ctx) error {
	type payload_game struct {
		Master string `json:"master"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_game)
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
		SetResult(responsegame{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"master":          client.Master,
		}).
		Post(PATH + "api/allgame")
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
	result := resp.Result().(*responsegame)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":           result.Status,
			"message":          result.Message,
			"record":           result.Record,
			"listcategorygame": result.Listcategorygame,
			"listprovidergame": result.Listprovidergame,
			"time":             time.Since(render_page).String(),
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
func Savegame(c *fiber.Ctx) error {
	type payload_savegame struct {
		Sdata               string `json:"sdata" `
		Page                string `json:"page" `
		Game_id             int    `json:"idgame"`
		Game_idcategame     string `json:"idcategame"`
		Game_idprovidergame string `json:"idprovidergame"`
		Game_name           string `json:"name" `
		Game_imgcover       string `json:"imgcover" `
		Game_imgthumb       string `json:"imgthumb" `
		Game_endpointurl    string `json:"endpointurl" `
		Game_status         string `json:"status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	log.Println("Hostname: ", hostname)
	client := new(payload_savegame)
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
			"game_id":             client.Game_id,
			"game_idcategame":     client.Game_idcategame,
			"game_idprovidergame": client.Game_idprovidergame,
			"game_name":           client.Game_name,
			"game_imgcover":       client.Game_imgcover,
			"game_imgthumb":       client.Game_imgthumb,
			"game_endpointurl":    client.Game_endpointurl,
			"game_status":         client.Game_status,
		}).
		Post(PATH + "api/savegame")
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
