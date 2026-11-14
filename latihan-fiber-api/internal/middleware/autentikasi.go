package middleware

import (
    "errors"
    "strconv"
    "strings"
    "time"

    "github.com/gofiber/fiber/v3"
    "gorm.io/gorm"

    "latihan-fiber-api/internal/auth"
    "latihan-fiber-api/internal/config"
    "latihan-fiber-api/internal/model"
)

func Autentikasi() fiber.Handler {
    konf := config.Muat()

    return func(c fiber.Ctx) error {
        header := c.Get("Authorization")

        if header == "" {
            return fiber.NewError(fiber.StatusUnauthorized, "Header Authorization tidak ditemukan")
        }

        bagian := strings.Split(header, " ")

        if len(bagian) != 2 || bagian[0] != "Bearer" {
            return fiber.NewError(fiber.StatusUnauthorized, "Format header Authorization tidak sesuai")
        }

        klaim, err := auth.PeriksaToken(bagian[1], konf.JwtSecret)
        if err != nil {
            return fiber.NewError(fiber.StatusUnauthorized, "Token tidak valid atau sudah kedaluwarsa")
        }

		// Memeriksa apakah token sudah masuk blacklist
var jumlah int64

	err = config.DB.Model(&model.TokenBlacklist{}).
		Where(
			"token_hash = ? AND expires_at > ?",
			auth.HashToken(bagian[1]),
			time.Now(),
		).
		Count(&jumlah).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"Gagal memeriksa status token",
			)
		}
	}

	if jumlah > 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Token sudah dicabut",
		)
	}
        id, err := strconv.ParseUint(klaim.Subject, 10, 64)
        if err != nil {
            return fiber.NewError(fiber.StatusUnauthorized, "Isi token tidak sesuai")
        }

        
	c.Locals("id_pengguna", uint(id))
	c.Locals("peran", klaim.Peran)

	// Simpan token dan waktu kedaluwarsa untuk proses logout
	c.Locals("token", bagian[1])
	c.Locals("token_exp", klaim.ExpiresAt.Time)

	return c.Next()

    }
}

func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
    return func(c fiber.Ctx) error {
        peran, cocok := c.Locals("peran").(string)

        if cocok == false {
            return fiber.NewError(fiber.StatusForbidden, "Peran pengguna tidak diketahui")
        }

        for i := 0; i < len(peranDiizinkan); i++ {
            if peranDiizinkan[i] == peran {
                return c.Next()
            }
        }

        return fiber.NewError(fiber.StatusForbidden, "Anda tidak berwenang mengakses sumber daya ini")
    }
}
