<script>
	import { BellOff, MonitorDown } from 'lucide-svelte';
	import MoreHorizontalModal from './MoreHorizontalModal.svelte';

	const CONTACTS = [
		{
			user: {
				id: '2',
				name: 'Linh Mõm',
				avatar:
					'https://scontent.fhan2-4.fna.fbcdn.net/v/t39.30808-1/395517388_3705379936359673_6358498159852438917_n.jpg?stp=dst-jpg_p100x100&_nc_cat=110&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeEo9Zc9spiAH4sSqlTI4nF9Yc1lCiUst4VhzWUKJSy3hSXGUG3al7W4ZrfT1XeLNsiGb-auNv4vuZStevKYL-7g&_nc_ohc=kLQCIcOS7k4AX84wSpP&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.fhan2-4.fna&oh=00_AfCCLn3PT2Q6rSNJpn7Zg68jqu654GUgbq1Ft24IiPJr7Q&oe=655CADAA'
			},
			lastMessage: {
				text: 'bạn dạo ni sao r',
				timestamp: '1700204276',
				sender: {
					id: '1'
				}
			}
		}
	];
</script>

<div class="relative h-screen w-90 border-r-1 py-2">
	<div class="px-4">
		<h1 class="text-2xl font-bold">Chats</h1>
		<input class="input" placeholder="Search Messenger" />
	</div>
	<div class=" mt-1 p-1.5">
		{#each CONTACTS as { user, lastMessage }}
			<div class="group flex items-center justify-between gap-3 rounded-xl bg-gray-100 p-2">
				<div class="flex gap-3">
					<img src={user.avatar} alt="" class="h-12 w-12 rounded-full" />
					<div>
						<p>{user.name}</p>
						<p class="text-sm text-gray-500">You: {lastMessage.text}</p>
					</div>
				</div>
				<div class="flex items-center gap-2">
					<MoreHorizontalModal />
					<BellOff class="text-gray-500" size={20} />
				</div>
			</div>
		{/each}
	</div>
	<div class="absolute bottom-0 w-full border border-t-1 px-2 py-3">
		<button class="flex w-full justify-center gap-2 rounded-xl px-3 py-2 hover:bg-gray-100">
			<MonitorDown />
			Try Messenger for Mac
		</button>
	</div>
</div>
<div class="h-screen w-[calc(100vw-26rem)]">
	<slot />
</div>
