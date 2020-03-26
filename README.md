# Fudge

I had some success with on Phishing engagements by base64 encoding the contents of a file, putting it into a HTML file and having it decode and drop when the user opened the HTML file. 


# Usage

      -n string
            Name of downloaded file (default "nastypasty.exe")
      -s string
            Source File (default "test.exe")

Use the above flags to set the file to encode, and the filename which will be given to the user when they open the HTML file. 
