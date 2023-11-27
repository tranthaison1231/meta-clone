<script lang="ts">
	export let loading: boolean = false;
</script>

<div class="h-auto w-auto">
	{#if loading}
		<div class="flex items-center justify-center">
			<slot name="loading-indicator">Loading...</slot>
		</div>
	{:else}
		<slot />
	{/if}
</div>
