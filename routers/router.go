package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/wl_frontendsuper/controller"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Static("/", "frontend/dist", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	app.Get("api/healthz", controller.HealthCheck)
	app.Post("/api/init", controller.Init)
	app.Post("/api/log", controller.Log)
	app.Post("/api/login", controller.CheckLogin)
	app.Post("/api/company", controller.Company)
	app.Post("/api/savecompany", controller.Companysave)
	app.Post("/api/companylistadmin", controller.Companylistadmin)
	app.Post("/api/savecompanylistadmin", controller.Companysavelistadmin)
	app.Post("/api/companylistagen", controller.Companylistagen)

	app.Post("/api/setting", controller.Setting)
	app.Post("/api/savesetting", controller.Settingsave)

	app.Post("/api/admin", controller.Admin)
	app.Post("/api/saveadmin", controller.Saveadmin)
	app.Post("/api/adminrule", controller.Adminrule)
	app.Post("/api/saveadminrule", controller.Saveadminrule)
	app.Post("/api/currency", controller.Currency)
	app.Post("/api/savecurrency", controller.Savecurr)
	app.Post("/api/categorygame", controller.Categame)
	app.Post("/api/savecategorygame", controller.Savecategame)
	app.Post("/api/categorybank", controller.Catebank)
	app.Post("/api/savecategorybank", controller.Savecatebank)
	app.Post("/api/banktype", controller.Banktype)
	app.Post("/api/savebanktype", controller.Savebanktype)
	app.Post("/api/providergame", controller.Providergame)
	app.Post("/api/saveprovidergame", controller.Saveprovidergame)
	app.Post("/api/game", controller.Game)
	app.Post("/api/savegame", controller.Savegame)
	return app
}
