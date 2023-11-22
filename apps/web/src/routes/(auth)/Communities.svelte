<script lang="ts">
	import { communitiesApi } from '$lib/apis/communities';
	import Tooltip from '$lib/components/ui/tooltip/tooltip.svelte';
	import { useQuery } from '@sveltestack/svelte-query';

	const result = useQuery(['communities'], () => communitiesApi.getAll());
</script>

<div class="space-y-4">
	{#each $result.data?.data?.communities ?? [] as { name, logo }}
		<Tooltip content={name}>
			<button class="rounded-xl p-1 hover:bg-gray-100">
				<img src={logo} alt={name} class="rounded-xl" />
			</button>
		</Tooltip>
	{/each}
</div>
