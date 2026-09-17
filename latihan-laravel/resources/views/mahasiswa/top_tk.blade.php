@extends('layouts.app')
@section('judul', 'Top 10 Mahasiswa Teknik Komputer')

@section('konten')
<h1 class="h3 mb-4">Top 10 Mahasiswa IPK Tertinggi - Teknik Komputer</h1>
<table class="table table-striped bg-white">
    <thead>
        <tr>
            <th>Peringkat</th>
            <th>NIM</th>
            <th>Nama</th>
            <th>IPK</th>
        </tr>
    </thead>
    <tbody>
        @foreach ($topMahasiswa as $index => $mhs)
        <tr>
            <td>{{ $index + 1 }}</td>
            <td>{{ $mhs->nim }}</td>
            <td>{{ $mhs->nama }}</td>
            <td>{{ $mhs->ipk }}</td>
        </tr>
        @endforeach
    </tbody>
</table>
@endsection