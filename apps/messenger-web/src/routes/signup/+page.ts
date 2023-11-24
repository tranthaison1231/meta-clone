import { redirect } from '@sveltejs/kit';

export function load() {
	const isLoggedIn = localStorage.getItem('token');
	if (isLoggedIn) {
		throw redirect(307, '/');
	}
}
