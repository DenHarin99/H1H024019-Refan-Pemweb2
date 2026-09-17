<?php

namespace Database\Seeders;

use App\Models\Mahasiswa;
use App\Models\MataKuliah;
use Illuminate\Database\Seeder;
use Database\Seeders\ProgramStudiSeeder;
use Database\Seeders\MataKuliahSeeder; // <-- Import ini wajib ada

class DatabaseSeeder extends Seeder
{
    public function run(): void
    {
        $this->call([
            ProgramStudiSeeder::class,
            MataKuliahSeeder::class,
        ]);

        Mahasiswa::factory()->count(30)->create();

        $mahasiswas = Mahasiswa::all();
        $matkuls = MataKuliah::all();

        foreach ($mahasiswas as $mhs) {
            $mhs->mataKuliah()->attach(
                $matkuls->random(rand(1, 3))->pluck('id'),
                ['nilai' => fake()->randomElement(['A', 'B+', 'B', 'C+', 'C'])]
            );
        }
    }
}