import { writable } from 'svelte/store';
import { browser } from '$app/environment';

// initialize with default theme
const initialTheme = browser ? localStorage.getItem('theme') || 'earth' : 'earth';

export const theme = writable<string>(initialTheme);

if (browser) {
    document.body.setAttribute('data-theme', initialTheme);
}

if (browser) {
    theme.subscribe((value) => {
        localStorage.setItem('theme', value);
        document.body.setAttribute('data-theme', value);
    });
}

export function setTheme(newTheme: string): void {
    theme.set(newTheme)
}