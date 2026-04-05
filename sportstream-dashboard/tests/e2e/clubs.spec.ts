import { test, expect } from '@playwright/test';

test.describe('Clubs Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/clubs');
    await page.waitForLoadState('networkidle');
  });

  test('page loads with Clubs title', async ({ page }) => {
    const heading = page.locator('h2').filter({ hasText: 'Clubs' });
    await expect(heading).toBeVisible();

    // Breadcrumb should show Management > Clubs
    await expect(page.getByText('Management', { exact: true })).toBeVisible();
  });

  test('club cards are visible in a grid', async ({ page }) => {
    // Wait for loading skeletons to disappear
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // The grid container uses CSS grid with sm:grid-cols-2 lg:grid-cols-3
    const grid = page.locator('div.grid').filter({
      has: page.locator('div[class*="rounded-2xl"]'),
    });
    await expect(grid.first()).toBeVisible();

    // Club cards should be present
    const cards = page.locator('div[class*="rounded-2xl"][class*="shadow-sm"]').filter({
      has: page.locator('h3'),
    });
    expect(await cards.count()).toBeGreaterThan(0);
  });

  test('each card shows club name, league, country, and sport badge', async ({ page }) => {
    // Wait for data to load
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // Get the first club card
    const firstCard = page.locator('div[class*="rounded-2xl"][class*="shadow-sm"]').filter({
      has: page.locator('h3'),
    }).first();
    await expect(firstCard).toBeVisible();

    // Club name (h3)
    const clubName = firstCard.locator('h3');
    await expect(clubName).toBeVisible();
    const nameText = await clubName.textContent();
    expect(nameText?.trim().length).toBeGreaterThan(0);

    // League — rendered as a p element with text-xs class below the name
    const league = firstCard.locator('p').filter({ hasText: /.+/ }).first();
    await expect(league).toBeVisible();

    // Country — displayed as text in the bottom section
    const cardText = await firstCard.textContent();
    // The card should contain a pipe separator between country and sport
    expect(cardText).toContain('|');

    // Sport badge — uppercase rounded-full badge
    const sportBadge = firstCard.locator('span[class*="rounded-full"][class*="uppercase"]').first();
    await expect(sportBadge).toBeVisible();
    const sportText = await sportBadge.textContent();
    expect(sportText?.trim().length).toBeGreaterThan(0);
  });

  test('at least 5 clubs are displayed from seed data', async ({ page }) => {
    // Wait for data to load
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    const cards = page.locator('div[class*="rounded-2xl"][class*="shadow-sm"]').filter({
      has: page.locator('h3'),
    });
    expect(await cards.count()).toBeGreaterThanOrEqual(5);
  });
});
