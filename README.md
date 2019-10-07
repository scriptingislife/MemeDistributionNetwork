# ~~Meme~~PopupDistributionNetwork

Client that downloads and opens ~~memes~~ popups on some poor soul's screen.

(But you can change it to memes by changing the `dir` variable.)

![Example](example.gif)

## Run Your Own

```bash
# 1. Start the web server on a red team machine
# Python 2
python -m SimpleHTTPServer

# Python 3
python3 -m http.server


# 2. Set the server IP and port in the Golang file
# host := "<serv-ip>:8000" // CHANGE ME 

# 3. Build the executables
make build



# 4. Copy executables to blue team machine



# 5. Run the program in a loop
# Linux
while true; do ./MemeDistributionNetork > /dev/null; sleep 3; done

# Windows
# Create a scheduled task

```