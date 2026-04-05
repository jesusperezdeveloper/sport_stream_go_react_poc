import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests/e2e',
  timeout: 30000,
  retries: 1,
  use: {
    baseURL: 'http://localhost:5173',
    screenshot: 'on',
    trace: 'on',
  },
  webServer: {
    command: 'npm run dev',
    port: 5173,
    reuseExistingServer: true,
    timeout: 30000,
    env: {
      VITE_API_URL: 'http://localhost:8080',
    },
  },
  projects: [
    { name: 'chromium', use: { browserName: 'chromium' } },
  ],
});
