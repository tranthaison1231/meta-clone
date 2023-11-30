<script lang="ts">
	import { chatsApi } from '$lib/apis/chats';
	import { messagesApi } from '$lib/apis/message';
	import { constructPayload } from '$lib/services/websocket';
	import { inboxUsers, inboxChat, setInboxChatData, appendNewMessage } from '$lib/stores/chat';
	import { me } from '$lib/stores/me';
	import { ws } from '$lib/stores/websocket';
	import { sendMessageSchema } from '$lib/utils/schema';
	import { useMutation } from '@sveltestack/svelte-query';
	import { Image, PlusCircle, SendIcon, Smile, StickyNote, ThumbsUp } from 'lucide-svelte';
	import { fieldProxy, superForm, superValidateSync } from 'sveltekit-superforms/client';

	const sendMessageMutate = useMutation(messagesApi.sendMessage, {
		onSuccess({ message }) {
			appendNewMessage(message);
			$content = '';
			$ws.send(constructPayload('MESSAGE', { message }));
		}
	});
	const createChatMutate = useMutation(chatsApi.createChat);

	const { form, enhance } = superForm(superValidateSync(sendMessageSchema), {
		SPA: true,
		validators: sendMessageSchema,
		taintedMessage: false,
		onUpdate: async ({ form }) => {
			if ($sendMessageMutate.isLoading) return;
			if (!$me?.id) return;

			if (form.valid) {
				const content = form.data.content ?? '👍';
				const currentChatId = $inboxChat?.id;
				if (currentChatId) {
					$sendMessageMutate.mutateAsync({
						chatId: currentChatId,
						content
					});
				} else {
					const createChatResult = await $createChatMutate.mutateAsync({
						memberIds: ($inboxUsers.map((user) => user.id) ?? []).concat([$me.id])
					});

					const chatId = createChatResult.chat.id;

					setInboxChatData(createChatResult.chat);

					$sendMessageMutate.mutateAsync({
						chatId,
						content
					});
				}
			}
		}
	});

	const content = fieldProxy(form, 'content');
</script>

<form
	method="POST"
	use:enhance
	class="absolute bottom-0 flex w-full items-center border-t px-2 py-3"
>
	<div class="flex text-blue-500">
		<button class="rounded-full p-2 hover:bg-gray-100">
			<PlusCircle />
		</button>
		<button class="rounded-full p-2 hover:bg-gray-100">
			<Image />
		</button>
		<button class="rounded-full p-2 hover:bg-gray-100">
			<StickyNote />
		</button>
		<button class="rounded-full p-2 hover:bg-gray-100">
			<Image />
		</button>
	</div>
	<div class="relative ml-4 w-full">
		<input
			class="input rounded-xl border-none bg-gray-100"
			placeholder="Aa"
			bind:value={$form.content}
		/>
		<Smile class="absolute bottom-2 right-2 text-blue-500" />
	</div>
	<button class="ml-2 rounded-full p-2 text-blue-500 hover:bg-gray-100" type="submit">
		{#if $form.content}
			<SendIcon />
		{:else}
			<ThumbsUp />
		{/if}
	</button>
</form>
