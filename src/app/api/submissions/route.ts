import { NextResponse } from 'next/server';
import { prisma } from '@/lib/prisma';
import { submissionSchema } from '@/lib/validation';
import { ZodError } from 'zod';

export async function GET() {
  const submissions = await prisma.submission.findMany({
    orderBy: { createdAt: 'desc' },
    take: 100,
  });
  return NextResponse.json(submissions);
}

export async function POST(req: Request) {
  try {
    const json = await req.json();
    const data = submissionSchema.parse(json);
    const created = await prisma.submission.create({ data });
    return NextResponse.json({ id: created.id, createdAt: created.createdAt });
  } catch (err) {
    if (err instanceof ZodError) {
      return NextResponse.json({ error: err.issues }, { status: 400 });
    }
    return NextResponse.json({ error: 'Server error' }, { status: 500 });
  }
}
