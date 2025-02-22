<script>
	let searchQuery = $state('');
	let topDivHeight = $state(400);

	import Calendar from '@event-calendar/core';
	import TimeGrid from '@event-calendar/time-grid';
	import SimplePane from '../../components/SimplePane.svelte';
	import SearchBar from '../../components/SearchBar.svelte';
	import QuickTable from './quicktable.svelte'

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

	function startResizing(event) {
		const startY = event.clientY;
		const startHeight = topDivHeight;

		function onMouseMove(e) {
			const newHeight = startHeight + (e.clientY - startY);
			const maxHeight = window.innerHeight - 200; // Leave space for calendar
			topDivHeight = Math.max(50, Math.min(newHeight, maxHeight));
		}

		function onMouseUp() {
			window.removeEventListener('mousemove', onMouseMove);
			window.removeEventListener('mouseup', onMouseUp);
		}

		window.addEventListener('mousemove', onMouseMove);
		window.addEventListener('mouseup', onMouseUp);
	}

	let entriesNumber = 100
</script>

<main class="no-select flex flex-col overflow-hidden">
	<div
		class="bg-surface-50 dark:bg-primary-700 text-color-black p-4 space-y-2 rounded-md shadow-xl flex-shrink-0"
		style="height: {topDivHeight}px; min-height: 110px;"
	>
		<div style="display: flex; align-items: center;">
			<SearchBar bind:value={searchQuery} />
			<button type="button" 
					class="icon-button bg-primary-300 dark:bg-surface-900 ml-4"> 
				Added {entriesNumber} new entries this week</button>
		</div>
			
		<div style="height: calc(100% - 3rem); overflow: auto;">
			<QuickTable/>
		</div>

	</div>

	<div
		class="slider flex-shrink-0"
		role="slider"
		tabindex="0"
		aria-valuenow={topDivHeight}
		onmousedown={startResizing}
	></div>

	<SimplePane>
		<div class="calendar-container">
			<Calendar {plugins} {options} />
		</div>
	</SimplePane>
</main>
