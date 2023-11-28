<script lang="ts">
	import { goto } from '$app/navigation';
	import { authApi } from '$lib/apis/auth';
	import { logoIcon } from '$lib/assets/svgs';
	import Button from '$lib/components/ui/button/button.svelte';
	import { setToken } from '$lib/services/storage';
	import { loginSchema } from '$lib/utils/schema';
	import { useMutation } from '@sveltestack/svelte-query';
	import { superForm, superValidateSync } from 'sveltekit-superforms/client';
	import BGSignup from '$lib/assets/images/signup-background.png';
	import { cn } from '$lib/utils/cn';

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

<div class="mx-auto h-full w-full max-w-6xl px-10 py-10 lg:px-0">
	<div class="mb-20 flex h-5 w-full items-center justify-between">
		<img src={logoIcon} alt="logo" class="w-11" />
		<div class="hidden cursor-pointer gap-7 text-lg font-semibold lg:flex">
			<p>Features</p>
			<p>Desktop app</p>
			<p>Privacy and safety</p>
			<p>For developers</p>
		</div>
	</div>
	<div class="flex items-start justify-center lg:justify-between">
		<div class="flex flex-col">
			<h1
				class="text-linear my-5 mt-8 max-w-sm text-4xl font-bold lg:text-[80px] lg:leading-[80px]"
			>
				Hang out anytime, anywhere
			</h1>
			<p class="mb-10 max-w-sm text-base lg:text-xl">
				Messenger makes it easy and fun to stay close to your favourite people.
			</p>
			<form method="POST" class="w-sm flex flex-col gap-4" use:enhance>
				<input
					class={cn(
						'input text-style h-12 transition-all',
						$errors.email &&
							'!border-1 !border-solid !border-red-300 !bg-red-100 placeholder:text-red-500'
					)}
					placeholder="Email address"
					bind:value={$form.email}
					{...$constraints.email}
				/>
				<input
					class={cn(
						'input text-style h-12 transition-all',
						$errors.password &&
							'!border-1 !border-solid !border-red-300 !bg-red-100 placeholder:text-red-500'
					)}
					type="password"
					placeholder="Password"
					bind:value={$form.password}
					{...$constraints.password}
				/>
				<div class="flex items-center gap-5">
					<Button class="w-fit rounded-3xl bg-[#0a7cff] p-6 font-bold" type="submit">Log in</Button>
					<a href="/signup" class="cursor-pointer font-medium text-[#0a7cff] underline"
						>Create new account?</a
					>
				</div>
			</form>
		</div>
		<img src={BGSignup} alt="logo" class="hidden w-full max-w-[650px] lg:block" />
	</div>
</div>

<style>
	.text-linear {
		background: -webkit-linear-gradient(83.84deg, #0088ff -6.87%, #a033ff 26.54%, #ff5c87 58.58%);
		-webkit-background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.text-style {
		@apply rounded-xl border-none bg-[#0000000a] px-4;
	}
</style>
