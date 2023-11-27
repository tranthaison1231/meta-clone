import type { User } from '$lib/types';

export function getUserName(user: Partial<User>) {
	return [user.firstName ?? '', user.lastName ?? ''].join(' ').trim();
}
