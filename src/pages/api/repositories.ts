import type { APIRoute } from 'astro';
import { getSession } from 'auth-astro/server';

export const GET: APIRoute = async ({ request }) => {
  const session = await getSession(request);
  
  if (!session?.accessToken) {
    return new Response(JSON.stringify({ error: 'Unauthorized' }), {
      status: 401,
    });
  }

  const response = await fetch('https://api.github.com/user/repos', {
    headers: {
      Authorization: `Bearer ${session.accessToken}`,
      Accept: 'application/vnd.github.v3+json',
    },
  });

  const repositories = await response.json();
  repositories.sort((a, b) => new Date(b.updated_at) - new Date(a.updated_at));

  return new Response(JSON.stringify(repositories), {
    status: 200,
    headers: {
      'Content-Type': 'application/json',
    },
  });
}; 