import { getToken } from './storage';

export const BASE_URL = import.meta.env.VITE_API_URL;

const { fetch: originalFetch } = window;

export const request = async (...args: Parameters<typeof originalFetch>) => {
	let [input, init] = args;
	input = input.toString().startsWith('http') ? input : `${BASE_URL}${input}`;
	const token = getToken();

	if (token) {
		init = {
			...init,
			headers: {
				...init?.headers,
				Authorization: `Bearer ${token}`
			}
		};
	}

	const res = await originalFetch(input, init);
	const data = await res.json();
	return data;
};
