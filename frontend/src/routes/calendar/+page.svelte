<script>
	let searchQuery = '';
	let topDivHeight = 150;

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

<div class="bg-white text-gray-900 p-4 rounded-md shadow-xl flex-1 overflow-hidden">
	<div class="calendar-container h-full w-full">
		<Calendar {plugins} {options} />
	</div>
</div>
