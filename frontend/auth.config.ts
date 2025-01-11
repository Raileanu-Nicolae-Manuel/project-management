import GitHub from '@auth/core/providers/github';
import { defineConfig } from 'auth-astro';

export default defineConfig({
  providers: [
    GitHub({
      clientId: import.meta.env.GITHUB_CLIENT_ID,
      clientSecret: import.meta.env.GITHUB_CLIENT_SECRET,
      authorization: {
        params: {
          scope: 'read:user repo'
        }
      }
    }),
  ],
  callbacks: {
    async session({ session, token }) {
      if (token) {
        session.accessToken = token.accessToken;
      }
      return session;
    },
    async jwt({ token, account }) {
      if (account) {
        token.accessToken = account.access_token;
      }
      return token;
    }
  },
});