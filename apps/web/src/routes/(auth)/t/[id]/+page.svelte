<script>
	import {
		Image,
		MoreHorizontal,
		Phone,
		Plus,
		PlusCircle,
		Smile,
		StickyNote,
		ThumbsUp,
		Video
	} from 'lucide-svelte';

	const user = {
		id: '2',
		name: 'Linh Mõm',
		avatar:
			'https://scontent.fhan2-4.fna.fbcdn.net/v/t39.30808-1/395517388_3705379936359673_6358498159852438917_n.jpg?stp=dst-jpg_p100x100&_nc_cat=110&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeEo9Zc9spiAH4sSqlTI4nF9Yc1lCiUst4VhzWUKJSy3hSXGUG3al7W4ZrfT1XeLNsiGb-auNv4vuZStevKYL-7g&_nc_ohc=kLQCIcOS7k4AX84wSpP&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.fhan2-4.fna&oh=00_AfCCLn3PT2Q6rSNJpn7Zg68jqu654GUgbq1Ft24IiPJr7Q&oe=655CADAA'
	};
</script>

<div class="relative h-screen">
	<div class="flex h-14 w-full items-center justify-between px-4 shadow-md">
		<div class="flex gap-2">
			<img
				src="https://scontent.fhan2-4.fna.fbcdn.net/v/t39.30808-1/395517388_3705379936359673_6358498159852438917_n.jpg?stp=dst-jpg_p100x100&_nc_cat=110&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeEo9Zc9spiAH4sSqlTI4nF9Yc1lCiUst4VhzWUKJSy3hSXGUG3al7W4ZrfT1XeLNsiGb-auNv4vuZStevKYL-7g&_nc_ohc=kLQCIcOS7k4AX84wSpP&_nc_ad=z-m&_nc_cid=0&_nc_ht=scontent.fhan2-4.fna&oh=00_AfCCLn3PT2Q6rSNJpn7Zg68jqu654GUgbq1Ft24IiPJr7Q&oe=655CADAA"
				alt=""
				class="h-11 w-11 rounded-full"
			/>
			<div>
				<p>{user.name}</p>
				<p class="text-sm text-gray-500">Active now</p>
			</div>
		</div>
		<div class="flex gap-4 text-blue-500">
			<Phone />
			<Video />
			<MoreHorizontal />
		</div>
	</div>
	<div class="absolute bottom-0 flex h-14 w-full items-center px-2 py-3">
		<div class="flex text-blue-500">
			<button class="rounded-full p-2 hover:bg-gray-100">
				<PlusCircle />
			</button>
			<button class="rounded-full p-2 hover:bg-gray-100">
				<Image />
			</button>
			<button class="rounded-full p-2 hover:bg-gray-100">
				<StickyNote />
			</button>
			<button class="rounded-full p-2 hover:bg-gray-100">
				<Image />
			</button>
		</div>
		<div class="relative ml-4 w-full">
			<input class="input rounded-2xl bg-gray-200" placeholder="Aa" />
			<Smile class="absolute bottom-2 right-2 text-blue-500" />
		</div>
		<button class="ml-2 rounded-full p-2 text-blue-500 hover:bg-gray-100">
			<ThumbsUp />
		</button>
	</div>
</div>
