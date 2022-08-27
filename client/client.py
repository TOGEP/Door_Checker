import time
import RPi.GPIO as GPIO
import socket

PORT = 5000

GPIO.setmode(GPIO.BCM)
GPIO.setup(18, GPIO.IN, pull_up_down=GPIO.PUD_UP)

client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, socket.IPPROTO_UDP)
client.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)
client.settimeout(0.5)

sw_status = 1

while True:
    try:
        sw_status = GPIO.input(18)
        if sw_status == 0:
            # close
            pass

        else:
            # open
            client.sendto(b'OpenTheDoor', ('<broadcast>', PORT))
            while sw_status == 1:
                sw_status = GPIO.input(18)
                time.sleep(1)
        time.sleep(1)

    except:
        client.close()
        break

GPIO.cleanup()
print("end")
