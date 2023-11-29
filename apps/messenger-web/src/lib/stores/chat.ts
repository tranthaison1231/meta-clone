import type { User } from '$lib/types';
import type { Chat } from '$lib/types/chat';
import { writable } from 'svelte/store';

export const inboxChat = writable<Chat | null>(null);

export const setInboxChatData = (chat: Chat) => {
	return inboxChat.update(() => chat);
};

export const inboxUsers = writable<User[]>([]);

export const setInboxUsers = (users: User[]) => {
	return inboxUsers.update(() => users);
};

export const setInboxUser = (user: User) => {
	return inboxUsers.update(() => [user]);
};
