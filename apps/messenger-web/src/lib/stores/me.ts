import type { User } from '$lib/types';
import {  writable } from 'svelte/store';

export const me = writable<User | null>(null);

