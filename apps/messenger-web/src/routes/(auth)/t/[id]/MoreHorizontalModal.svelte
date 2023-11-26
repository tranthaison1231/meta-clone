<script>
	import * as Popover from '$lib/components/ui/popover';
	import { cn } from '$lib/utils/cn';
	import {
		AlertTriangle,
		Archive,
		BellOff,
		Check,
		LogOut,
		MoreHorizontal,
		Phone,
		Trash,
		Video
	} from 'lucide-svelte';

	let open;
</script>

<Popover.Root bind:open>
	<Popover.Trigger>
		<button
			class={cn(
				'hidden rounded-full bg-white p-1.5 shadow-md hover:bg-gray-100 group-hover:block',
				{
					block: open
				}
			)}
		>
			<MoreHorizontal />
		</button>
	</Popover.Trigger>
	<Popover.Content class="w-80 px-2 py-1">
		<div class="item">
			<button class="btn">
				<Check size={16} />
			</button>
			<p>Mark as unread</p>
		</div>
		<div class="item">
			<button class="btn">
				<BellOff size={16} />
			</button>
			<p>Unmute notifications</p>
		</div>
		<hr class="my-2" />
		<div class="item">
			<button class="btn">
				<Phone size={16} />
			</button>
			<p>Audio Call</p>
		</div>
		<div class="item">
			<button class="btn">
				<Video size={16} />
			</button>
			<p>Video Chat</p>
		</div>
		<hr class="my-2" />
		<div class="item">
			<button class="btn">
				<Archive size={16} />
			</button>
			<p>Archive Chat</p>
		</div>
		<div class="item">
			<button class="btn">
				<Trash size={16} />
			</button>
			<p>Delete Chat</p>
		</div>
		<div class="item">
			<button class="btn">
				<AlertTriangle size={16} />
			</button>
			<p>Report</p>
		</div>
		<div class="item">
			<button class="btn">
				<LogOut size={16} />
			</button>
			<p>Leave Group</p>
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
		@apply rounded-full bg-gray-100 p-2;
	}
</style>
