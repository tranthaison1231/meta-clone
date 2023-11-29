<script lang="ts">
	import type { Message } from '$lib/types/message';
	import { cn } from '$lib/utils/cn';
	import { me } from '$lib/stores/me';
	import ReactMessageModal from '../ReactMessageModal.svelte';

	export let message: Message;
</script>

<div class={cn('group', $me?.id === message.ownerId ? 'message-reverse' : 'message')}>
	<p class={cn($me?.id === message.ownerId ? 'message-box-reverse' : 'message-box')}>
		{message.content}
	</p>
	<div
		class="pointer-events-none flex items-center opacity-0 transition-all group-hover:pointer-events-auto group-hover:opacity-100"
	>
		<ReactMessageModal />
	</div>
</div>

<style>
	.message {
		@apply flex w-full gap-2;
	}
	.message-reverse {
		@apply flex w-full flex-row-reverse gap-2;
	}
	.message-box {
		@apply h-auto max-w-[350px] break-all rounded-3xl rounded-bl-none bg-[#F5F5F5] px-6 py-3 shadow-sm;
	}
	.message-box-reverse {
		@apply h-auto max-w-[350px] break-all rounded-3xl rounded-br-none bg-[#F5F5F5] px-6 py-4;
	}
</style>
