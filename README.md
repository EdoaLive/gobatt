# gobatt
This simple piece of software came to solve a simple problem: my old laptop's battery has a high internal resistance and the BMS wrongly tells the OS that the battery is low when it actually still has several tens of minutes to run.
## How it works
In this phase we just read the battery voltage and compare with a threshold. If it's lower than the threshold it makes a beep and after N beeps it send a hibernation command.

## Future plans:
* Separate OS specific functions (now Windows only) to make it work also on Linux
* Estimate internal resistance from delta V and delta load to obtain an estimated no-load voltage value in order to avoid triggers due to temporary load.