<script lang="ts">
	import { goto } from '$app/navigation';
	import { authApi } from '$lib/apis/auth';
	import logo from '$lib/assets/images/logo.svg';
	import Button from '$lib/components/ui/button/button.svelte';
	import { setToken } from '$lib/services/storage';
	import { loginSchema } from '$lib/utils/schema';
	import { useMutation } from '@sveltestack/svelte-query';
	import { superForm, superValidateSync } from 'sveltekit-superforms/client';

	const loginMutate = useMutation(authApi.login);

	const { form, errors, constraints, enhance } = superForm(superValidateSync(loginSchema), {
		SPA: true,
		validators: loginSchema,
		taintedMessage: false,
		onUpdate: async ({ form }) => {
			if (form.valid) {
				const result = await $loginMutate.mutateAsync(form.data);
				if (result.data?.token) {
					goto('/');
					setToken(result.data?.token);
				}
			}
		}
	});
</script>

<div class="flex w-full items-center justify-center">
	<div class="flex flex-col items-center">
		<img src={logo} alt="logo" class="w-20" />
		<h1 class="my-11 text-4xl">Connect with your favourite people</h1>
		<form method="POST" class="flex w-80 flex-col items-center" use:enhance>
			<input
				class="input"
				placeholder="Email address"
				bind:value={$form.email}
				{...$constraints.email}
			/>
			{#if $errors.email}<span class="mt-1 w-full text-red-500">{$errors.email}</span>{/if}
			<input
				class="input mt-3"
				type="password"
				placeholder="Password"
				bind:value={$form.password}
				{...$constraints.password}
			/>
			{#if $errors.password}<span class="mt-1 w-full text-red-500">{$errors.password}</span>{/if}
			<Button class="btn mt-6 w-fit rounded-full" type="submit">Continue</Button>
		</form>
	</div>
</div>
