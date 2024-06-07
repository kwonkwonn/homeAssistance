#include <stdio.h>
#include <pigpio.h>
#include <signal.h>

#define INA 18  // GPIO18 (physical pin 12 on Raspberry Pi)
#define INB 23  // GPIO23 (physical pin 16 on Raspberry Pi)

void cleanup(int signum) {
    gpioPWM(INA, 0);  // Set INA to low
    gpioPWM(INB, 0);  // Set INB to low
    gpioTerminate();  // Terminate pigpio library
    printf("Program terminated and GPIO cleaned up.\n");
    exit(0);
}

void setup() {
    gpioSetMode(INA, PI_OUTPUT);
    gpioSetMode(INB, PI_OUTPUT);
    gpioPWM(INA, 0);  // Initialize PWM with 0 duty cycle
    gpioPWM(INB, 0);  // Initialize PWM with 0 duty cycle
}

void loop() {
    gpioPWM(INA, 255);  // Set pin INA to high with full PWM duty cycle (255)
    gpioWrite(INB, PI_LOW); // Set pin INB to low
    time_sleep(5);          // Wait for 5 seconds
   
    gpioPWM(INA, 0);   // Set pin INA to low
    gpioWrite(INB, PI_LOW); // Set pin INB to low
    time_sleep(0.2);        // Wait for 0.2 seconds
   
    gpioWrite(INA, PI_LOW); // Set pin INA to low
    gpioPWM(INB, 255);  // Set pin INB to high with full PWM duty cycle (255)
    time_sleep(5);          // Wait for 5 seconds
   
    gpioPWM(INA, 0);   // Set pin INA to low
    gpioWrite(INB, PI_LOW); // Set pin INB to low
    time_sleep(0.2);        // Wait for 0.2 seconds
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
    while (1) {
        loop();
    }

    gpioTerminate();
    return 0;
}