<script>
	import {
		Image,
		MoreHorizontal,
		Phone,
		PlusCircle,
		Smile,
		StickyNote,
		ThumbsUp,
		Video
	} from 'lucide-svelte';
	import ReactMessageModal from './ReactMessageModal.svelte';
	import MoreHorizontalModal from './MoreHorizontalModal.svelte';

	const user = {
		id: '2',
		name: 'Linh Mõm',
		avatar:
			'https://kenh14cdn.com/thumb_w/660/203336854389633024/2021/5/27/photo2021-05-2712-11-40-1622093561531643326457.jpg'
	};
</script>

<div class="relative h-screen">
	<div class="flex h-14 w-full items-center justify-between px-4 shadow-sm">
		<div class="flex items-center gap-2">
			<img
				src="https://kenh14cdn.com/thumb_w/660/203336854389633024/2021/5/27/photo2021-05-2712-11-40-1622093561531643326457.jpg"
				alt=""
				class="h-8 w-8 rounded-full object-cover"
			/>
			<div>
				<p class="font-bold">{user.name}</p>
				<p class="text-sm text-gray-500">Active now</p>
			</div>
		</div>
		<div class="flex gap-4 text-transparent">
			<Phone class="fill-[#AA00FF]" />
			<Video class="fill-[#AA00FF] text-[#AA00FF]" />
			<MoreHorizontal class="fill-[#AA00FF]" />
		</div>
	</div>
	<div class="flex h-[calc(100vh-7rem)] w-full flex-col justify-end px-4 py-2">
		<div class="message-reverse">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message group">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
			<ReactMessageModal />
		</div>
		<div class="message-reverse">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message-reverse">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message-reverse">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message-reverse">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message-reverse">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
		</div>
		<div class="message">
			<p class="test">yeh có cái khay đồ ăn của nó r</p>
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
			<input class="input rounded-xl border-none bg-gray-100" placeholder="Aa" />
			<Smile class="absolute bottom-2 right-2 text-blue-500" />
		</div>
		<button class="ml-2 rounded-full p-2 text-blue-500 hover:bg-gray-100">
			<ThumbsUp />
		</button>
	</div>
</div>

<style>
	.message {
		@apply flex w-full gap-2;
	}
	.message-reverse {
		@apply flex w-full flex-row-reverse gap-2;
	}
	.test {
		@apply h-auto max-w-[250px] break-all rounded-full bg-[#F5F5F5] px-4 py-2;
	}
</style>
