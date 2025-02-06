<script>
	let searchQuery = $state('');
	let topDivHeight = $state(150);

	import Calendar from '@event-calendar/core';
	import TimeGrid from '@event-calendar/time-grid';
	import SimplePane from '../../components/SimplePane.svelte';

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
</script>

<main
	class="flex flex-col h-[calc(100vh-8rem)] overflow-hidden"
>
	<div
		class="bg-surface-50 dark:bg-primary-500 text-color-black p-4 space-y-2 rounded-md shadow-xl flex-shrink-0"
		style="height: {topDivHeight}px; min-height: 110px;"
	>
		<input
			id="search"
			type="search"
			placeholder="Dammi un Attimo..."
			bind:value={searchQuery}
			oninput={() => console.log(searchQuery)}
			class="bg-surface-100 dark:bg-surface-700 text-surface-900 dark:text-surface-50 placeholder:text-surface-500 mt-1 block w-4/5 rounded-md shadow-sm"
		/>
	</div>

	<div
		class="slider flex-shrink-0"
		role="slider"
		tabindex="0"
		aria-valuenow={topDivHeight}
		onmousedown={startResizing}
	></div>

	<SimplePane>
		<div class="calendar-container h-full w-full">
			<Calendar {plugins} {options} />
		</div>
	</SimplePane>
</main>
