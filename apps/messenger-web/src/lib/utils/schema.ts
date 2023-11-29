import * as z from 'zod';
import { validator } from './validator';

export const loginSchema = z.object({
	email: z.string().email('Email must be valid'),
	password: z.string().regex(validator.password, 'Password must be valid')
});

export const signUpSchema = z
	.object({
		email: z.string().email('Email must be valid'),
		firstName: z.string(),
		lastName: z.string(),
		password: z.string().regex(validator.password, 'Password must be valid'),
		confirmPassword: z.string(),
		gender: z.string(),
		avatar: z.string()
	})
	.superRefine(({ confirmPassword, password }, context) => {
		if (confirmPassword !== password) {
			context.addIssue({
				code: 'custom',
				message: "Confirm passwords doesn't match"
			});
		}
	});

export const sendMessageSchema = z.object({
	content: z.string().optional()
});
