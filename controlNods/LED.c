#include <stdio.h>
#include <pigpio.h>
#include <signal.h>

#define LED_PIN 6  // GPIO6 (physical pin 31 on Raspberry Pi)

void cleanup(int signum) {
    gpioWrite(LED_PIN, PI_LOW);  // Turn off the LED
    gpioTerminate();  // Terminate pigpio library
    printf("Program terminated and GPIO cleaned up.\n");
    exit(0);
}

void setup() {
    gpioSetMode(LED_PIN, PI_OUTPUT);
    gpioWrite(LED_PIN, PI_HIGH);  // Turn on the LED
}

int main(void) {
    if (gpioInitialise() < 0) {
        fprintf(stderr, "pigpio initialization failed\n");
        return 1;
    }

    // Register the signal handler for cleanup
    signal(SIGINT, cleanup);
    signal(SIGTERM, cleanup);

    setup();
    printf("LED is on. Press Ctrl+C to turn off.\n");

    // Keep the program running to wait for the signal
    while (1) {
        time_sleep(1);
    }

    return 0;
}