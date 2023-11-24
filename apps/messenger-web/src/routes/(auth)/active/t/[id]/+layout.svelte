<script>
	import { BellOff, Divide, MonitorDown, MoreHorizontal } from 'lucide-svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import { chatsApi } from '$lib/apis/chats';
	import { each } from 'lodash-es';
	import { usersApi } from '$lib/apis/user';
	import UserCard from './UserCard.svelte';
	import Spinner from '$lib/components/ui/wrapper/Spinner.svelte';

	const result = useQuery(['get-all-users'], () => usersApi.getAll({}));

	$: users = $result.data?.data?.items ?? [];
	$: console.log($result)
</script>

<div class="border-r-1 w-90 relative flex h-screen flex-col gap-8 py-2">
	<div class="flex flex-col gap-2">
		<div class="border-b-1 border-solid px-4">
			<h1 class="text-2xl font-bold">Active</h1>
		</div>

		<div class="px-4">Empty List</div>
	</div>

	<div class="flex flex-col gap-2">
		<div class="border-b-1 border-solid px-4">
			<h1 class="text-2xl font-bold">Social</h1>
		</div>
		<div class="flex flex-col gap-2 px-4">
			<Spinner loading={$result.isFetching}>
				{#each users as user}
					<UserCard
						email={user.email}
						avatar={user.avatar}
						userId={user.id}
						friendStatus={user.friend_status}
						refetchUserList={$result.refetch}
					/>
				{/each}
			</Spinner>
		</div>
	</div>
</div>
<div class="h-screen w-[calc(100vw-26rem)]">
	<slot />
</div>
