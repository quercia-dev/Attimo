<script lang="ts">
	import Calendar from '@event-calendar/core';
	import TimeGrid from '@event-calendar/time-grid';
	import SimplePane from '../components/SimplePane.svelte';
	import SearchBar from '../components/SearchBar.svelte';
	import QuickTable from './home/quicktable.svelte';

	let searchQuery = $state('');
	let plugins = [TimeGrid];
	let options = {
		view: 'timeGridDay',
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

	let entriesNumber = 100;
</script>

<div class="space-x-2" style="display: flex; height: 100%;">
	<SimplePane>
		<div style="display: flex; align-items: center;">
			<SearchBar bind:value={searchQuery} />
			<button class="button ml-2">
				{entriesNumber} new entries this week
			</button>
		</div>

		<div style="height: calc(100% - 3rem); overflow: auto;">
			<QuickTable />
		</div>
	</SimplePane>
	
	<div class="bg-surface-50 dark:bg-primary-700 dark:text-white p-4 space-y-2 rounded-md shadow-xl overflow-hidden">
		<div style="height: 100%; overflow: auto;">
			<Calendar {plugins} {options} />
		</div>
	</div>
</div>
