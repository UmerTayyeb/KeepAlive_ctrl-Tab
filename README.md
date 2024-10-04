Note:
Ensure that you have Go installed and set up properly on your system

Description:
This Go program simulates a periodic Ctrl + Tab keypress to keep the screen awake. The interval between key presses is randomized between 25 and 50 seconds, ensuring that the system doesn't enter sleep mode.

Steps to Clone and Run the Program:
Clone the Repository
        git clone https://github.com/UmerTayyeb/KeepAlive_ctrl-Tab.git
        cd KeepAlive_ctrl-Tab

Install Dependencies
        go mod tidy

Run the Program
        go run ctrl+tab.go


Creating an Executable File:

Build the Executable(This will create an executable named keep_awake in the current directory (for ubuntu/windows))
        
        go build -o keep_awake

            

    Run the Executable
        ./keep_awake