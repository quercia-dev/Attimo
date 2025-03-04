<script lang="ts">
	import Calendar from '@event-calendar/core';
	import List from '@event-calendar/list';
	import DayGrid from '@event-calendar/day-grid';
	import TimeGrid from '@event-calendar/time-grid';
	import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';

	import SimplePane from '../../components/SimplePane.svelte';
	import SearchBar from '../../components/SearchBar.svelte';

	let searchQuery = '';

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

	let value: number = 0;
</script>

<SimplePane>
	<div class="flex-shrink-0 flex items-center w-auto h-auto">
		<SearchBar bind:value={searchQuery} />

		<RadioGroup class= "ml-2 thin-border" rounded="rounded-md" border="border-black" hover="hover:bg-surface-500 dark:hover:bg-surface-800">
			<RadioItem class="button" bind:group={value} name="time of day" value={0}>
				<svg class="feather">
					<use href="/icons/feather-sprite.svg#clock" />
				</svg>
			</RadioItem>
			<RadioItem class="button" bind:group={value} name="duration" value={1}>
				<svg class="feather">
					<use href="/icons/feather-sprite.svg#move" />
				</svg>
			</RadioItem>
			<RadioItem class="button" bind:group={value} name="label" value={2}>
				<svg class="feather">
					<use href="/icons/feather-sprite.svg#tag" />
				</svg>
			</RadioItem>
		</RadioGroup>
	</div>

	<div style="height: calc(100% - 3rem); overflow: auto;">
		<Calendar {plugins} {options} />
	</div>
</SimplePane>
