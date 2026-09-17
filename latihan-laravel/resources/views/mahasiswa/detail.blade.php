@extends('layouts.app')
@section('judul', 'Detail Mahasiswa')

@section('konten')
<div class="card mb-4">
    <div class="card-header font-weight-bold">Detail Mahasiswa</div>
    <div class="card-body">
        <p><strong>NIM:</strong> {{ $mahasiswa->nim }}</p>
        <p><strong>Nama:</strong> {{ $mahasiswa->nama }}</p>
        <p><strong>Program Studi:</strong> {{ $mahasiswa->programStudi->nama }}</p>
        <p><strong>IPK:</strong> {{ $mahasiswa->ipk }}</p>
    </div>
</div>

<h4 class="h5 mb-3">Mata Kuliah yang Diambil</h4>
<table class="table table-bordered bg-white">
    <thead>
        <tr>
            <th>Kode</th>
            <th>Nama Mata Kuliah</th>
            <th>SKS</th>
            <th>Nilai</th>
        </tr>
    </thead>
    <tbody>
        @forelse ($mahasiswa->mataKuliah as $mk)
        <tr>
            <td>{{ $mk->kode }}</td>
            <td>{{ $mk->nama }}</td>
            <td>{{ $mk->sks }}</td>
            <td>{{ $mk->pivot->nilai }}</td>
        </tr>
        @empty
        <tr>
            <td colspan="4" class="text-center">Belum ada mata kuliah yang diambil.</td>
        </tr>
        @endforelse
    </tbody>
</table>
<a href="{{ route('mahasiswa.data') }}" class="btn btn-secondary">Kembali</a>
@endsection