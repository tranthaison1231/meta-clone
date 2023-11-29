import type { Message, User } from '$lib/types';
import type { Chat } from '$lib/types/chat';
import { writable } from 'svelte/store';

export const inboxChat = writable<Chat | null>(null);

export const setInboxChatData = (chat: Chat | null) => {
	return inboxChat.update(() => chat);
};

export const inboxUsers = writable<User[]>([]);

export const setInboxUsers = (users: User[]) => {
	return inboxUsers.set(users);
};

export const setInboxUser = (user: User) => {
	return inboxUsers.set([user]);
};

export const inboxMessages = writable<Message[] >([])

export const appendNewMessage = (message: Message) => {
	return inboxMessages.update(prev => [...prev, message]);
}

export const setInboxMessages = (messages: Message[]) => {
	return inboxMessages.set(messages);
}
