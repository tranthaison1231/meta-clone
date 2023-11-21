<script lang="ts">
	import { communitiesApi } from '$lib/apis/communities';
	import { useQuery } from '@sveltestack/svelte-query';

	const result = useQuery(['communities'], () => communitiesApi.getAll());
</script>

<div>
	{#each $result.data?.communities ?? [] as { name, logo }}
		<div>
			<img src={logo} alt={name} />
		</div>
	{/each}
</div>
