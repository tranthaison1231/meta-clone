import { redirect } from '@sveltejs/kit';

export function load() {
	const defaultID = 124124;
	throw redirect(307, `/t/${defaultID}`);
}
