<script lang="ts">
	import { TableHandler, Datatable, ThSort, ThFilter } from '@vincjo/datatables';
	import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';
	import SimplePane from '../../components/SimplePane.svelte';
	import SearchBar from '../../components/SearchBar.svelte';

	const data = [
		{ id: 1, first_name: 'Tobie', last_name: 'Vint', email: 'tvint0@fotki.com' },
		{ id: 2, first_name: 'Zacharias', last_name: 'Cerman', email: 'zcerman1@sciencedirect.com' },
		{ id: 3, first_name: 'Gérianna', last_name: 'Bunn', email: 'gbunn2@foxnews.com' },
		{ id: 4, first_name: 'Bee', last_name: 'Saurin', email: 'bsaurin3@live.com' }
	];
	const table = new TableHandler(data, { rowsPerPage: 10 });

	let searchQuery: string = '';
	let value: number = 0;
</script>

<SimplePane>
	<div class="flex-shrink-0 flex items-center w-auto h-auto space-x-4">
		<SearchBar bind:value={searchQuery} />
		<button
			type="button"
			onclick={() => console.log('Plus clicked')}
			aria-label="create new entry"
			class="small-button bg-green-300 dark:bg-green-700"
			style="margin-left: 8px;"
		>
			<svg class="feather">
				<use href="/icons/feather-sprite.svg#plus" />
			</svg>
		</button>

		<RadioGroup
			class="thin-border"
			rounded="rounded-md"
			border="border-black"
			hover="hover:bg-surface-500 dark:hover:bg-surface-800"
		>
			<RadioItem class="button" bind:group={value} name="spreadsheet" value={0}>
				<svg class="feather">
					<use href="/icons/feather-sprite.svg#layers" />
				</svg>
			</RadioItem>
			<RadioItem class="button" bind:group={value} name="justify" value={1}>
				<svg class="feather">
					<use href="/icons/feather-sprite.svg#file-text" />
				</svg>
			</RadioItem>
		</RadioGroup>
	</div>

	<Datatable basic {table}>
		<table>
			<thead>
				<tr>
					<ThSort {table} field="first_name">First Name</ThSort>
					<ThSort {table} field="last_name">Last Name</ThSort>
					<ThSort {table} field="email">Email</ThSort>
				</tr>
				<tr>
					<ThFilter {table} field="first_name" />
					<ThFilter {table} field="last_name" />
					<ThFilter {table} field="email" />
				</tr>
			</thead>
			<tbody>
				{#each table.rows as row}
					<tr>
						<td>{row.first_name}</td>
						<td>{row.last_name}</td>
						<td>{row.email}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</Datatable>
</SimplePane>
