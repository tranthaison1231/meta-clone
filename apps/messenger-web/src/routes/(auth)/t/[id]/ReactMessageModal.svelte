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
	import {
		angryIcon,
		hahaIcon,
		sadIcon,
		loveIcon,
		wowIcon,
		likeIcon,
		heartIcon
	} from '$lib/assets/svgs';
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
				src={hahaIcon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={angryIcon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={loveIcon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={likeIcon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={heartIcon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={sadIcon}
				alt="logo"
				class="icon-animation hover:-translate-y-6 hover:translate-x-1 hover:rotate-12 hover:scale-150
            "
			/>
			<img
				src={wowIcon}
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
