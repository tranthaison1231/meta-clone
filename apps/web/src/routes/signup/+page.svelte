<script lang="ts">
	import { authApi } from '$lib/apis/auth';
	import logo from '$lib/assets/images/logo.svg';
	import Button from '$lib/components/ui/button/button.svelte';
	import FormItem from '$lib/components/ui/form/FormItem.svelte';
	import { signUpSchema } from '$lib/utils/schema';
	import { useMutation } from '@sveltestack/svelte-query';
	import { superForm, superValidateSync } from 'sveltekit-superforms/client';
	import { omit } from 'lodash-es';
	import { notification } from '$lib/components/ui/notification';

	const signUpMutate = useMutation(authApi.signUp);

	const { form, errors, constraints, enhance } = superForm(superValidateSync(signUpSchema), {
		SPA: true,
		validators: signUpSchema,
		taintedMessage: false,
		onUpdate: async ({ form }) => {
			if (form.valid) {
				try {
					await $signUpMutate.mutateAsync(omit(form.data, ['confirmPassword']));

					window.location.assign('/login');
				} catch (error) {
					notification.error({ title: (error as Error).message });
				}
			}
		}
	});
</script>

<div class="flex h-full w-full flex-col items-center justify-center gap-4">
	<img src={logo} alt="logo" class="w-20" />
	<h1 class="my-8 text-3xl">Create Your Account</h1>
	<form method="POST" class="flex w-80 flex-col gap-4" use:enhance>
		<FormItem label="Email" errors={$errors.email} required>
			<input
				class="input"
				placeholder="Email address"
				bind:value={$form.email}
				{...$constraints.email}
			/>
		</FormItem>
		<FormItem label="Password" errors={$errors.password} required>
			<input
				class="input"
				placeholder="Password"
				type="password"
				bind:value={$form.password}
				{...$constraints.password}
			/>
		</FormItem>
		<FormItem label="Confirm Password" errors={$errors.confirmPassword} required>
			<input
				class="input"
				type="password"
				placeholder="Confirm Password"
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
		<FormItem label="Avatar" errors={$errors.avatar}>
			<input
				class="input"
				placeholder="Paste your avatar's url here"
				bind:value={$form.avatar}
				{...$constraints.avatar}
			/>
		</FormItem>
		<Button type="submit">Sign Up</Button>
	</form>
</div>
