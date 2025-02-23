<script lang="ts">
	let searchQuery = '';
	let topDivHeight = 150;

	import Calendar from '@event-calendar/core';
	import List from '@event-calendar/list';
	import DayGrid from '@event-calendar/day-grid';
	import TimeGrid from '@event-calendar/time-grid';
	import SimplePane from '../../components/SimplePane.svelte';

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

<SimplePane>
	<div class="calendar-container">
		<Calendar {plugins} {options} />
	</div>
</SimplePane>
