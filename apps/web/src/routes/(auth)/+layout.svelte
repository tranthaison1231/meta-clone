<script lang="ts">
	import { page } from '$app/stores';
	import { cn } from '$lib/utils/cn';
	import { MessageCircle, MessageSquare, PanelLeft, Store, Trash, Users } from 'lucide-svelte';
	import ProfilePopover from './ProfilePopover.svelte';
	import Groups from './Groups.svelte';

	let isSidebarOpen = false;

	let id = $page.url.pathname.split('/t/')[1];
	const SIDEBAR = [
		{
			to: `/t/${id}`,
			icon: MessageCircle,
			name: 'Chat'
		},
		{
			to: `/active/t/${id}`,
			icon: Users,
			name: 'People'
		},
		{
			to: `/marketplace/t/${id}`,
			icon: Store,
			name: 'Marketplace'
		},
		{
			to: `/requests/t/${id}`,
			icon: MessageSquare,
			name: 'Requests'
		},
		{
			to: `/archived/t/${id}`,
			icon: Trash,
			name: 'Archived'
		}
	];
</script>

<div class="flex">
	<div
		class={cn('border-r-1 gap-y relative flex h-screen w-14 flex-col items-center p-2', {
			'w-60': isSidebarOpen
		})}
	>
		{#each SIDEBAR as { to, icon: Icon, name }}
			<a
				href={to}
				class={cn('rounded-xl p-2.5', {
					'bg-gray-100': $page.url.pathname === to,
					'flex w-full gap-2': isSidebarOpen
				})}
			>
				<Icon />
				{#if isSidebarOpen}
					<p>{name}</p>
				{/if}
			</a>
		{/each}
		<hr class="my-6 h-0.5 w-full bg-gray-100" />
		<Groups />
		<div
			class={cn('absolute bottom-0 flex flex-col items-center space-y-3 p-2', {
				'w-full cursor-pointer flex-row items-center justify-between gap-4': isSidebarOpen
			})}
		>
			<ProfilePopover {isSidebarOpen} />
			<button
				class="rounded-full bg-gray-100 p-1"
				on:click={() => (isSidebarOpen = !isSidebarOpen)}
			>
				<PanelLeft />
			</button>
		</div>
	</div>
	<slot />
</div>
