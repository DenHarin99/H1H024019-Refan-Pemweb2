<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\StoreMataKuliahRequest;
use App\Http\Requests\UpdateMataKuliahRequest;
use App\Http\Resources\MataKuliahResource;
use App\Models\MataKuliah;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Resources\Json\AnonymousResourceCollection;

class MataKuliahController extends Controller
{
    public function index(): AnonymousResourceCollection
    {
        return MataKuliahResource::collection(
            MataKuliah::orderBy('kode')->paginate(10)
        );
    }

    public function store(StoreMataKuliahRequest $request): JsonResponse
    {
        $matakuliah = MataKuliah::create($request->validated());

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dibuat',
            'data' => new MataKuliahResource($matakuliah),
        ], 201);
    }

    public function show(MataKuliah $matakuliah): MataKuliahResource
    {
        return new MataKuliahResource($matakuliah);
    }

    public function update(
        UpdateMataKuliahRequest $request,
        MataKuliah $matakuliah
    ): JsonResponse {
        $matakuliah->update($request->validated());

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil diperbarui',
            'data' => new MataKuliahResource($matakuliah),
        ]);
    }

    public function destroy(MataKuliah $matakuliah): JsonResponse
    {
        $matakuliah->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mata kuliah berhasil dihapus',
        ]);
    }
}