<script>
	let searchQuery = '';
	let topDivHeight = 150;

	import Calendar from '@event-calendar/core';
	import TimeGrid from '@event-calendar/time-grid';
	import List from '@event-calendar/list';
	import DayGrid from '@event-calendar/day-grid';

	let plugins = [TimeGrid, List, DayGrid];
	let options = {
		headerToolbar: {
			left: 'prev,next today',
			center: 'title',
			right: 'timeGridWeek,dayGridMonth,listWeek'
		},
		initialView: 'timeGridWeek',
		firstDay: 1,
		slotWidth: 10,
		views: {
			timeGridWeek: {
				type: 'timeGrid',
				duration: { weeks: 1 }
			},
			dayGridMonth: {
				type: 'dayGrid',
				duration: { months: 1 }
			},
			listWeek: {
				type: 'list',
				duration: { weeks: 1 }
			}
		},
		events: [
			{
				title: 'Event 1',
				start: '2025-02-03T10:00:00',
				end: '2025-02-03T12:00:00'
			}
		]
	};
</script>

<div class="bg-white text-gray-900 p-4 rounded-md shadow-xl flex-1 overflow-auto">
	<div class="calendar-container h-full w-full">
		<Calendar {plugins} {options} />
	</div>
</div>
