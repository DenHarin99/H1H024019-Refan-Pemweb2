package handler

import (
	"errors"
	"strconv"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// GET: Menampilkan semua mata kuliah
func DaftarMataKuliah(c fiber.Ctx) error {
	daftar, err := repository.AmbilSemuaMataKuliah()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}

// GET: Menampilkan detail mata kuliah
func DetailMataKuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID tidak valid",
		})
	}

	matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   matakuliah,
	})
}

// POST: Menambahkan mata kuliah
func BuatMataKuliah(c fiber.Ctx) error {
	var dto PermintaanMataKuliah

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	matakuliah := model.MataKuliah{
		ProgramStudiID: dto.ProgramStudiID,
		Kode:           dto.Kode,
		Nama:           dto.Nama,
		SKS:            dto.SKS,
		Semester:       dto.Semester,
	}

	if err := repository.SimpanMataKuliah(&matakuliah); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menyimpan mata kuliah",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil ditambahkan",
		"data":   matakuliah,
	})
}

// PUT: Memperbarui mata kuliah
func PerbaruiMataKuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID tidak valid",
		})
	}

	matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	var dto PermintaanMataKuliah
	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	matakuliah.ProgramStudiID = dto.ProgramStudiID
	matakuliah.Kode = dto.Kode
	matakuliah.Nama = dto.Nama
	matakuliah.SKS = dto.SKS
	matakuliah.Semester = dto.Semester

	if err := repository.PerbaruiMataKuliah(&matakuliah); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memperbarui mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil diperbarui",
		"data":   matakuliah,
	})
}

// DELETE: Menghapus mata kuliah
func HapusMataKuliah(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID tidak valid",
		})
	}

	matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	if err := repository.HapusMataKuliah(&matakuliah); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menghapus mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil dihapus",
	})
}
