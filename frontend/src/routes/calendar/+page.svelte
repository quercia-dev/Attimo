<script>
	let searchQuery = '';
	let topDivHeight = 150;

	import Calendar from '@event-calendar/core';
	import List from '@event-calendar/list';
	import DayGrid from '@event-calendar/day-grid';
	import TimeGrid from '@event-calendar/time-grid';

	let plugins = [TimeGrid, List, DayGrid];
	let options = {
		initialView: 'list',
		firstDay: 1,
		slotWidth: 10,
		events: [
			{
				title: 'Event 1',
				start: '2025-02-03T10:00:00',
				end: '2025-02-03T12:00:00'
			}
		],
		headerToolbar: {
			start: 'prev, next, today',
			center: 'title',
			end: 'dayGridMonth, timeGridWeek, timeGridDay, listMonth'
		}
	};
</script>

<div class="bg-white text-gray-900 p-4 rounded-md shadow-xl flex-1 overflow-auto">
	<div class="calendar-container h-full w-full">
		<Calendar {plugins} {options} />
	</div>
</div>
