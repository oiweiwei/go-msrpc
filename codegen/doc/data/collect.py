from selenium import webdriver

from selenium.webdriver.common.by import By
from selenium.webdriver.support.wait import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


browser = webdriver.Chrome()

selector = (By.CSS_SELECTOR, 'li:nth-child(1) > blockquote > cite')

for l in open('types.in.tsv', 'r').read().split('\n'):
    l = l.split('\t')
    if len(l) < 2:
        continue
    if len(l) < 3:
        browser.get(f'https://learn.microsoft.com/en-us/search/?terms={l[1]}')
        WebDriverWait(browser, 10).until(EC.presence_of_element_located(selector))
        l.append(
            browser.find_element(By.CSS_SELECTOR, 'li:nth-child(1) > blockquote > cite').text.split('/')[-1]
        )
    print('\t'.join(l))
