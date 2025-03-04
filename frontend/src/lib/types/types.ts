export type FieldType = 'text' | 'select' | 'textarea' | 'file' | 'range' | 'number' | 'date' | 'email' | 'tel';

export interface FormField {
    id: string;
    label: string;
    type: FieldType;
    placeholder?: string;
    options?: string[];
    required?: boolean;
    validation?: {
        pattern?: string;
        minLength?: number;
        maxLength?: number;
        min?: number;
        max?: number;
    };
    defaultValue?: any;
}

// Category Types
export interface Category {
    id: string;
    name: string;
    description?: string;
    icon?: string;
}

// Form Values Type
export type FormValues = Record<string, any>;

// Form Validation Types
export interface ValidationError {
    fieldId: string;
    message: string;
}

export interface FormState {
    values: FormValues;
    errors: ValidationError[];
    isDirty: boolean;
    isSubmitting: boolean;
    isValid: boolean;
}

// Search/Filter Types
export interface SearchOptions {
    query: string;
    category?: string;
    includeDisabled?: boolean;
}