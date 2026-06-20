import { test, expect } from '@playwright/test';

test.describe('SEO', () => {
  test('has correct page title', async ({ page }) => {
    await page.goto('/dashboard');
    await expect(page).toHaveTitle(/StudyPanel/);
  });

  test('has meta description', async ({ page }) => {
    await page.goto('/dashboard');
    const meta = page.locator('meta[name="description"]');
    await expect(meta).toHaveAttribute('content', /.+/);
  });

  test('has meta robots tag', async ({ page }) => {
    await page.goto('/dashboard');
    const robots = page.locator('meta[name="robots"]');
    await expect(robots).toHaveAttribute('content', /index/);
  });

  test('has Open Graph title', async ({ page }) => {
    await page.goto('/dashboard');
    const ogTitle = page.locator('meta[property="og:title"]');
    await expect(ogTitle).toHaveAttribute('content', /StudyPanel/);
  });

  test('has Open Graph type', async ({ page }) => {
    await page.goto('/dashboard');
    const ogType = page.locator('meta[property="og:type"]');
    await expect(ogType).toHaveAttribute('content', /.+/);
  });

  test('has lang attribute set to pt-BR', async ({ page }) => {
    await page.goto('/dashboard');
    const html = page.locator('html');
    await expect(html).toHaveAttribute('lang', 'pt-BR');
  });

  test('robots.txt is accessible', async ({ page }) => {
    const response = await page.goto('/robots.txt');
    expect(response?.status()).toBe(200);
  });

  test('sitemap.xml is accessible', async ({ page }) => {
    const response = await page.goto('/sitemap.xml');
    expect(response?.status()).toBe(200);
  });
});
