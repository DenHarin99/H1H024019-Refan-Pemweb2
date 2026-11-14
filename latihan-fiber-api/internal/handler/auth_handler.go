package handler

import (
    "github.com/gofiber/fiber/v3"
    "golang.org/x/crypto/bcrypt"

    "latihan-fiber-api/internal/auth"
    "latihan-fiber-api/internal/config"
    "latihan-fiber-api/internal/model"
    "latihan-fiber-api/internal/validasi"
	"time"
)

func Registrasi(c fiber.Ctx) error {
    var permintaan PermintaanRegistrasi

    err := c.Bind().Body(&permintaan)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
    }

    galat := validasi.Periksa(permintaan)
    if galat != nil {
        return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
            "sukses": false,
            "pesan":  "Data yang dikirim tidak valid",
            "galat":  galat,
        })
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(permintaan.KataSandi), bcrypt.DefaultCost)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Gagal memproses kata sandi")
    }

    pengguna := model.Pengguna{
        Nama:      permintaan.Nama,
        Email:     permintaan.Email,
        KataSandi: string(hash),
        Peran:     "mahasiswa",
    }

    err = config.DB.Create(&pengguna).Error
    if err != nil {
        return fiber.NewError(fiber.StatusUnprocessableEntity, "Email sudah terdaftar")
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "sukses": true,
        "pesan":  "Registrasi berhasil",
        "data":   pengguna,
    })
}

func Login(c fiber.Ctx) error {
    var permintaan PermintaanLogin

    err := c.Bind().Body(&permintaan)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Format body tidak valid")
    }

    galat := validasi.Periksa(permintaan)
    if galat != nil {
        return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
            "sukses": false,
            "pesan":  "Data yang dikirim tidak valid",
            "galat":  galat,
        })
    }

    var pengguna model.Pengguna

    err = config.DB.Where("email = ?", permintaan.Email).First(&pengguna).Error
    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Email atau kata sandi tidak sesuai")
    }

    err = bcrypt.CompareHashAndPassword([]byte(pengguna.KataSandi), []byte(permintaan.KataSandi))
    if err != nil {
        return fiber.NewError(fiber.StatusUnauthorized, "Email atau kata sandi tidak sesuai")
    }

    konf := config.Muat()

    token, err := auth.TerbitkanToken(pengguna.ID, pengguna.Peran, konf.JwtSecret, konf.JwtTtlMenit)
    if err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Gagal menerbitkan token")
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "pesan":  "Login berhasil",
        "data": fiber.Map{
            "pengguna": pengguna,
            "token":    token,
            "tipe":     "Bearer",
        },
    })
}

func Profil(c fiber.Ctx) error {
    idPengguna := c.Locals("id_pengguna")

    var pengguna model.Pengguna

    err := config.DB.First(&pengguna, idPengguna).Error
    if err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Pengguna tidak ditemukan")
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "data":   pengguna,
    })
}


func Segarkan(c fiber.Ctx) error {
    // Mengambil ID pengguna dari hasil pemeriksaan middleware
    idPengguna, cocok := c.Locals("id_pengguna").(uint)
    if !cocok {
        return fiber.NewError(
            fiber.StatusUnauthorized,
            "Identitas pengguna tidak valid",
        )
    }

    // Mengambil peran pengguna dari hasil pemeriksaan middleware
    peran, cocok := c.Locals("peran").(string)
    if !cocok {
        return fiber.NewError(
            fiber.StatusUnauthorized,
            "Peran pengguna tidak valid",
        )
    }

    // Membaca konfigurasi JWT
    konf := config.Muat()

    // Menerbitkan token baru
    token, err := auth.TerbitkanToken(
        idPengguna,
        peran,
        konf.JwtSecret,
        konf.JwtTtlMenit,
    )
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "Gagal menerbitkan token baru",
        )
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "pesan":  "Token berhasil diperbarui",
        "data": fiber.Map{
            "token": token,
            "tipe":  "Bearer",
        },
    })
}


func Logout(c fiber.Ctx) error {
    // Mengambil token dari middleware
    token, cocok := c.Locals("token").(string)
    if !cocok || token == "" {
        return fiber.NewError(
            fiber.StatusUnauthorized,
            "Token tidak ditemukan",
        )
    }

    // Mengambil waktu kedaluwarsa token
    expiresAt, cocok := c.Locals("token_exp").(time.Time)
    if !cocok {
        return fiber.NewError(
            fiber.StatusUnauthorized,
            "Waktu kedaluwarsa token tidak ditemukan",
        )
    }

    // Menghitung hash token
    hash := auth.HashToken(token)

    // Menyimpan token ke blacklist
    daftar := model.TokenBlacklist{
        TokenHash: hash,
        ExpiresAt: expiresAt,
    }

    err := config.DB.Create(&daftar).Error
    if err != nil {
        return fiber.NewError(
            fiber.StatusInternalServerError,
            "Gagal melakukan logout",
        )
    }

    return c.JSON(fiber.Map{
        "sukses": true,
        "pesan":  "Logout berhasil",
    })
}
