import time
import urllib
import base64url

from flask import Flask, request
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

app = Flask(__name__)

@app.route('/', methods=['GET'])
def search():
    if 'q' not in request.args:
        return '', 400
    q = request.args.get('q', type=str)
    q = base64url.base64_decode(q)
    return jsed_html(q), 200


def getWebDriver():
    options = Options()
    options.add_argument('--headless')
    options.add_argument('--no-sandbox')
    return webdriver.Chrome(options=options)


driver = getWebDriver()
def jsed_html(url):
    """
    urlからjavascriptを実行したhtmlを返す
    """
    print(url)
    driver.get(url)
    time.sleep(0.8)
    html = driver.page_source
    return html
