import requests
import RPi.GPIO as GPIO
import time


GPIO.setmode(GPIO.BCM)
GPIO.setwarnings(False)
GPIO.setup(18,GPIO.OUT)
GPIO.setup(15,GPIO.OUT)
#sets up the basic LED functionality
#pin 18 sends pwr for green light, 15 sends for red light

abc=0

while abc < 1000:
    # not sure if this while loop actually works as intended, but it works
    url='http://www.nleiva.com:8081'
    # URL for web server
    state = requests.get(url)
    # HTTP get on the web server
    state_text = state.text
    # pulls the actual text from url
    if state_text == 'on':
        print "Green LED on"
        GPIO.output(18,GPIO.HIGH)
        #this command tells the Pi to send power on pin 18

        time.sleep(10)
        print "LED off"
        # GPIO.output(18,GPIO.LOW)
        abc += 1
    else :
        GPIO.output(18,GPIO.LOW)
        #this commnds tells the Pi to stop sending power on pin 18


# below code isn't currently in use, controls the red pin on the LED
# print "Red LED on"
# GPIO.output(15,GPIO.HIGH)
# time.sleep(1)
# print "LED off"
# GPIO.output(15,GPIO.LOW)
