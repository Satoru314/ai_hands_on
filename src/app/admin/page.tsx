'use client';

import { useEffect, useState } from 'react';

interface Submission {
  id: number;
  name: string;
  email: string;
  message: string;
  createdAt: string;
}

export default function AdminPage() {
  const [data, setData] = useState<Submission[]>([]);

  useEffect(() => {
    fetch('/api/submissions')
      .then((res) => res.json())
      .then(setData);
  }, []);

  return (
    <main className="max-w-md mx-auto p-4">
      <h1 className="text-xl mb-4">Submissions</h1>
      <ul className="space-y-2">
        {data.map((s) => (
          <li key={s.id} className="border p-2">
            <p className="font-bold">
              {s.name} ({s.email})
            </p>
            <p>{s.message}</p>
            <p className="text-xs text-gray-500">
              {new Date(s.createdAt).toLocaleString()}
            </p>
          </li>
        ))}
      </ul>
    </main>
  );
}
