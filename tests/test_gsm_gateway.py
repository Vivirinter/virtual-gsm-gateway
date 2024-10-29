import requests

def test_send_sms():
    url = "http://localhost:8080/send-sms"
    payload = {"from": "12345", "to": "67890", "message": "Hello"}
    response = requests.post(url, json=payload)
    assert response.status_code == 201

def test_get_sms():
    url = "http://localhost:8080/messages"
    response = requests.get(url)
    assert response.status_code == 200
    messages = response.json()
    assert len(messages) > 0

def test_delete_sms():
    url = "http://localhost:8080/delete-messages"
    response = requests.post(url)
    assert response.status_code == 200
    response = requests.get("http://localhost:8080/messages")
    messages = response.json()
    assert len(messages) == 0

def test_send_ussd():
    url = "http://localhost:8080/send-ussd"
    payload = {"from": "12345", "code": "*123#"}
    response = requests.post(url, json=payload)
    assert response.status_code == 200
    ussd_response = response.json()
    assert ussd_response["response"] == "USSD response for code: *123#"

def test_get_ussd():
    url = "http://localhost:8080/ussd-requests"
    response = requests.get(url)
    assert response.status_code == 200
    ussd_requests = response.json()
    assert len(ussd_requests) > 0

def test_delete_ussd():
    url = "http://localhost:8080/delete-ussd-requests"
    response = requests.post(url)
    assert response.status_code == 200
    response = requests.get("http://localhost:8080/ussd-requests")
    ussd_requests = response.json()
    assert len(ussd_requests) == 0