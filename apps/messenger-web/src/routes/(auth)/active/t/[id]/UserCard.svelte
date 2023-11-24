<script lang="ts">
	import { usersApi } from '$lib/apis/user';
	import DefaultAvatar from '$lib/assets/images/default-avatar.jpeg';
	import { me } from '$lib/stores/me';
	import { FriendStatus } from '$lib/types';
	import { useMutation } from '@sveltestack/svelte-query';
	import { CheckIcon, UserCheck2, UserPlus2, XIcon } from 'lucide-svelte';

	export let email: string,
		avatar: string | undefined = undefined,
		userId: number,
		friendStatus: FriendStatus | undefined = undefined;

	const addFriendMutate = useMutation(usersApi.addFriend);
	const acceptFriendMutate = useMutation(usersApi.acceptFriend);

	const onAddFriend = async () => {
		if ($me?.id) {
			await $addFriendMutate.mutate({
				friend_id: userId,
				user_id: $me?.id
			});
		}
	};

	const onAcceptFriend = async () => {
		if ($me?.id) {
			await $acceptFriendMutate.mutate({
				friend_id: $me?.id,
				user_id: userId
			});
		}
	};
</script>

<div
	class="flex cursor-pointer items-center justify-between gap-2 rounded-xl p-2 hover:bg-slate-100"
>
	<div class="flex items-center gap-2">
		<span>
			<img src={avatar ?? DefaultAvatar} alt="avatar" class="rounded-full" width="45px" />
		</span>
		<span>
			{email}
		</span>
	</div>

	{#if friendStatus === FriendStatus.Friend}
		<UserCheck2 class="text-green-600" />
	{:else if friendStatus === FriendStatus.RequireAccept}
		<div class="flex gap-4">
			<button on:click={onAcceptFriend}>
				<CheckIcon class="rounded-lg bg-green-500 text-white hover:bg-green-600" />
			</button>
			<button on:click={onAddFriend}>
				<XIcon class="rounded-lg bg-red-500 text-white hover:bg-red-600" />
			</button>
		</div>
	{:else if friendStatus === FriendStatus.Pending}
		<button on:click={onAddFriend}>
			<UserCheck2 class="text-slate-400 hover:text-slate-500" />
		</button>
	{:else}
		<button on:click={onAddFriend}>
			<UserPlus2 class="text-slate-400 hover:text-slate-500" />
		</button>
	{/if}
</div>
