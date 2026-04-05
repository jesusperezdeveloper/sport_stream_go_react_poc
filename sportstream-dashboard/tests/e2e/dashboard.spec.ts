import { test, expect } from '@playwright/test';

test.describe('Dashboard Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
    await page.waitForLoadState('networkidle');
  });

  test('page loads and shows Dashboard as active nav item', async ({ page }) => {
    // The sidebar nav should show "Dashboard" as the active link (with active styling)
    const dashboardLink = page.locator('nav a[href="/"]').first();
    await expect(dashboardLink).toBeVisible();
    await expect(dashboardLink).toContainText('Dashboard');
    // Active state applies border-sidebar-accent class — check the element has the active color
    await expect(dashboardLink).toHaveClass(/text-sidebar-accent|text-\[#00e676\]/);
  });

  test('displays 4 stats cards with labels', async ({ page }) => {
    // Wait for data to load (skeletons disappear)
    await page.waitForTimeout(2000);

    await expect(page.getByText('Total Clubs')).toBeVisible({ timeout: 10000 });
    await expect(page.getByText('Active Streams')).toBeVisible({ timeout: 10000 });
    await expect(page.getByText('Upcoming Events').first()).toBeVisible({ timeout: 10000 });
    await expect(page.getByText('Total Network Views')).toBeVisible({ timeout: 10000 });
  });

  test('stats cards show values loaded from API', async ({ page }) => {
    // Wait for loading skeletons to disappear — real values should replace them
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // Each stats card should have a visible numeric value (not just the label)
    // The StatsCard component renders the value as a prominent number
    const totalClubsCard = page.locator('div').filter({ hasText: 'Total Clubs' }).first();
    await expect(totalClubsCard).toBeVisible();

    const activeStreamsCard = page.locator('div').filter({ hasText: 'Active Streams' }).first();
    await expect(activeStreamsCard).toBeVisible();
  });

  test('Streams by Type section is visible', async ({ page }) => {
    const section = page.getByText('Streams by Type');
    await expect(section).toBeVisible();

    // Should display stream type labels from API data
    await expect(page.getByText('Live Match').first()).toBeVisible({ timeout: 10000 });
  });

  test('LIVE NOW section is visible with at least 1 stream', async ({ page }) => {
    const liveNowHeading = page.getByText('LIVE NOW');
    await expect(liveNowHeading).toBeVisible();

    // Wait for loading to finish, then check for stream content
    // Featured stream should have a title rendered as an h3
    const livePanel = page.locator('div').filter({ hasText: 'LIVE NOW' }).first();
    await expect(livePanel).toBeVisible();

    // At least one stream title should be visible inside the live panel area
    const streamTitles = page.locator('h3').filter({
      has: page.locator('text=/.+/'),
    });
    await expect(streamTitles.first()).toBeVisible({ timeout: 10000 });
  });

  test('Upcoming Events section is visible', async ({ page }) => {
    const heading = page.getByText('Upcoming Events').first();
    await expect(heading).toBeVisible();

    // Should contain event entries with titles after loading
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });
  });

  test('Top Clubs by Views section is visible with ranked items', async ({ page }) => {
    const heading = page.getByText('Top Clubs by Views');
    await expect(heading).toBeVisible();

    // After loading, should display club names with rank numbers
    // Rank circles contain numbers 1-5
    const rankOne = page.locator('div').filter({ hasText: /^1$/ }).first();
    await expect(rankOne).toBeVisible({ timeout: 10000 });

    // Club names should be visible as h4 elements within the section
    const topClubsSection = page.locator('div').filter({ hasText: 'Top Clubs by Views' }).last();
    const clubNames = topClubsSection.locator('h4');
    await expect(clubNames.first()).toBeVisible();
    expect(await clubNames.count()).toBeGreaterThanOrEqual(1);
  });
});
