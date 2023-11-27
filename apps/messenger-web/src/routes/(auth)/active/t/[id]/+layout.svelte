<script>
	import { usersApi } from '$lib/apis/users';
	import Spinner from '$lib/components/ui/wrapper/Spinner.svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import UserCard from './UserCard.svelte';
	import { me } from '$lib/stores/me';
	import FriendCard from './FriendCard.svelte';

	const getUsersResult = useQuery(
		['users', $me?.id],
		() => usersApi.getAll({ page: 1, limit: 20 }),
		{
			keepPreviousData: false,
			refetchOnMount: true
		}
	);

	const getFriendResult = useQuery(
		['friends', $me?.id],
		() => usersApi.getFriends({ page: 1, limit: 20, userId: $me?.id }),
		{
			enabled: !!$me?.id
		}
	);

	$: users = $getUsersResult?.data?.items ?? [];
	$: friends = $getFriendResult?.data?.items ?? [];
</script>

<div class="w-90 border-r-1 relative flex h-screen flex-col gap-8 py-2">
	<div class="flex flex-col gap-2">
		<div class="border-b-1 border-solid px-6">
			<h1 class="text-2xl font-bold">Active</h1>
		</div>
		<div class="flex flex-col gap-2 overflow-y-auto px-4">
			<Spinner loading={$getFriendResult.isLoading}>
				{#if friends && friends.length > 0}
					{#each friends as user}
						<FriendCard {user} />
					{/each}
				{:else}
					<div class="px-4">Empty List</div>
				{/if}
			</Spinner>
		</div>
	</div>

	<div class="flex flex-col gap-2">
		<div class="border-b-1 border-solid px-6">
			<h1 class="text-2xl font-bold">Social</h1>
		</div>
		<div class="flex flex-col gap-2 overflow-y-auto px-4">
			<Spinner loading={$getUsersResult.isLoading}>
				{#if users && users.length > 0}
					{#each users as user}
						<UserCard
							{user}
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
