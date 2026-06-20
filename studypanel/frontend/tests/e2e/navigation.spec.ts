import { test, expect } from '@playwright/test';

test.describe('Navigation', () => {
  test('redirects root to /dashboard', async ({ page }) => {
    await page.goto('/');
    await expect(page).toHaveURL(/\/dashboard/);
  });

  test('navigates to Estudos e Labs page', async ({ page }) => {
    await page.goto('/estudos-labs');
    await expect(page.locator('h1')).toBeVisible();
  });

  test('navigates to Projetos page', async ({ page }) => {
    await page.goto('/projetos');
    await expect(page.locator('h1')).toBeVisible();
  });

  test('navigates to Projetos Pessoais page', async ({ page }) => {
    await page.goto('/projetos/pessoais');
    await expect(page.locator('h1')).toBeVisible();
  });

  test('navigates to Projetos Profissionais page', async ({ page }) => {
    await page.goto('/projetos/profissionais');
    await expect(page.locator('h1')).toBeVisible();
  });

  test('navigates to Backend study section', async ({ page }) => {
    await page.goto('/backend');
    await expect(page.locator('h1')).toBeVisible();
  });

  test('navigates to DevOps study section', async ({ page }) => {
    await page.goto('/devops');
    await expect(page.locator('h1')).toBeVisible();
  });

  test('navigates to Cloud study section', async ({ page }) => {
    await page.goto('/cloud');
    await expect(page.locator('h1')).toBeVisible();
  });

  test('navigates to Rollout Service page', async ({ page }) => {
    await page.goto('/rollout-service');
    await expect(page.locator('h1')).toBeVisible();
  });
});
