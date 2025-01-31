import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	extensions: ['.svelte'],
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: [vitePreprocess()],
	
	kit: {
		adapter: adapter({
			out: 'build',
			pages: 'build',
			fallback: 'index.html', 
			assets: 'build', 
		  })
	}
};
export default config;