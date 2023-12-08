<script lang="ts">
	import Header from './header/Header.svelte';
	import Footer from './footer/Footer.svelte';
	import Body from './body/Body.svelte';
	import { inboxUsers } from '$lib/stores/chat';
</script>

{#if $inboxUsers.length > 0}
	<div class="relative h-screen">
		<Header />
		<Body />
		<Footer />
	</div>
{:else}
	<div class="flex h-full items-center justify-center text-slate-600">
		Connect with your friends!
	</div>
{/if}
