<script>
	import '../app.css';
	import { Notification } from '$lib/components/ui/notification';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				refetchIntervalInBackground: false,
				refetchOnWindowFocus: false,
				retry: false
			}
		}
	});
</script>

<svelte:head>
	<title>Messenger</title>
</svelte:head>
<Notification />

<QueryClientProvider client={queryClient}>
	<div class="flex h-screen">
		<slot />
	</div>
</QueryClientProvider>
