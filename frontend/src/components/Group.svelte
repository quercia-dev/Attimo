<script lang="ts">
    import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';
</script>

<RadioGroup
	class="thin-border"
	rounded="rounded-md"
	border="border-black"
	hover="hover:bg-surface-500 dark:hover:bg-surface-800"
>
	<slot></slot>
</RadioGroup>