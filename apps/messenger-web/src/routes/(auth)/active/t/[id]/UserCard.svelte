<script lang="ts">
	import { chatsApi } from '$lib/apis/chats';
	import { usersApi } from '$lib/apis/users';
	import DefaultAvatar from '$lib/assets/images/default-avatar.jpeg';
	import { notification } from '$lib/components/ui/notification';
	import { getUserName } from '$lib/services/user';
	import { setInboxChatData, setInboxUser } from '$lib/stores/chat';
	import { me } from '$lib/stores/me';
	import { FriendStatus, type User } from '$lib/types';
	import { useMutation } from '@sveltestack/svelte-query';
	import { CheckIcon, UserCheck2, UserPlus2, XIcon } from 'lucide-svelte';

	export let friendStatus: FriendStatus | undefined = undefined,
		refetchUserList: (() => void) | undefined = undefined,
		user: User;

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
				friendId: user.id,
				userId: $me?.id
			});
		}
	};

	const onAcceptFriend = async (isRejecting: boolean) => {
		if ($me?.id) {
			await $acceptFriendMutate.mutate({
				friendId: $me?.id,
				userId: user.id,
				isRejecting
			});
		}
	};

	const onCardClick = async () => {
		console.log(user);

		if ($me?.id) {
			const result = await chatsApi.getAll({ memberIds: [user.id, $me?.id], isSingleChat: true });

			const firstChat = result.items[0];

			if (firstChat) {
				setInboxChatData(firstChat);
			} else {
				setInboxChatData(null);
			}

			setInboxUser(user);
		}
	};
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	class="flex cursor-pointer items-center justify-between gap-2 rounded-xl p-2 hover:bg-slate-100"
	on:click={onCardClick}
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

	{#if friendStatus === FriendStatus.Friend}
		<UserCheck2 class="text-green-600" />
	{:else if friendStatus === FriendStatus.RequireAccept}
		<div class="flex gap-4">
			<button on:click={() => onAcceptFriend(false)}>
				<CheckIcon class="rounded-lg bg-green-500 text-white hover:bg-green-600" />
			</button>
			<button on:click={() => onAcceptFriend(true)}>
				<XIcon class="rounded-lg bg-red-500 text-white hover:bg-red-600" />
			</button>
		</div>
	{:else if friendStatus === FriendStatus.Pending}
		<UserCheck2 class="text-slate-600 hover:text-slate-500" />
	{:else}
		<button on:click={onAddFriend}>
			<UserPlus2 class="text-slate-600 hover:text-slate-500" />
		</button>
	{/if}
</div>
