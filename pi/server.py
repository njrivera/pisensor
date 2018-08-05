import Adafruit_DHT
import time
import os
import sys
import requests

sensor = Adafruit_DHT.DHT22

pin = os.environ['GPIO_PIN']
addr = os.environ['LISTENER_ADDRESS']

if pin == '':
    print("GPIO_PIN not set")
    sys.exit()
if addr == '':
    print('LISTENER_ADDRESS not set')
    sys.exit()

addr += '/testserial/testmodel'

while True:
    time.sleep(5)
    humidity, temperature = Adafruit_DHT.read_retry(sensor, pin)
    temperature = temperature * 1.8 + 32
    requests.post(addr, json={'num':temperature})
    
