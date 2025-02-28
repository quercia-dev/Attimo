import { writable } from 'svelte/store';
import type { FormField, Category } from '$lib/types/types';

export const categories = writable<Category[]>([
    { id: 'general', name: 'General' },
    { id: 'financial', name: 'Financial' },
    { id: 'personal', name: 'Personal' }
]);

export const formFields = writable<Record<string, FormField[]>>({
    general: [
        { id: 'title', label: 'Title', type: 'text', placeholder: 'Enter title', required: true },
        { id: 'description', label: 'Description', type: 'textarea', placeholder: 'Enter description' }
    ],
    financial: [
        { id: 'amount', label: 'Amount', type: 'number', placeholder: '0.00', required: true },
        { id: 'currency', label: 'Currency', type: 'select', options: ['USD', 'EUR', 'GBP'] }
    ],
    personal: [
        { id: 'name', label: 'Full Name', type: 'text', placeholder: 'John Doe', required: true },
        { id: 'document', label: 'Upload Document', type: 'file' }
    ]
});