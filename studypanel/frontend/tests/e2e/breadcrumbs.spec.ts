import { test, expect } from '@playwright/test';

test.describe('Breadcrumbs', () => {
  test('shows no breadcrumbs on /dashboard', async ({ page }) => {
    await page.goto('/dashboard');
    await expect(page.locator('nav[aria-label="Navegação estrutural"]')).not.toBeVisible();
  });

  test('shows breadcrumbs on /estudos-labs', async ({ page }) => {
    await page.goto('/estudos-labs');
    const nav = page.locator('nav[aria-label="Navegação estrutural"]');
    await expect(nav).toBeVisible();
    await expect(nav).toContainText('Dashboard');
    await expect(nav).toContainText('Estudos e Labs');
  });

  test('shows correct breadcrumbs on /backend', async ({ page }) => {
    await page.goto('/backend');
    const nav = page.locator('nav[aria-label="Navegação estrutural"]');
    await expect(nav).toContainText('Dashboard');
    await expect(nav).toContainText('Estudos e Labs');
    await expect(nav).toContainText('Backend');
  });

  test('shows correct breadcrumbs on /projetos', async ({ page }) => {
    await page.goto('/projetos');
    const nav = page.locator('nav[aria-label="Navegação estrutural"]');
    await expect(nav).toContainText('Dashboard');
    await expect(nav).toContainText('Projetos');
  });

  test('shows correct breadcrumbs on /projetos/pessoais', async ({ page }) => {
    await page.goto('/projetos/pessoais');
    const nav = page.locator('nav[aria-label="Navegação estrutural"]');
    await expect(nav).toContainText('Dashboard');
    await expect(nav).toContainText('Projetos');
    await expect(nav).toContainText('Projetos Pessoais');
  });

  test('shows correct breadcrumbs on /projetos/profissionais', async ({ page }) => {
    await page.goto('/projetos/profissionais');
    const nav = page.locator('nav[aria-label="Navegação estrutural"]');
    await expect(nav).toContainText('Dashboard');
    await expect(nav).toContainText('Projetos');
    await expect(nav).toContainText('Projetos Profissionais');
  });

  test('Dashboard breadcrumb is a link', async ({ page }) => {
    await page.goto('/backend');
    const dashboardLink = page.locator('nav a').filter({ hasText: 'Dashboard' });
    await expect(dashboardLink).toBeVisible();
  });

  test('breadcrumb Dashboard link navigates to /dashboard', async ({ page }) => {
    await page.goto('/backend');
    await page.locator('nav a').filter({ hasText: 'Dashboard' }).click();
    await expect(page).toHaveURL(/\/dashboard/);
  });
});
