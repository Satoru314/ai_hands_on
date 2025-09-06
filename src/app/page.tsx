'use client';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { submissionSchema, SubmissionInput } from '@/lib/validation';

export default function Home() {
  const { register, handleSubmit, reset, formState: { errors } } = useForm<SubmissionInput>({
    resolver: zodResolver(submissionSchema),
  });

  const onSubmit = async (data: SubmissionInput) => {
    await fetch('/api/submissions', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    reset();
  };

  return (
    <main className="max-w-md mx-auto p-4">
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-2">
        <input {...register('name')} placeholder="Name" className="border p-2" />
        {errors.name && <p className="text-red-500 text-sm">{errors.name.message}</p>}
        <input {...register('email')} placeholder="Email" className="border p-2" />
        {errors.email && <p className="text-red-500 text-sm">{errors.email.message}</p>}
        <textarea {...register('message')} placeholder="Message" className="border p-2" />
        {errors.message && <p className="text-red-500 text-sm">{errors.message.message}</p>}
        <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">Submit</button>
      </form>
    </main>
  );
}
