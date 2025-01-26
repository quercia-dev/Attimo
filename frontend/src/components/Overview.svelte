<script>
  import {
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
  } from "flowbite-svelte";
  import openIcon from "../assets/images/plus.svg";
  import closeIcon from "../assets/images/minus.svg";
  let items = [
    {
      id: 1,
      name: "One",
      canOpen: true,
      totOpened: 10,
      canClose: false,
      tot: 40,
    },
    {
      id: 2,
      name: "Two",
      canOpen: false,
      totOpened: 15,
      canClose: true,
      tot: 45,
    },
    {
      id: 3,
      name: "Three",
      canOpen: true,
      totOpened: 30,
      canClose: true,
      tot: 55,
    },
  ];
</script>

<Table>
  <TableHead defaultRow={false}>
    <tr>
      <TableHeadCell>Table Name</TableHeadCell>
      <TableHeadCell></TableHeadCell>
      <TableHeadCell>Close</TableHeadCell>
      <TableHeadCell></TableHeadCell>
      <TableHeadCell>Entries</TableHeadCell>
    </tr>
  </TableHead>
  <TableBody tableBodyClass="divide-y">
    {#each items as { id, name, canOpen, totOpened, canClose, tot }}
      <TableBodyRow>
        <TableBodyCell align="left">{name}</TableBodyCell>
        <TableBodyCell>
          {#if canOpen}
            <a
              href="/tables"
              class="font-medium text-primary-600 hover:underline dark:text-primary-500"
            >
              <img
                src={openIcon}
                alt="Open"
                class="w-4 h-4"
                style="filter: invert(48%) sepia(13%) saturate(2878%) hue-rotate(86deg) brightness(118%) contrast(119%);"
              />
            </a>
          {/if}
        </TableBodyCell>
        <TableBodyCell>{totOpened}</TableBodyCell>
        <TableBodyCell>
          {#if canClose}
            <a
              href="/tables"
              class="font-medium text-primary-600 hover:underline dark:text-primary-500"
            >
              <img
                src={closeIcon}
                alt="Close"
                class="w-4 h-4"
                style="filter: invert(47%) sepia(89%) saturate(2629%) hue-rotate(346deg) brightness(99%) contrast(97%);"
              />
            </a>
          {/if}
        </TableBodyCell>
        <TableBodyCell>{tot}</TableBodyCell>
      </TableBodyRow>
    {/each}
  </TableBody>
</Table>
