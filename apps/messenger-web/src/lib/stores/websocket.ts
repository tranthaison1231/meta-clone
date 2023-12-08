import { writable } from 'svelte/store';

export const ws = writable<WebSocket>(new WebSocket(import.meta.env.VITE_WS_URL));
