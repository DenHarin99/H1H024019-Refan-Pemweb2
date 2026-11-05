package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"
)

func Daftarkan(app *fiber.App) {
	api := app.Group("/api", middleware.KunciApi())

	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	mahasiswa := api.Group("/mahasiswa")

	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
	mahasiswa.Get("/nim/:nim", handler.CariMahasiswaByNIM)
}
