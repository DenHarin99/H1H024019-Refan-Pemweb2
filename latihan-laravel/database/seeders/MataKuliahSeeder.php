<?php

namespace Database\Seeders;

use App\Models\MataKuliah;
use Illuminate\Database\Seeder;

class MataKuliahSeeder extends Seeder
{
    public function run(): void
    {
        $matkul = [
            ['kode' => 'MK01', 'nama' => 'Pemrograman Web 2', 'sks' => 3, 'semester' => 4],
            ['kode' => 'MK02', 'nama' => 'Basis Data', 'sks' => 3, 'semester' => 2],
            ['kode' => 'MK03', 'nama' => 'Jaringan Komputer', 'sks' => 3, 'semester' => 3],
        ];

        foreach ($matkul as $item) {
            MataKuliah::create($item);
        }
    }
}