# coding: utf-8
from turtle import position
import pytest
import responses
import kiteconnect as kc
import kiteconnect.exceptions as ex

import utils

def test_setAccessToken(kc):
    """check for token exception when invalid token is set."""
    kc.root = "https://api.kite.trade"
    kc.set_access_token("invalid_token")
    with pytest.raises(ex.TokenException):
        kc.positions()

@responses.activate
def test_positions(kc):
    """ Test Positions """
    responses.add(
        responses.GET,
        "{0}{1}".format(kc.root, kc._routes["portfolio.positions"]),
        body = utils.get_response("portfolio.positions"),
        content_type="application/json"
    )
    positions = kc.positions()
    assert type(positions) == dict
    assert "day" in positions
    assert "net" in positions

