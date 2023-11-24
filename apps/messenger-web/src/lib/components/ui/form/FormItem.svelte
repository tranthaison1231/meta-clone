<script lang="ts">
	export let errors: string[] | undefined = [],
		label: string | undefined = undefined,
		layout: 'vertical' | 'horizontal' = 'vertical',
		required: boolean | undefined = false;

	const labelClass = layout === 'vertical' ? 'flex flex-col gap-2' : 'flex gap-4 items-center my-1';
</script>

<div class="flex w-full flex-col gap-2">
	<div class={labelClass}>
		<span class="flex items-center gap-1">
			{label}
			{#if required}
				<span class="text-red-500"> * </span>
			{/if}
		</span>
		<slot />
	</div>
	{#if errors && errors.length > 0}
		{#each errors as error}
			<span class="mt-1 w-full text-red-500">{error}</span>
		{/each}
	{/if}
</div>
