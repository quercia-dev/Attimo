<script>
	let searchQuery = $state('');
	let topDivHeight = $state(150);
</script>
<div
	class="bg-surface-200 dark:bg-surface-500 p-4 space-y-2 rounded-md shadow-xl flex-shrink-0"
	style="height: {topDivHeight}px; min-height: 110px;"
>
	<h1 class="text-2xl font-bold">Search</h1>
	<!-- New flex container for horizontal layout -->
	<div class="flex items-center space-x-2">
		<input
			id="search"
			type="search"
			placeholder="dammi un Attimo..."
			bind:value={searchQuery}
			oninput={() => console.log(searchQuery)}
			class="mt-1 block w-1/3 rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-gray-800"
		/>
        <button
        type="button"
        onclick={() => console.log('search dashboard')}
        aria-label="Open drawer"
        class="icon-button hover:bg-surface-400 dark:hover:bg-primary-600"
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
                class="feather feather-plus"
                >
                <line x1="12" y1="5" x2="12" y2="19"></line>
                <line x1="5" y1="12" x2="19" y2="12"></line>
            </svg>
        </button>
	</div>
</div>
