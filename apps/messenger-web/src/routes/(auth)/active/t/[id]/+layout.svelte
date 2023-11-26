<script>
	import { BellOff, Divide, MonitorDown, MoreHorizontal } from 'lucide-svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import { chatsApi } from '$lib/apis/chats';
	import { usersApi } from '$lib/apis/users';
	import UserCard from './UserCard.svelte';
	import Spinner from '$lib/components/ui/wrapper/Spinner.svelte';
	import { me } from '$lib/stores/me';

	const getUsersResult = useQuery(['users'], () => usersApi.getAll({ page: 1, limit: 20 }));
	const getFriendsResult = useQuery(['friends'], () => usersApi.getFriends({ userId: $me?.id }), {
		enabled: !!$me?.id
	});

	$: users = $getUsersResult.data?.items ?? [];
	$: friends = $getFriendsResult.data?.items ?? [];
</script>

<div class="relative flex h-screen w-90 flex-col gap-8 border-r-1 py-2">
	<div class="flex flex-col gap-2">
		<div class="border-b-1 border-solid px-4">
			<h1 class="text-2xl font-bold">Active</h1>
		</div>

		<div class="px-4">Empty List</div>
	</div>

	<div class="flex flex-col gap-2 overflow-y-auto">
		<div class="border-b-1 border-solid px-4">
			<h1 class="text-2xl font-bold">Social</h1>
		</div>
		<div class="flex flex-col gap-2 px-4">
			<Spinner loading={$getUsersResult.isLoading}>
				{#if users && users.length > 0}
					{#each users as user}
						<UserCard
							email={user.email}
							avatar={user.avatar}
							userId={user.id}
							friendStatus={user.friendStatus}
							refetchUserList={$getUsersResult.refetch}
						/>
					{/each}
				{:else}
					<div class="px-4">Empty List</div>
				{/if}
			</Spinner>
		</div>
	</div>
</div>
<div class="h-screen w-[calc(100vw-26rem)]">
	<slot />
</div>
