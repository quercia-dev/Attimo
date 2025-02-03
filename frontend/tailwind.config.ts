import { join } from 'path'
import type { Config } from 'tailwindcss'
import forms from '@tailwindcss/forms';
import { skeleton } from '@skeletonlabs/tw-plugin'
import { QuercusTheme } from './src/styles/quercia-theme'
import { EarthTheme } from './src/styles/earth-theme'

export default {
	darkMode: 'selector',
	content: ['./src/**/*.{html,js,svelte,ts}', join(require.resolve('@skeletonlabs/skeleton'), '../**/*.{html,js,svelte,ts}')],
	theme: {
		extend: {},
	},
	plugins: [
		forms,
		skeleton({
			themes: {
				custom: [
					QuercusTheme, EarthTheme
				],
				preset: ['rocket', 'vintage'],
			},
		}),
	],
} satisfies Config;
		  