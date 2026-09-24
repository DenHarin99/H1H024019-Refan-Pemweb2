<?php

namespace App\Http\Requests;

use Illuminate\Contracts\Validation\ValidationRule;
use Illuminate\Foundation\Http\FormRequest;

class UpdateMataKuliahRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        return [
            'kode' => ['sometimes', 'string', 'max:10', 'unique:matakuliahs,kode,' . $this->route('matakuliah')],
            'nama' => ['sometimes', 'string', 'max:100'],
            'sks' => ['sometimes', 'integer', 'min:1', 'max:6'],
            'semester' => ['sometimes', 'integer', 'min:1', 'max:8'],
        ];
    }
}
