<script lang="ts">
	import { goto } from '$app/navigation';
	import { authApi } from '$lib/apis/auth';
	import { logoIcon } from '$lib/assets/svgs';
	import Button from '$lib/components/ui/button/button.svelte';
	import FormItem from '$lib/components/ui/form/FormItem.svelte';
	import { notification } from '$lib/components/ui/notification';
	import { signUpSchema } from '$lib/utils/schema';
	import { useMutation } from '@sveltestack/svelte-query';
	import { superForm, superValidateSync } from 'sveltekit-superforms/client';
	import BGSignup from '$lib/assets/images/signup-background.png';

	const signUpMutate = useMutation(authApi.signUp);

	const { form, errors, constraints, enhance } = superForm(superValidateSync(signUpSchema), {
		SPA: true,
		validators: signUpSchema,
		taintedMessage: false,
		onUpdate: async ({ form }) => {
			if (form.valid) {
				try {
					const { confirmPassword: _, ...signUpDto } = form.data;
					await $signUpMutate.mutateAsync(signUpDto);
					goto('/');
				} catch (error) {
					notification.error({ title: (error as Error).message });
				}
			}
		}
	});
</script>

<div class="mx-auto h-full w-full max-w-6xl px-10 py-10 lg:px-0">
	<div class="flex h-5 w-full items-center justify-between">
		<img src={logoIcon} alt="logo" class="w-11" />
		<div class="hidden cursor-pointer gap-7 text-lg font-semibold lg:flex">
			<p>Features</p>
			<p>Desktop app</p>
			<p>Privacy and safety</p>
			<p>For developers</p>
		</div>
	</div>
	<div class="flex items-center justify-center lg:justify-between">
		<div class="flex flex-col">
			<h1 class="text-linear my-5 mt-8 max-w-sm text-4xl font-bold lg:text-6xl">
				Create your account, anywhere
			</h1>
			<p class="mb-10 max-w-sm text-base lg:text-xl">
				Messenger makes it easy and fun to stay close to your favourite people.
			</p>
			<form method="POST" class="w-sm flex flex-col gap-4" use:enhance>
				<div class="flex gap-2">
					<FormItem label="First Name" errors={$errors.firstName} required>
						<input
							class="input text-style"
							placeholder="Enter your first Name"
							bind:value={$form.firstName}
							{...$constraints.firstName}
						/>
					</FormItem>
					<FormItem label="Last Name" errors={$errors.lastName} required>
						<input
							class="input text-style"
							placeholder="Enter your last Name"
							bind:value={$form.lastName}
							{...$constraints.lastName}
						/>
					</FormItem>
				</div>
				<FormItem label="Email" errors={$errors.email} required>
					<input
						class="input text-style"
						placeholder="Enter your email address"
						bind:value={$form.email}
						{...$constraints.email}
					/>
				</FormItem>

				<FormItem label="Password" errors={$errors.password} required>
					<input
						class="input text-style"
						placeholder="Enter your password"
						type="password"
						bind:value={$form.password}
						{...$constraints.password}
					/>
				</FormItem>
				<FormItem label="Confirm Password" errors={$errors.confirmPassword} required>
					<input
						class="input text-style"
						type="password"
						placeholder="Enter your confirm Password"
						bind:value={$form.confirmPassword}
						{...$constraints.confirmPassword}
					/>
				</FormItem>
				<FormItem label="Gender" layout="horizontal" errors={$errors.gender} required>
					<select bind:value={$form.gender} class="border border-solid p-2">
						<option value="" disabled selected>Select your gender</option>
						<option value="male">Male</option>
						<option value="female">Female</option>
					</select>
				</FormItem>
				<FormItem label="Upload your avatar" errors={$errors.avatar}>
					<input
						class="input text-style"
						placeholder="Enter your avatar's url here"
						bind:value={$form.avatar}
						{...$constraints.avatar}
					/>
				</FormItem>
				<Button class="w-fit rounded-3xl bg-[#0a7cff] p-6 font-bold" type="submit">Sign Up</Button>
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
	select {
		/* Reset */
		appearance: none;
		border: 0;
		outline: 0;
		font: inherit;
		/* Personalize */
		width: 100%;
		padding: 0.5rem 4rem 0.5rem 1rem;
		background: #0000000a;
		border-radius: 0.75rem;
		cursor: pointer;
		&::-ms-expand {
			display: none;
		}

		&:focus {
			outline: none;
		}
	}
</style>
