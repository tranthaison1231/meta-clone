<script>
	import Header from './header/Header.svelte';
	import Footer from './footer/Footer.svelte';
	import Body from './body/Body.svelte';
	import { inboxUsers, setInboxChatData } from '$lib/stores/chat';
	import { useQuery } from '@sveltestack/svelte-query';
	import { chatsApi } from '$lib/apis/chats';
	import { me } from '$lib/stores/me';
</script>

{#if $inboxUsers.length > 0}
	<div class="flex h-screen flex-col">
		<Header />
		<Body />
		<Footer />
	</div>
{:else}
	<div class="flex h-full items-center justify-center text-slate-600">
		Connect with your friends!
	</div>
{/if}
