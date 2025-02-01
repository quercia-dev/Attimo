<script lang="ts">
	import { openPath } from '@tauri-apps/plugin-opener';
	import { AppBar } from '@skeletonlabs/skeleton';
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import { showAppRail } from '../stores/appStore';

	let linkElement: HTMLAnchorElement;
	async function openMailto(mailtoLink: string) {
		try {
			await openPath(mailtoLink);
		} catch (error) {
			console.error('Error opening mailto link:', error);
		}
	}
</script>

<AppBar>
	<svelte:fragment slot="lead">
		<button
			type="button"
			on:click={() => showAppRail.update(value => !value)}
			aria-label="Open drawer"
			class="icon-button hover:bg-primary-400"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				width="32"
				height="32"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="3"
				stroke-linecap="round"
				stroke-linejoin="round"
				class="feather feather-menu"
			>
				<line x1="3" y1="12" x2="21" y2="12"></line>
				<line x1="3" y1="6" x2="21" y2="6"></line>
				<line x1="3" y1="18" x2="21" y2="18"></line>
			</svg>
		</button>
	</svelte:fragment>
	<svelte:fragment slot="trail">
  <!-- Container to hold LightSwitch and the link icon side by side -->
  <div class="icon-container">
    <!-- LightSwitch component -->
    <LightSwitch />
    
    <!-- Mail link with icon -->
    <a
      bind:this={linkElement}
      href="mailto:example@example.com"
      target="_blank"
      aria-label="Send email"
      class="icon-button hover:bg-primary-400"
      on:click={() => openMailto(linkElement.href)}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="32"
        height="32"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="3"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="feather feather-message-square"
      >
        <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
      </svg>
    </a>
  </div>
</svelte:fragment>

</AppBar>

<style>
	.icon-container {
	  display: flex;
	  align-items: center; 
	  gap: 10px;
	}
  
	.icon-button {
	  display: inline-flex;
	  align-items: center;
	}
  </style>