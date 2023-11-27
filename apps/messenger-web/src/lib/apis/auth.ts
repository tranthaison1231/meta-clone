import { request } from '$lib/services/request';
import type { SignInRequest, SignUpRequest, User } from '$lib/types';
import type { BaseResponseType } from '$lib/types/response';

export interface GetMeResponse {
	user: User;
}

export const authApi = {
	login: async ({ email, password }: SignInRequest) => {
		const data = await request(`/sign-in`, {
			method: 'POST',
			body: JSON.stringify({
				email: email,
				password: password
			})
		});
		return data;
	},
	getMe: async () => {
		const data: BaseResponseType<GetMeResponse> = await request(`/me`);
		return data;
	},
	signUp: async (input: SignUpRequest) => {
		const data = await request('/sign-up', {
			method: 'POST',
			body: JSON.stringify({
				...input
			})
		});
		return data;
	}
};
