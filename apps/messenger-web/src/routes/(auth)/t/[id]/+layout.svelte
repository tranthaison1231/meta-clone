<script lang="ts">
	import { MonitorDown } from 'lucide-svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import { chatsApi } from '$lib/apis/chats';
	import { me } from '$lib/stores/me';
	import Loading from '$lib/components/ui/Loading/Loading.svelte';
	import Inbox from './Inbox.svelte';
	import { setInboxChatData, setInboxUsers } from '$lib/stores/chat';

	$: getChatByMemberIds = useQuery(
		['chats', $me?.id],
		() =>
			chatsApi.getAll({
				memberIds: [String($me?.id)]
			}),
		{
			enabled: !!$me?.id
		}
	);

	$: chats = $getChatByMemberIds.data?.items;

	const onInboxClick = async (id: string) => {
		const { chat } = await chatsApi.getChatById(id);
		setInboxChatData(chat);
		setInboxUsers(chat.members);
	};

	const CONTACTS_ONLINE = [
		{
			user: {
				id: '1',
				name: 'Minh Vip',
				avatar:
					'https://kenh14cdn.com/thumb_w/660/203336854389633024/2021/5/27/photo2021-05-2712-11-40-1622093561531643326457.jpg',
				isOnline: true
			}
		},
		{
			user: {
				id: '2',
				name: 'Kiên Mõm',
				avatar:
					'https://gamek.mediacdn.vn/133514250583805952/2021/6/25/photo-1-16246064774091716463237.jpg',
				isOnline: true
			}
		},
		{
			user: {
				id: '3',
				name: 'Đạt Vip',
				avatar: 'https://nguoinoitieng.tv/images/nnt/97/0/bb39.jpg',
				isOnline: true
			}
		},
		{
			user: {
				id: '4',
				name: 'Sơn Trần',
				avatar:
					'https://t3.ftcdn.net/jpg/05/56/29/10/360_F_556291020_q2ieMiOCKYbtoLITrnt7qcSL1LJYyWrU.jpg',
				isOnline: false
			}
		}
	];
</script>

<div class="relative h-screen w-90 border-r-1 py-2 transition-all max-lg:w-20">
	<div class="px-4 max-lg:hidden">
		<h1 class="my-4 mb-7 text-2xl font-bold">Chats</h1>
		<input class="input border-none bg-[#F5F5F5] outline-none" placeholder="Search (Ctrl + K)" />
	</div>
	<div class="mt-5 px-4 max-lg:hidden">
		<div class="flex gap-5">
			{#each CONTACTS_ONLINE as { user }}
				<div class="flex cursor-pointer flex-col items-center">
					<div class="relative">
						<img src={user.avatar} alt="" class=" h-12 w-12 rounded-full object-cover" />

						{#if user.isOnline}
							<div
								class="absolute bottom-0 right-0 h-5 w-5 translate-x-1 rounded-full border-4 border-white bg-green-500"
							/>
						{/if}
					</div>
					<div class="w-10 max-lg:hidden">
						<p class="break-words text-center text-sm">{user.name}</p>
					</div>
				</div>
			{/each}
		</div>
	</div>
	<div class="mt-5">
		{#if $getChatByMemberIds.isLoading}
			<Loading />
		{:else if chats && chats?.length > 0}
			{#each chats as chat}
				<Inbox {chat} onClick={onInboxClick} />
			{/each}
		{:else}
			<div class="flex h-full items-center justify-center px-4 text-lg">
				You Have No Conversation
			</div>
		{/if}
	</div>
	<div class="absolute bottom-0 flex h-14 w-full items-center border-t-1 px-2">
		<button class="flex w-full justify-center gap-2 rounded-xl px-3 hover:bg-gray-100">
			<MonitorDown />
			Try Messenger for Mac
		</button>
	</div>
</div>
<div class="h-screen w-[calc(100vw-26rem)] max-lg:w-[calc(100vw-8.5rem)]">
	<slot />
</div>
