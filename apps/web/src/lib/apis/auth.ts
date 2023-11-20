import { BASE_URL } from '$lib/services/request';

export const authApi = {
	login: async ({ email, password }: { email: string; password: string }) => {
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
		const rest = await fetch(`${BASE_URL}/me`);
		return rest.json();
	}
};
