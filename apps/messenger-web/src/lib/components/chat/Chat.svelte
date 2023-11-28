<script>
	import Header from './header/Header.svelte';
	import Footer from './footer/Footer.svelte';
	import Body from './body/Body.svelte';
	import { inboxUsers, setInboxChatData } from '$lib/stores/chat';
	import { useQuery } from '@sveltestack/svelte-query';
	import { chatsApi } from '$lib/apis/chats';
	import { me } from '$lib/stores/me';

	$: currentUserIds = ($inboxUsers ?? []).map((user) => user.id);
	$: memberIds = $me.id ? currentUserIds.concat([$me.id]) : [];

	$: getChatByMemberIds = useQuery(
		['chats', { memberIds }],
		() =>
			chatsApi.getAll({
				memberIds
			}),
		{
			enabled: memberIds.length > 1,
			onSuccess(data) {
				const firstChat = data.items?.[0];
				if (firstChat) {
					setInboxChatData(firstChat);
				}
			}
		}
	);
</script>

{#if $inboxUsers.length > 0}
	<div class="relative h-screen">
		<Header />
		<Body loading={$getChatByMemberIds.isFetching} />
		<Footer />
	</div>
{:else}
	<div class="flex h-full items-center justify-center text-slate-600">
		Connect with your friends!
	</div>
{/if}
