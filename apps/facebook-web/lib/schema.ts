import * as z from 'zod';
import { validator } from './validator';

export const loginSchema = z.object({
  email: z.string().email('Email must be valid'),
  password: z.string().regex(validator.password, 'Password must be valid'),
});
