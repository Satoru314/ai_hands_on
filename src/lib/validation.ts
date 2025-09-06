import { z } from 'zod';

export const submissionSchema = z.object({
  name: z.string().min(1),
  email: z.string().email(),
  message: z.string().min(1),
});

export type SubmissionInput = z.infer<typeof submissionSchema>;
