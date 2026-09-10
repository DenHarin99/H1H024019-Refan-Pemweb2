<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{

    protected function data(): array
    {
        return [
            ['kode' => 'TK101', 'nama' => 'Algoritma dan Pemrograman', 'sks' => 3],
            ['kode' => 'TK102', 'nama' => 'Struktur Data',             'sks' => 3],
            ['kode' => 'TK203', 'nama' => 'Basis Data',                'sks' => 3],
            ['kode' => 'TK204', 'nama' => 'Jaringan Komputer',         'sks' => 2],
            ['kode' => 'TK305', 'nama' => 'Pemrograman Web II',        'sks' => 2],
        ];
    }

    public function index(Request $request)
    {
        $kataKunci = $request->query('q', '');
        $daftarMatakuliah = $this->data();

        if ($kataKunci !== '') {
            $daftarMatakuliah = array_values(array_filter(
                $daftarMatakuliah,
                function ($mk) use ($kataKunci) {
                    return str_contains(strtolower($mk['kode']), strtolower($kataKunci))
                        || str_contains(strtolower($mk['nama']), strtolower($kataKunci));
                }
            ));
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah,
            'kataKunci' => $kataKunci,
        ]);
    }

    /**
     * Menampilkan detail satu matakuliah berdasarkan kode.
     */
    public function show(string $kode)
    {
        $matakuliah = collect($this->data())->firstWhere('kode', $kode);

        return view('matakuliah.show', [
            'kode' => $kode,
            'matakuliah' => $matakuliah,
        ]);
    }
}
