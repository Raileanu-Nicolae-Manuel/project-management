// @ts-check
import { defineConfig } from 'astro/config';

import solidJs from '@astrojs/solid-js';

import tailwind from '@astrojs/tailwind';

import auth from 'auth-astro';

import node from '@astrojs/node';

// https://astro.build/config
export default defineConfig({
  output: 'server',
  integrations: [solidJs(), tailwind(), auth()],

  adapter: node({
    mode: 'standalone'
  })
});