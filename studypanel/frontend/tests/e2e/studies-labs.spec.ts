import { test, expect } from '@playwright/test';

test.describe('Estudos e Labs', () => {
  test('shows 9 study cards', async ({ page }) => {
    await page.goto('/estudos-labs');
    const cards = page.locator('.studyCard');
    await expect(cards).toHaveCount(9);
  });

  test('shows Backend card', async ({ page }) => {
    await page.goto('/estudos-labs');
    await expect(page.locator('.studyCard').filter({ hasText: 'Backend' })).toBeVisible();
  });

  test('shows DevOps card', async ({ page }) => {
    await page.goto('/estudos-labs');
    await expect(page.locator('.studyCard').filter({ hasText: 'DevOps' })).toBeVisible();
  });

  test('shows Cloud card', async ({ page }) => {
    await page.goto('/estudos-labs');
    await expect(page.locator('.studyCard').filter({ hasText: 'Cloud' })).toBeVisible();
  });

  test('StudyPanel eyebrow is visible', async ({ page }) => {
    await page.goto('/estudos-labs');
    await expect(page.locator('.studyPage__eyebrow')).toContainText('StudyPanel');
  });

  test('back button is NOT visible on /estudos-labs (root route)', async ({ page }) => {
    await page.goto('/estudos-labs');
    await expect(page.locator('app-back-button button')).not.toBeVisible();
  });

  test('back button IS visible on /backend', async ({ page }) => {
    await page.goto('/backend');
    await expect(page.locator('app-back-button button')).toBeVisible();
  });

  test('clicking Backend card navigates to /backend', async ({ page }) => {
    await page.goto('/estudos-labs');
    await page.locator('.studyCard').filter({ hasText: 'Backend' }).locator('button').click();
    await expect(page).toHaveURL(/\/backend/);
  });

  test('back button on /backend navigates to /estudos-labs', async ({ page }) => {
    await page.goto('/backend');
    await page.locator('app-back-button button').click();
    await expect(page).toHaveURL(/\/estudos-labs/);
  });
});
