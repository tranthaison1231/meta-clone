<script lang="ts">
	import Loading from '$lib/components/ui/Loading/Loading.svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import Message from './Message.svelte';
	import { inboxChat } from '$lib/stores/chat';
	import { messagesApi } from '$lib/apis/message';

	export let loading: boolean = false;

	$: getMessagesResult = useQuery(
		['messages', { targetMessageId: $inboxChat?.lastMessage?.id }],
		() =>
			messagesApi.getMessages({
				targetMessageId: $inboxChat?.lastMessage.id,
				isUp: true,
				chatId: String($inboxChat?.id)
			}),
		{
			enabled: !!$inboxChat?.lastMessage?.id && !!$inboxChat.id,
			keepPreviousData: false
		}
	);
	$: messages = !$inboxChat?.id ? [] : $getMessagesResult.data?.items;
</script>

<div
	class="flex h-[calc(100vh-7rem)] w-full flex-col justify-end overflow-y-auto overflow-x-hidden px-4 py-4"
>
	{#if loading || $getMessagesResult.isFetching}
		<Loading />
	{:else if messages && messages.length > 0}
		<div class="flex flex-col gap-4">
			{#each messages as message}
				<Message {message} />
			{/each}
		</div>
	{:else}
		<div class="flex h-full items-center justify-center text-slate-600">
			Send your first message!
		</div>
	{/if}
</div>
