<script lang="ts">
	import { groupsApi } from '$lib/apis/groups';
	import { useQuery } from '@sveltestack/svelte-query';

	const result = useQuery(['groups'], () => groupsApi.getAll());
</script>

<div>
	{#each $result.data?.groups ?? [] as { name, logo }}
		<div>
			<img src={logo} alt={name} />
		</div>
	{/each}
</div>
