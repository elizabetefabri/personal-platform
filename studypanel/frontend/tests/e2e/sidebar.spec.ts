import { test, expect } from '@playwright/test';

test.describe('Sidebar', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/dashboard');
  });

  test('shows 3 main navigation items', async ({ page }) => {
    const navItems = page.locator('.sidebarNav .navItem');
    await expect(navItems).toHaveCount(3);
  });

  test('shows Dashboard nav item', async ({ page }) => {
    await expect(page.locator('.navItem').filter({ hasText: 'Dashboard' })).toBeVisible();
  });

  test('shows Estudos e Labs nav item', async ({ page }) => {
    await expect(page.locator('.navItem').filter({ hasText: 'Estudos e Labs' })).toBeVisible();
  });

  test('shows Projetos nav item', async ({ page }) => {
    await expect(page.locator('.navItem').filter({ hasText: 'Projetos' })).toBeVisible();
  });

  test('highlights Dashboard when on /dashboard', async ({ page }) => {
    const dashboardItem = page.locator('.navItem--active').filter({ hasText: 'Dashboard' });
    await expect(dashboardItem).toBeVisible();
  });

  test('highlights Estudos e Labs when on /backend', async ({ page }) => {
    await page.goto('/backend');
    const item = page.locator('.navItem--active').filter({ hasText: 'Estudos e Labs' });
    await expect(item).toBeVisible();
  });

  test('highlights Projetos when on /projetos/pessoais', async ({ page }) => {
    await page.goto('/projetos/pessoais');
    const item = page.locator('.navItem--active').filter({ hasText: 'Projetos' });
    await expect(item).toBeVisible();
  });

  test('sidebar toggle button is visible', async ({ page }) => {
    await expect(page.locator('.toggleButton')).toBeVisible();
  });

  test('sidebar collapses when toggle is clicked', async ({ page }) => {
    await page.locator('.toggleButton').click();
    await expect(page.locator('.sidebar')).toHaveClass(/sidebar--collapsed/);
  });
});
