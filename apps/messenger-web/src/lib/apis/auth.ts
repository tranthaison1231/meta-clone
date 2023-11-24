import { BASE_URL } from '$lib/services/request';
import type { SignInRequest, SignUpRequest, User } from '$lib/types';
import type { BaseResponseType } from '$lib/types/response';

export interface GetMeResponse {
	user: User
}

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
		return await rest.json() as BaseResponseType<GetMeResponse>;
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
