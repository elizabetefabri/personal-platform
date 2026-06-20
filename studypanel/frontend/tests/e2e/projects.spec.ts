import { test, expect } from '@playwright/test';

test.describe('Projects Pages', () => {
  test('Projetos landing shows 2 cards', async ({ page }) => {
    await page.goto('/projetos');
    const cards = page.locator('.studyCard');
    await expect(cards).toHaveCount(2);
  });

  test('Projetos landing has Projetos Pessoais card', async ({ page }) => {
    await page.goto('/projetos');
    await expect(page.locator('.studyCard').filter({ hasText: 'Projetos Pessoais' })).toBeVisible();
  });

  test('Projetos landing has Projetos Profissionais card', async ({ page }) => {
    await page.goto('/projetos');
    await expect(page.locator('.studyCard').filter({ hasText: 'Projetos Profissionais' })).toBeVisible();
  });

  test('Projetos Pessoais page shows project cards', async ({ page }) => {
    await page.goto('/projetos/pessoais');
    const cards = page.locator('.projectCard');
    await expect(cards.first()).toBeVisible();
  });

  test('Projetos Pessoais cards have repo or deploy buttons', async ({ page }) => {
    await page.goto('/projetos/pessoais');
    const firstCard = page.locator('.projectCard').first();
    const actionButton = firstCard.locator('.projectCard__btn');
    await expect(actionButton.first()).toBeVisible();
  });

  test('Projetos Profissionais page shows Rollout Service card', async ({ page }) => {
    await page.goto('/projetos/profissionais');
    await expect(page.locator('.projectCard').filter({ hasText: 'Rollout Service' })).toBeVisible();
  });

  test('back button is visible on /projetos/pessoais', async ({ page }) => {
    await page.goto('/projetos/pessoais');
    await expect(page.locator('app-back-button button')).toBeVisible();
  });

  test('back button navigates to /projetos from /projetos/pessoais', async ({ page }) => {
    await page.goto('/projetos/pessoais');
    await page.locator('app-back-button button').click();
    await expect(page).toHaveURL(/\/projetos$/);
  });
});
