<script>
	let searchQuery = '';

	import Calendar from '@event-calendar/core';
	import TimeGrid from '@event-calendar/time-grid';

	let plugins = [TimeGrid];
	let options = {
		view: 'timeGridWeek',
		slotMinTime: '08:00:00',
		slotMaxTime: '20:00:00',
		slotWidth: 10,
		events: [
			{
				title: 'Event 1',
				start: '2025-01-29T10:00:00',
				end: '2025-01-29T12:00:00'
			}
		]
	};
</script>

<main class="bg-surface-50 p-4 space-y-4 flex flex-col flex-1">
    <div class="bg-surface-500 p-4 space-y-2 rounded-md shadow-xl flex-grow">
        <h1 class="text-2xl font-bold">Search</h1>
        <input
            id="search"
            type="search"
            placeholder="dammi un attimo..."
            bind:value={searchQuery}
            on:input={() => console.log(searchQuery)}
            class="mt-1 block w-4/5 rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-gray-800"
        />
    </div>
    <div class="bg-white text-gray-900 p-4 rounded-md shadow-xl overflow-auto mt-auto">
        <div class="calendar-container">
            <Calendar {plugins} {options} />
        </div>
    </div>
</main>

<style>
    .calendar-container {
        width: 100%; 
        max-height: 400px;
        aspect-ratio: 16 / 9;
        overflow: auto;
    }

	.page-pane {
		min-width: 400px;
		min-height: 400px;
		padding: 1rem;
		flex: 1;
		border-radius: 10px;
	}
</style>
