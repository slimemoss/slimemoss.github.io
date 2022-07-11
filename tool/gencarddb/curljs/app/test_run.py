import base64url
import pytest
import requests

from run import app

SAMPLE_URL = 'https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=1&sess=1&pid=1101000&rp=99999&request_locale=ja'

def test_search():
    client = app.test_client()
    encoded_url = base64url.base64_encode(SAMPLE_URL)
    resp = client.get('/?q=' + encoded_url)
    jsed_data = resp.get_data().decode('UTF-8')

    resp = requests.get(SAMPLE_URL)
    nojs_data = resp.text

    assert len(jsed_data) > len(nojs_data) + 100
