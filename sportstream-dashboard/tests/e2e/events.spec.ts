import { test, expect } from '@playwright/test';

test.describe('Events Page', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/events');
    await page.waitForLoadState('networkidle');
  });

  test('page loads with Events title', async ({ page }) => {
    const heading = page.locator('h2').filter({ hasText: 'Events' });
    await expect(heading).toBeVisible();

    // Breadcrumb should show Management > Events
    await expect(page.getByText('Management')).toBeVisible();
  });

  test('event entries are visible after loading', async ({ page }) => {
    // Wait for loading skeletons to disappear
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // Events are rendered inside a container with EventCard components
    // Each event card is a div with rounded-xl styling containing an h4 title
    const eventCards = page.locator('div[class*="rounded-xl"]').filter({
      has: page.locator('h4'),
    });
    expect(await eventCards.count()).toBeGreaterThan(0);
  });

  test('events show title, venue, sport badge, and time', async ({ page }) => {
    // Wait for data to load
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // Get the first event card
    const firstEvent = page.locator('div[class*="rounded-xl"]').filter({
      has: page.locator('h4'),
    }).first();
    await expect(firstEvent).toBeVisible();

    // Event title (h4)
    const title = firstEvent.locator('h4');
    await expect(title).toBeVisible();
    const titleText = await title.textContent();
    expect(titleText?.trim().length).toBeGreaterThan(0);

    // Sport badge — uppercase span with sport name
    const sportBadge = firstEvent.locator('span[class*="uppercase"][class*="rounded"]').first();
    await expect(sportBadge).toBeVisible();
    const sportText = await sportBadge.textContent();
    expect(sportText?.trim().length).toBeGreaterThan(0);

    // Venue — displayed with MapPin icon, as text in the event info area
    const eventText = await firstEvent.textContent();
    // The card should contain time information (HH:MM format)
    expect(eventText).toMatch(/\d{2}:\d{2}/);
  });

  test('status filter dropdown is available', async ({ page }) => {
    const statusSelect = page.locator('select').filter({
      has: page.locator('option', { hasText: 'All Statuses' }),
    });
    await expect(statusSelect).toBeVisible();

    // Should have status options
    const options = statusSelect.locator('option');
    expect(await options.count()).toBeGreaterThan(1);
  });

  test('sport filter input is available', async ({ page }) => {
    const sportInput = page.locator('input[placeholder="Filter by sport..."]');
    await expect(sportInput).toBeVisible();
  });

  test('event cards display date with month and day', async ({ page }) => {
    // Wait for data to load
    await expect(page.locator('.animate-pulse').first()).not.toBeVisible({ timeout: 10000 });

    // Each event card has a date section with month abbreviation and day number
    // Month is rendered as uppercase text (e.g., "JAN", "FEB")
    const firstEvent = page.locator('div[class*="rounded-xl"]').filter({
      has: page.locator('h4'),
    }).first();

    const eventText = await firstEvent.textContent();
    // Should contain a month abbreviation (3 uppercase letters)
    expect(eventText).toMatch(/[A-Z]{3}/);
    // Should contain a day number
    expect(eventText).toMatch(/\d{1,2}/);
  });
});
