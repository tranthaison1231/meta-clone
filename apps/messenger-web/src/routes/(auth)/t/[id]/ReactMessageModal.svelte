<script>
	import * as Popover from '$lib/components/ui/popover';
	import { cn } from '$lib/utils/cn';
	import {
		AlertTriangle,
		Archive,
		BellOff,
		Camera,
		Check,
		LogOut,
		MoreHorizontal,
		Phone,
		Trash,
		Video
	} from 'lucide-svelte';
	import { twMerge } from 'tailwind-merge';
	import angry_icon from '$lib/assets/images/angry.svg';
	import haha_icon from '$lib/assets/images/haha.svg';
	import sad_icon from '$lib/assets/images/sad.svg';
	import love_icon from '$lib/assets/images/love.svg';
	import wow_icon from '$lib/assets/images/wow.svg';
	import like_icon from '$lib/assets/images/like.svg';
	import heart_icon from '$lib/assets/images/heart.svg';
</script>

<Popover.Root>
	<Popover.Trigger>
		<div class="rounded-full bg-[#F5F5F5] p-1 shadow-xl">
			<MoreHorizontal />
		</div>
	</Popover.Trigger>
	<Popover.Content class="w-90 relative h-14 items-center px-2 py-1">
		<div class="flex items-start justify-between gap-2">
			<img
				src={haha_icon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={angry_icon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={love_icon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={like_icon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={heart_icon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={sad_icon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={wow_icon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
		</div>
	</Popover.Content>
</Popover.Root>

<style lang="postcss">
	.item {
		@apply flex cursor-pointer items-center gap-2 rounded-xl p-2;
		&:hover {
			@apply bg-gray-100;
		}
	}
	.btn {
		@apply rounded-full bg-gray-200 p-2;
	}
	.icon-animation {
		@apply h-10 w-10 transition-all;
	}
</style>
