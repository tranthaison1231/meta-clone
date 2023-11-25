<script lang="ts">
	import { usersApi } from '$lib/apis/users';
	import DefaultAvatar from '$lib/assets/images/default-avatar.jpeg';
	import { notification } from '$lib/components/ui/notification';
	import { me } from '$lib/stores/me';
	import { FriendStatus } from '$lib/types';
	import { useMutation } from '@sveltestack/svelte-query';
	import { CheckIcon, UserCheck2, UserPlus2, XIcon } from 'lucide-svelte';

	export let email: string,
		avatar: string | undefined = undefined,
		userId: number,
		friendStatus: FriendStatus | undefined = undefined,
		refetchUserList: (() => void) | undefined = undefined;

	const addFriendMutate = useMutation(usersApi.addFriend, {
		onSuccess() {
			refetchUserList?.();
		},
		onError(error) {
			notification.error({
				title: (error as Error).message
			});
		}
	});
	const acceptFriendMutate = useMutation(usersApi.acceptFriend, {
		onSuccess() {
			refetchUserList?.();
		},
		onError(error) {
			notification.error({
				title: (error as Error).message
			});
		}
	});

	const onAddFriend = async () => {
		if ($me?.id) {
			await $addFriendMutate.mutate({
				friendId: userId,
				userId: $me?.id
			});
		}
	};

	const onAcceptFriend = async () => {
		if ($me?.id) {
			await $acceptFriendMutate.mutate({
				friendId: $me?.id,
				userId: userId
			});
		}
	};
</script>

<div
	class="flex cursor-pointer items-center justify-between gap-2 rounded-xl p-2 hover:bg-slate-100"
>
	<div class="flex items-center gap-2">
		<span>
			<img src={avatar ?? DefaultAvatar} alt="avatar" class="h-12 w-12 rounded-full object-cover" />
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
		<UserCheck2 class="text-slate-400 hover:text-slate-500" />
	{:else}
		<button on:click={onAddFriend}>
			<UserPlus2 class="text-slate-400 hover:text-slate-500" />
		</button>
	{/if}
</div>
