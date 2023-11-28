<script lang="ts">
	import DefaultAvatar from '$lib/assets/images/default-avatar.jpeg';
	import { getUserName } from '$lib/services/user';
	import { setInboxUser } from '$lib/stores/chat';
	import type { User } from '$lib/types';
	import { MessageCircleIcon } from 'lucide-svelte';

	export let user: User;

	const onOpenChat = () => {
		setInboxUser(user);
	};
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	class="flex cursor-pointer items-center justify-between gap-2 rounded-xl p-2 hover:bg-slate-100"
	on:click={onOpenChat}
>
	<div class="flex items-center gap-2">
		<span>
			<img
				src={user.avatar ?? DefaultAvatar}
				alt="avatar"
				class="h-12 w-12 rounded-full object-cover"
			/>
		</span>
		<div class="text-slate-800">
			<span class="text-lg">
				{user.email}
			</span>
			<div class="text-sm font-medium">
				{getUserName(user)}
			</div>
		</div>
	</div>
	<MessageCircleIcon class="cursor-pointer text-green-500" />
</div>
