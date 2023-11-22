import { BASE_URL } from '$lib/services/request';
import type { SignInRequest, SignUpRequest } from '$lib/types';

export const authApi = {
	login: async ({ email, password }: SignInRequest) => {
		const rest = await fetch(`${BASE_URL}/sign-in`, {
			method: 'POST',
			body: JSON.stringify({
				email: email,
				password: password
			})
		});
		return rest.json();
	},
	getMe: async () => {
		const rest = await fetch(`${BASE_URL}/me`, {
			headers: {
				Authorization: `Bearer ${localStorage.getItem('token')}`
			}
		});
		return rest.json();
	},
	signUp: async ({email, gender, password, avatar }: SignUpRequest) => {
		const rest = await fetch(`${BASE_URL}/sign-up`, {
			method: 'POST',
			body: JSON.stringify({
				email,
				gender,
				password,
				avatar
			})
		});
		return rest.json();
	}
};
