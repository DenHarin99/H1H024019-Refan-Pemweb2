@extends('layouts.app')

@section('judul', 'Daftar Matakuliah')

@section('konten')
    <h1 class="h3 mb-4">Daftar Matakuliah</h1>

    <x-kartu-info judul="Pencarian">
        <form method="GET" action="{{ route('matakuliah.index') }}" class="row g-2">
            <div class="col-auto">
                <input
                    type="text"
                    name="q"
                    value="{{ $kataKunci }}"
                    class="form-control"
                    placeholder="Cari kode atau nama matakuliah"
                >
            </div>
            <div class="col-auto">
                <button type="submit" class="btn btn-primary">Cari</button>
            </div>
            @if ($kataKunci !== '')
                <div class="col-auto">
                    <a href="{{ route('matakuliah.index') }}" class="btn btn-outline-secondary">Reset</a>
                </div>
            @endif
        </form>
    </x-kartu-info>

    <table class="table table-bordered bg-white">
        <thead>
            <tr>
                <th>Kode</th>
                <th>Nama</th>
                <th>SKS</th>
                <th>Aksi</th>
            </tr>
        </thead>
        <tbody>
            @forelse ($daftarMatakuliah as $mk)
                <tr>
                    <td>{{ $mk['kode'] }}</td>
                    <td>{{ $mk['nama'] }}</td>
                    <td><x-badge-sks :sks="$mk['sks']" /></td>
                    <td>
                        <a href="{{ route('matakuliah.show', $mk['kode']) }}" class="btn btn-sm btn-primary">
                            Detail
                        </a>
                    </td>
                </tr>
            @empty
                <tr>
                    <td colspan="4">Tidak ada matakuliah yang cocok dengan pencarian.</td>
                </tr>
            @endforelse
        </tbody>
    </table>
@endsection
