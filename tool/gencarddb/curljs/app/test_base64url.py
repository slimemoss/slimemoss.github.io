import base64url

def test_type():
    res = base64url.base64_encode("b")
    assert type("a") == type(res)

    res = base64url.base64_decode("Yg")
    assert type("a") == type(res)

def test_value():
    res = base64url.base64_encode("b")
    assert "Yg" == res

    res = base64url.base64_decode("Yg")
    assert "b" == res
