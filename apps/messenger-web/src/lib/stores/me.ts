import type { User } from '$lib/types';
import { get, writable } from 'svelte/store';

export const me = writable<User | null>(null);

export const isAddedFriend = (friendId: number) => {
	return !!get(me)?.friends?.find((user) => user.id === friendId);
};
