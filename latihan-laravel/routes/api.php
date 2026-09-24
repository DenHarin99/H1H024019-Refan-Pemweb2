<?php

use App\Http\Controllers\Api\MahasiswaController;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api\MataKuliahController;
use App\Http\Controllers\Api\ProgramStudiController;

Route::get('/status', function () {
    return response()->json([
        'sukses' => true,
        'pesan' => 'API Pemweb II aktif',
        'waktu' => now()->toIso8601String(),
    ]);
});

Route::apiResource('mahasiswa', MahasiswaController::class);
Route::apiResource('matakuliah', MataKuliahController::class);
Route::get(
    '/program-studi/{programStudi}/mahasiswa',
    [ProgramStudiController::class, 'mahasiswa']
);