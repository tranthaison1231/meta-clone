<script lang="ts">
	import { goto } from '$app/navigation';
	import * as Popover from '$lib/components/ui/popover';
	import { cn } from '$lib/utils/cn';
	import { LogOut } from 'lucide-svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import { authApi } from '$lib/apis/auth';

	export let isSidebarOpen = false;

	const logout = () => {
		localStorage.removeItem('token');
		goto('/login');
	};

	const result = useQuery(['me'], () => authApi.getMe());

	$: me = $result.data?.data?.user;
</script>

<Popover.Root
	positioning={{
		placement: 'top'
	}}
>
	<Popover.Trigger>
		<div
			class={cn('flex', {
				'w-full items-center gap-2 rounded-lg p-2 hover:bg-gray-100': isSidebarOpen
			})}
		>
			<img src={me?.avatar} class="h-8 w-8 rounded-full" alt="avatar" />
			{#if isSidebarOpen}
				<p>{me?.email}</p>
			{/if}
		</div>
	</Popover.Trigger>
	<Popover.Content class="w-80 px-2 py-1">
		<button
			class="flex w-full cursor-pointer gap-3 rounded-md p-2 hover:bg-gray-100"
			on:click={logout}
		>
			<LogOut />
			Log Out
		</button>
	</Popover.Content>
</Popover.Root>
