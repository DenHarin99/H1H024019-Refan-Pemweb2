@php
    $warna = $sks < 3 ? 'bg-success' : 'bg-danger';
@endphp

<span class="badge {{ $warna }}">
    {{ $sks }} SKS
</span>
