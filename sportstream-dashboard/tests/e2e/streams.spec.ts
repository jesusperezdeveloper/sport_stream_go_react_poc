import { test, expect } from '@playwright/test';

test.describe('Streams Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/streams');
    await page.waitForLoadState('networkidle');
  });

  test('page loads with Streams title', async ({ page }) => {
    const heading = page.locator('h2').filter({ hasText: 'Streams' });
    await expect(heading).toBeVisible();

    // Breadcrumb should show Management > Streams
    await expect(page.getByText('Management', { exact: true })).toBeVisible();
  });

  test('stream entries are visible in the table on desktop', async ({ page }) => {
    // Wait for loading skeleton to disappear
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // Desktop table should have column headers
    const table = page.locator('table');
    await expect(table).toBeVisible();

    await expect(page.locator('th').filter({ hasText: 'Title' })).toBeVisible();
    await expect(page.locator('th').filter({ hasText: 'Club' })).toBeVisible();
    await expect(page.locator('th').filter({ hasText: 'Type' })).toBeVisible();
    await expect(page.locator('th').filter({ hasText: 'Status' })).toBeVisible();
    await expect(page.locator('th').filter({ hasText: 'Views' })).toBeVisible();
    await expect(page.locator('th').filter({ hasText: 'Duration' })).toBeVisible();
    await expect(page.locator('th').filter({ hasText: 'Actions' })).toBeVisible();

    // Table body should have rows
    const rows = table.locator('tbody tr');
    expect(await rows.count()).toBeGreaterThan(0);
  });

  test('filter dropdowns are visible', async ({ page }) => {
    // Status filter dropdown
    const statusSelect = page.locator('select').filter({ has: page.locator('option', { hasText: 'All Statuses' }) });
    await expect(statusSelect).toBeVisible();

    // Type filter dropdown
    const typeSelect = page.locator('select').filter({ has: page.locator('option', { hasText: 'All Types' }) });
    await expect(typeSelect).toBeVisible();
  });

  test('stream entries show title, club name, type badge, and status badge', async ({ page }) => {
    // Wait for data to load
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    const firstRow = page.locator('table tbody tr').first();
    await expect(firstRow).toBeVisible();

    // Title cell — contains a span with stream title text
    const titleCell = firstRow.locator('td').nth(0);
    await expect(titleCell.locator('span')).toBeVisible();

    // Club cell — second column
    const clubCell = firstRow.locator('td').nth(1);
    const clubText = await clubCell.textContent();
    expect(clubText?.trim().length).toBeGreaterThan(0);

    // Type badge — should contain a type like LIVE, VOD, HIGHLIGHT, or BTS
    const typeCell = firstRow.locator('td').nth(2);
    const typeBadge = typeCell.locator('span');
    await expect(typeBadge).toBeVisible();
    const typeText = await typeBadge.textContent();
    expect(typeText?.trim()).toMatch(/live|vod|highlight|bts/i);

    // Status badge — fourth column (uses div or span for badge)
    const statusCell = firstRow.locator('td').nth(3);
    await expect(statusCell).not.toBeEmpty();
  });

  test('New Stream button is visible', async ({ page }) => {
    const newStreamButton = page.locator('button').filter({ hasText: /New Stream/ });
    await expect(newStreamButton).toBeVisible();
  });

  test('Watch button is visible for non-scheduled streams', async ({ page }) => {
    // Wait for data to load
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // At least one Watch button should exist in the actions column
    const watchButtons = page.locator('table button').filter({ hasText: 'Watch' });
    expect(await watchButtons.count()).toBeGreaterThan(0);
    await expect(watchButtons.first()).toBeVisible();
  });
});
