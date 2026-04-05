import { test, expect } from '@playwright/test';

test.describe('Navigation', () => {
  test('sidebar links navigate between pages', async ({ page }) => {
    await page.goto('/');
    await page.waitForLoadState('networkidle');

    // Click on Streams in the sidebar
    const streamsLink = page.locator('aside nav a[href="/streams"]');
    await expect(streamsLink).toBeVisible();
    await streamsLink.click();
    await page.waitForLoadState('networkidle');

    // Should be on the Streams page
    await expect(page).toHaveURL(/\/streams/);
    const streamsHeading = page.locator('h2').filter({ hasText: 'Streams' });
    await expect(streamsHeading).toBeVisible();
  });

  test('Dashboard to Streams to Clubs to Events navigation works', async ({ page }) => {
    // Start at Dashboard
    await page.goto('/');
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL('/');

    // Navigate to Streams
    await page.locator('aside nav a[href="/streams"]').click();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/streams/);
    await expect(page.locator('h2').filter({ hasText: 'Streams' })).toBeVisible();

    // Navigate to Clubs
    await page.locator('aside nav a[href="/clubs"]').click();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/clubs/);
    await expect(page.locator('h2').filter({ hasText: 'Clubs' })).toBeVisible();

    // Navigate to Events
    await page.locator('aside nav a[href="/events"]').click();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/events/);
    await expect(page.locator('h2').filter({ hasText: 'Events' })).toBeVisible();

    // Navigate back to Dashboard
    await page.locator('aside nav a[href="/"]').click();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL('/');
  });

  test('active nav item is highlighted', async ({ page }) => {
    // On Dashboard page, the Dashboard link should have active styling
    await page.goto('/');
    await page.waitForLoadState('networkidle');

    const dashboardLink = page.locator('aside nav a[href="/"]');
    await expect(dashboardLink).toHaveClass(/text-sidebar-accent/);

    // Streams link should NOT have active styling
    const streamsLink = page.locator('aside nav a[href="/streams"]');
    await expect(streamsLink).not.toHaveClass(/text-sidebar-accent/);

    // Navigate to Streams
    await streamsLink.click();
    await page.waitForLoadState('networkidle');

    // Now Streams should be active
    await expect(page.locator('aside nav a[href="/streams"]')).toHaveClass(/text-sidebar-accent/);
    // Dashboard should no longer be active
    await expect(page.locator('aside nav a[href="/"]')).not.toHaveClass(/text-sidebar-accent/);
  });

  test('browser back and forward navigation works', async ({ page }) => {
    await page.goto('/');
    await page.waitForLoadState('networkidle');

    // Navigate to Streams
    await page.locator('aside nav a[href="/streams"]').click();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/streams/);

    // Navigate to Clubs
    await page.locator('aside nav a[href="/clubs"]').click();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/clubs/);

    // Go back — should return to Streams
    await page.goBack();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/streams/);
    await expect(page.locator('h2').filter({ hasText: 'Streams' })).toBeVisible();

    // Go back again — should return to Dashboard
    await page.goBack();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL('/');

    // Go forward — should go to Streams
    await page.goForward();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/streams/);
    await expect(page.locator('h2').filter({ hasText: 'Streams' })).toBeVisible();

    // Go forward again — should go to Clubs
    await page.goForward();
    await page.waitForLoadState('networkidle');
    await expect(page).toHaveURL(/\/clubs/);
    await expect(page.locator('h2').filter({ hasText: 'Clubs' })).toBeVisible();
  });
});
