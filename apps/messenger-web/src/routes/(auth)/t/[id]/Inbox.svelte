<script lang="ts">
	import MoreHorizontalModal from '$lib/components/chat/MoreHorizontalModal.svelte';
	import { getUserName } from '$lib/services/user';
	import { me } from '$lib/stores/me';
	import type { Chat } from '$lib/types';
	import { BellOff } from 'lucide-svelte';
	import { twMerge } from 'tailwind-merge';

	export let isActive = false,
		chat: Chat,
		isSeen = false,
		onClick: ((id: string) => void) | undefined = undefined;

	$: members = chat.members.filter((user) => user.id !== $me?.id);
	$: lastMessage = chat.lastMessage;
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<!-- svelte-ignore a11y-click-events-have-key-events -->
<div
	class={twMerge(
		'group flex cursor-pointer items-center justify-between gap-3 p-2 hover:bg-[#F5F5F5] max-lg:bg-transparent',
		isActive && 'bg-[#F5F5F5]'
	)}
	on:click={() => onClick?.(chat.id)}
>
	<div class="flex gap-3">
		{#if members.length > 1}
			<img src={members[0]?.avatar} alt="" class="h-12 w-12 rounded-full object-cover" />
		{:else}
			<img src={members[0]?.avatar} alt="" class="h-12 w-12 rounded-full object-cover" />
		{/if}
		<div class="max-lg:hidden">
			<p class="font-bold">
				{members.length > 1
					? members.map((member) => getUserName(member)).join(', ')
					: getUserName(members[0])}
			</p>
			{#if lastMessage}
				<p class={twMerge('text-sm font-bold ', isSeen && 'font-normal text-gray-500')}>
					{lastMessage.ownerId === $me?.id
						? getUserName($me)
						: getUserName(members.find((user) => user.id === lastMessage?.ownerId))}: {lastMessage.content}
				</p>
			{/if}
		</div>
	</div>
	<div class="z-40 flex items-center gap-2 max-lg:hidden">
		<MoreHorizontalModal />
		<BellOff class="text-gray-500" size={20} />
		{#if !isSeen}
			<div class="bg-primary h-3 w-3 rounded-full" />
		{/if}
	</div>
</div>
