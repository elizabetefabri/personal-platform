import { test, expect } from '@playwright/test';

test.describe('Responsive Grid', () => {
  test('shows 1 column on mobile (375px)', async ({ page }) => {
    await page.setViewportSize({ width: 375, height: 812 });
    await page.goto('/estudos-labs');
    const grid = page.locator('.studyGrid');
    const columns = await grid.evaluate((el) =>
      getComputedStyle(el).getPropertyValue('grid-template-columns'),
    );
    expect(columns.trim().split(' ').length).toBe(1);
  });

  test('shows 2 columns on small tablet (500px)', async ({ page }) => {
    await page.setViewportSize({ width: 500, height: 768 });
    await page.goto('/estudos-labs');
    const grid = page.locator('.studyGrid');
    const columns = await grid.evaluate((el) =>
      getComputedStyle(el).getPropertyValue('grid-template-columns'),
    );
    expect(columns.trim().split(' ').length).toBe(2);
  });

  test('shows 3 columns on tablet (768px)', async ({ page }) => {
    await page.setViewportSize({ width: 768, height: 1024 });
    await page.goto('/estudos-labs');
    const grid = page.locator('.studyGrid');
    const columns = await grid.evaluate((el) =>
      getComputedStyle(el).getPropertyValue('grid-template-columns'),
    );
    expect(columns.trim().split(' ').length).toBe(3);
  });

  test('shows 4 columns on desktop (1024px)', async ({ page }) => {
    await page.setViewportSize({ width: 1024, height: 768 });
    await page.goto('/estudos-labs');
    const grid = page.locator('.studyGrid');
    const columns = await grid.evaluate((el) =>
      getComputedStyle(el).getPropertyValue('grid-template-columns'),
    );
    expect(columns.trim().split(' ').length).toBe(4);
  });

  test('shows 6 columns on wide screen (1440px)', async ({ page }) => {
    await page.setViewportSize({ width: 1440, height: 900 });
    await page.goto('/estudos-labs');
    const grid = page.locator('.studyGrid');
    const columns = await grid.evaluate((el) =>
      getComputedStyle(el).getPropertyValue('grid-template-columns'),
    );
    expect(columns.trim().split(' ').length).toBe(6);
  });

  test('project grid is visible on /projetos/pessoais', async ({ page }) => {
    await page.setViewportSize({ width: 1024, height: 768 });
    await page.goto('/projetos/pessoais');
    await expect(page.locator('.projectGrid')).toBeVisible();
  });
});
