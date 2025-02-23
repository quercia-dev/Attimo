<script lang="ts">
	import { openPath } from '@tauri-apps/plugin-opener';
	import { AppBar } from '@skeletonlabs/skeleton';
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import { showAppRail } from '../stores/appStore';

	let linkElement: HTMLAnchorElement = $state();
	async function openMailto(mailtoLink: string) {
		try {
			await openPath(mailtoLink);
		} catch (error) {
			console.error('Error opening mailto link:', error);
		}
	}
</script>

<AppBar background="bg-primary-400 dark:bg-primary-900" padding="p-2">
	{#snippet lead()}
		<button
			type="button"
			onclick={() => showAppRail.update((value) => !value)}
			aria-label="Open drawer"
			class="button hover:bg-primary-300 dark:hover:bg-primary-900"
		>
			<svg class="feather">
				<use href="/icons/feather-sprite.svg#menu" />
			</svg>
		</button>
	{/snippet}
	{#snippet trail()}
		<div class="icon-container">
			<LightSwitch />

			<a
				bind:this={linkElement}
				href="mailto:example@example.com"
				target="_blank"
				aria-label="Send email"
				class="button hover:bg-primary-300 dark:hover:bg-primary-900"
				onclick={() => openMailto(linkElement.href)}
			>
				<svg class="feather">
					<use href="/icons/feather-sprite.svg#message-square" />
				</svg>
			</a>
		</div>
	{/snippet}
</AppBar>
