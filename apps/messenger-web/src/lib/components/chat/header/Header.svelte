<script lang="ts">
	import { MoreHorizontal, Phone, Video } from 'lucide-svelte';
	import { inboxUsers } from '$lib/stores/chat';
	import { getUserName } from '$lib/services/user';
	import defaultAvatar from '$lib/assets/images/default-avatar.jpeg';

	$: firstUser = $inboxUsers?.[0];
	$: userName = firstUser ? getUserName(firstUser) : 'Anonymous';
	$: userAvatar = firstUser?.avatar ?? defaultAvatar;
</script>

<div class="flex h-14 w-full items-center justify-between border-b px-4">
	<div class="flex items-center gap-2">
		<img src={userAvatar} alt="" class="h-8 w-8 rounded-full object-cover" />
		<div>
			<p class="font-bold">{userName}</p>
			<p class="text-sm text-gray-500">Active now</p>
		</div>
	</div>
	<div class="flex gap-4 text-transparent">
		<Phone class="fill-[#AA00FF]" />
		<Video class="fill-[#AA00FF] text-[#AA00FF]" />
		<MoreHorizontal class="fill-[#AA00FF]" />
	</div>
</div>
