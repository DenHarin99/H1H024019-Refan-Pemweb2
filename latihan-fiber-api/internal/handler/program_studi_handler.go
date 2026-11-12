package handler

import (
    "github.com/gofiber/fiber/v3"

    "latihan-fiber-api/internal/repository"
)

func DaftarProgramStudi(c fiber.Ctx) error {
    daftar, err := repository.AmbilProgramStudiDenganJumlahMahasiswa()

    if err != nil {
        return c.Status(500).JSON(fiber.Map{
            "sukses": false,
            "pesan":  "Gagal mengambil data program studi",
        })
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "data":   daftar,
    })
}