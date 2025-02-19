<script lang="ts">
	import { TableHandler, Datatable, ThSort, ThFilter } from '@vincjo/datatables';

	const data = [
		{ id: 1, first_name: 'Tobie', last_name: 'Vint', email: 'tvint0@fotki.com' },
		{ id: 2, first_name: 'Zacharias', last_name: 'Cerman', email: 'zcerman1@sciencedirect.com' },
		{ id: 3, first_name: 'Gérianna', last_name: 'Bunn', email: 'gbunn2@foxnews.com' },
		{ id: 4, first_name: 'Bee', last_name: 'Saurin', email: 'bsaurin3@live.com' }
	];
	const table = new TableHandler(data, { rowsPerPage: 10 });
</script>

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
